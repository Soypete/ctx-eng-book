package tui

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/soypete/authorpedro/internal/ai/agent"
	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
	"github.com/soypete/authorpedro/internal/tui/styles"
)

type ToolCallMsg struct {
	ToolName string
	Args     map[string]any
}

type StreamingMsg struct {
	Content string
}

type StreamTickMsg struct{}

type AgentDoneMsg struct {
	Iterations   int
	ToolCalls    int
	FinalOutput  string
	Conversation []struct{ Role, Content string }
	Err          error
}

type Model struct {
	config       config.Config
	book         outline.Book
	agent        *agent.Agent
	currentCh    int
	currentMod   int
	writing      textinput.Model
	viewport     viewport.Model
	agentOutput  string
	toolCalls    []agent.ToolCall
	status       string
	width        int
	height       int
	agentRunning bool
	streamChan   chan string
}

func NewModel(cfg config.Config) (Model, error) {
	book, err := outline.ParseChaptersMarkdown("chapters.md")
	if err != nil {
		return Model{}, err
	}

	ag, err := agent.New(cfg, book)
	if err != nil {
		return Model{}, err
	}

	m := Model{
		config:      cfg,
		book:        book,
		agent:       ag,
		currentCh:   0,
		currentMod:  0,
		toolCalls:   []agent.ToolCall{},
		agentOutput: "Welcome to AuthorPedro.\n\n" + ag.GetRecentModules(5) + "\n\nWhat would you like to work on today?",
		status:      getStatus(book, 0, 0),
	}

	ag.SetToolCallback(func(tc agent.ToolCall) {
		log.Printf("Tool called: %s %+v", tc.Name, tc.Args)
	})

	m.writing = textinput.New()
	m.writing.Placeholder = "Start writing..."
	m.writing.Focus()

	m.viewport = viewport.New(80, 20)
	m.viewport.SetContent(m.agentOutput)

	m.streamChan = make(chan string, 1000)

	return m, nil
}

func getStatus(book outline.Book, ch, mod int) string {
	if ch >= len(book.Chapters) {
		return "No chapter selected"
	}
	chapter := book.Chapters[ch]
	if mod >= len(chapter.Modules) {
		return chapter.Title
	}
	module := chapter.Modules[mod]
	return chapter.Title + " → " + module.Title
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case StreamingMsg:
		m.agentOutput += msg.Content
		return m, nil

	case StreamTickMsg:
		if !m.agentRunning {
			return m, nil
		}
		select {
		case chunk := <-m.streamChan:
			m.agentOutput += chunk
		default:
		}
		return m, tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg { return StreamTickMsg{} })

	case AgentDoneMsg:
		fmt.Printf("[TUI] AgentDoneMsg: iter=%d toolcalls=%d outputLen=%d err=%v\n",
			msg.Iterations, msg.ToolCalls, len(msg.FinalOutput), msg.Err)
		m.agentRunning = false
		if msg.Err != nil {
			m.agentOutput += styles.Error.Render("Error: "+msg.Err.Error()) + "\n"
		} else {
			m.agentOutput += styles.Success.Render(fmt.Sprintf("Done (iterations:%d, tool_calls:%d)\n", msg.Iterations, msg.ToolCalls))
			m.agentOutput += msg.FinalOutput + "\n"
		}
		m.toolCalls = nil
		log.Printf("[TUI] Conversation has %d messages", len(msg.Conversation))
		for i, msg := range msg.Conversation {
			role := string(msg.Role)
			log.Printf("[TUI] Msg %d: role=%q contentLen=%d", i, role, len(msg.Content))
			if role == "assistant" {
				preview := msg.Content
				if len(preview) > 500 {
					preview = preview[:500] + "..."
				}
				m.agentOutput += styles.AgentOutput.Render("Assistant: " + preview + "\n")
			} else if role == "tool" {
				preview := msg.Content
				preview = strings.ReplaceAll(preview, "<tool_response>", "")
				preview = strings.ReplaceAll(preview, "</tool_response>", "")
				preview = strings.ReplaceAll(preview, "<tool_name>", "")
				preview = strings.ReplaceAll(preview, "</tool_name>", "")
				preview = strings.ReplaceAll(preview, "<result>", "")
				preview = strings.ReplaceAll(preview, "</result>", "")
				preview = strings.ReplaceAll(preview, "<error>", "")
				preview = strings.ReplaceAll(preview, "</error>", "")
				if len(preview) > 500 {
					preview = preview[:500] + "..."
				}
				m.agentOutput += styles.Help.Render("Tool result: " + preview + "\n")
			}
		}
		fmt.Printf("[TUI] Updated agentOutput, new len=%d\n", len(m.agentOutput))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "up":
			m.viewport.LineUp(3)
			return m, nil

		case "down":
			m.viewport.LineDown(3)
			return m, nil

		case "pgup":
			m.viewport.PageUp()
			return m, nil

		case "pgdown":
			m.viewport.PageDown()
			return m, nil

		case "ctrl+shift+c":
			if m.agentOutput != "" {
				err := clipboard.WriteAll(m.agentOutput)
				if err == nil {
					m.agentOutput += styles.Success.Render("Copied to clipboard!") + "\n"
				}
			}
			return m, nil

		case "enter":
			content := m.writing.Value()
			if content == "/exit" {
				return m, tea.Quit
			}
			if content != "" && !m.agentRunning {
				m.agentRunning = true
				m.agentOutput += styles.Title.Render("🤔 Agent working...") + "\n"
				m.agentOutput += styles.Help.Render("Query: " + content + "\n")
				m.agentOutput += styles.Help.Render("Tools: enabled\n")
				m.writing.SetValue("")

				ag := m.agent

				cmd := func() tea.Msg {
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()

					streamChan := m.streamChan
					ag.SetStreamCallback(func(chunk string) {
						select {
						case streamChan <- chunk:
						default:
						}
					})

					result, err := ag.Execute(ctx, content)
					if err != nil {
						return AgentDoneMsg{Err: err}
					}
					log.Printf("[TUI] AgentExecute returned: Iterations=%d, ToolCalls=%d, FinalResponse len=%d, Conversation len=%d",
						result.Iterations, result.ToolCallsMade, len(result.FinalResponse), len(result.Conversation))
					for i, m := range result.Conversation {
						log.Printf("[TUI] result.Conversation[%d]: role=%q contentLen=%d", i, string(m.Role), len(m.Content))
					}
					conv := make([]struct{ Role, Content string }, len(result.Conversation))
					for i, msg := range result.Conversation {
						conv[i] = struct{ Role, Content string }{string(msg.Role), msg.Content}
					}
					return AgentDoneMsg{
						Iterations:   result.Iterations,
						ToolCalls:    result.ToolCallsMade,
						FinalOutput:  result.FinalResponse,
						Conversation: conv,
						Err:          err,
					}
				}
				return m, tea.Batch(cmd, tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg { return StreamTickMsg{} }))
			}

		case "ctrl+a":
			m.agentOutput += styles.Help.Render("[Thinking toggled - now always on for agent harness]\n")
			return m, nil

		case "ctrl+p":
			if m.agentRunning {
				return m, nil
			}
			query := m.writing.Value()
			if query == "" {
				query = "Plan what to work on today for the book"
			}
			m.agentRunning = true
			m.agentOutput += styles.Title.Render("📋 Planning mode...") + "\n"
			m.agentOutput += styles.Help.Render("Query: " + query + "\n")
			m.writing.SetValue("")
			cmd := func() tea.Msg {
				log.Println("[Agent] Starting plan mode...")
				result, err := m.agent.ExecutePlanMode(context.Background(), query)
				conv := make([]struct{ Role, Content string }, len(result.Conversation))
				for i, msg := range result.Conversation {
					conv[i] = struct{ Role, Content string }{string(msg.Role), msg.Content}
				}
				return AgentDoneMsg{
					Iterations:   result.Iterations,
					ToolCalls:    result.ToolCallsMade,
					FinalOutput:  result.FinalResponse,
					Conversation: conv,
					Err:          err,
				}
			}
			return m, cmd

		case "ctrl+n":
			if m.currentMod < len(m.book.Chapters[m.currentCh].Modules)-1 {
				m.currentMod++
			} else if m.currentCh < len(m.book.Chapters)-1 {
				m.currentCh++
				m.currentMod = 0
			}
			m.status = getStatus(m.book, m.currentCh, m.currentMod)
			m.loadCurrentModule()
		}
	}

	m.writing, cmd = m.writing.Update(msg)
	return m, cmd
}

func (m *Model) saveCurrentModule(content string) error {
	if m.currentCh >= len(m.book.Chapters) {
		return os.ErrInvalid
	}
	chapter := m.book.Chapters[m.currentCh]
	if m.currentMod >= len(chapter.Modules) {
		return os.ErrInvalid
	}
	module := chapter.Modules[m.currentMod]

	relPath, err := m.book.ResolvePath(chapter.Slug, module.Slug)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(m.config.BookPath, relPath)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	existing, _ := os.ReadFile(fullPath)
	if len(existing) > 0 {
		existing = append(existing, '\n')
		content = string(existing) + content
	}

	return os.WriteFile(fullPath, []byte(content), 0644)
}

func (m *Model) loadCurrentModule() {
	if m.currentCh >= len(m.book.Chapters) {
		return
	}
	chapter := m.book.Chapters[m.currentCh]
	if m.currentMod >= len(chapter.Modules) {
		return
	}
	module := chapter.Modules[m.currentMod]

	relPath, err := m.book.ResolvePath(chapter.Slug, module.Slug)
	if err != nil {
		m.agentOutput += styles.Error.Render("Error resolving path: "+err.Error()) + "\n"
		return
	}

	fullPath := filepath.Join(m.config.BookPath, relPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		m.agentOutput += styles.AgentOutput.Render("New module: "+module.Title) + "\n"
		return
	}

	m.agentOutput += styles.Success.Render("Loaded: "+module.Title) + "\n"
	m.agentOutput += string(content) + "\n"
}

func (m Model) View() string {
	header := styles.Title.Render("AuthorPedro")
	if m.agentRunning {
		header = styles.Title.Render("AuthorPedro") + " " + styles.Error.Render("⟳")
	}
	help := styles.Help.Render("↑↓ scroll | ctrl+a ask (think) | ctrl+p plan | ctrl+n next | ctrl+shift+c copy | /exit quit")

	m.viewport.Width = m.width - 4
	m.viewport.Height = m.height - 12
	if m.width > 4 && m.height > 12 {
		m.viewport.SetContent(m.agentOutput)
		if m.viewport.TotalLineCount() > m.viewport.Height {
			m.viewport.GotoBottom()
		}
	}

	outputBox := styles.AgentOutput.Width(m.width - 4).Render(m.viewport.View())

	var toolCallsBox string
	if len(m.toolCalls) > 0 {
		toolCallsBox = styles.Help.Width(m.width - 4).Render("Tool calls: ")
		for _, tc := range m.toolCalls {
			toolCallsBox += styles.Help.Render(fmt.Sprintf("\n  • %s(%v)", tc.Name, tc.Args))
		}
	}

	statusBox := styles.StatusBar.Width(m.width - 4).Render(m.status)
	inputBox := styles.Input.Width(m.width - 4).Render("> " + m.writing.View())

	return lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		"",
		outputBox,
		toolCallsBox,
		"",
		statusBox,
		"",
		inputBox,
		"",
		help,
	)
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	var s string
	for i > 0 {
		s = string(rune('0'+i%10)) + s
		i /= 10
	}
	return s
}

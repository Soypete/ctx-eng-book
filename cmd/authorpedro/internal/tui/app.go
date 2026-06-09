package tui

import (
	"context"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/soypete/authorpedro/internal/agent"
	"github.com/soypete/authorpedro/internal/config"
	"github.com/soypete/authorpedro/internal/outline"
	"github.com/soypete/authorpedro/internal/tui/styles"
)

type Model struct {
	config       config.Config
	book         outline.Book
	agent        *agent.Agent
	currentCh    int
	currentMod   int
	writing      textinput.Model
	agentOutput  string
	status       string
	width        int
	height       int
	agentRunning bool
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

	ti := textinput.New()
	ti.Placeholder = "Start writing..."
	ti.Focus()

	m := Model{
		config:      cfg,
		book:        book,
		agent:       ag,
		currentCh:   0,
		currentMod:  0,
		writing:     ti,
		agentOutput: "Welcome to AuthorPedro. Ready to write.\n\nPress ctrl+a to ask the agent for help.\n",
		status:      getStatus(book, 0, 0),
	}

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

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "enter":
			content := m.writing.Value()
			if content != "" {
				err := m.saveCurrentModule(content)
				if err != nil {
					m.agentOutput += styles.Error.Render("Error: "+err.Error()) + "\n"
				} else {
					preview := content
					if len(content) > 50 {
						preview = content[:50] + "..."
					}
					m.agentOutput += styles.Success.Render("Saved: "+preview) + "\n"
				}
				m.writing.SetValue("")
			}

		case "ctrl+a":
			if m.agentRunning {
				return m, nil
			}
			query := m.writing.Value()
			if query == "" {
				query = "Write an introduction for this module"
			}
			m.agentRunning = true
			m.agentOutput += styles.Title.Render("Agent working...") + "\n"
			return m, func() tea.Msg {
				result, err := m.agent.Execute(context.Background(), query)
				m.agentRunning = false
				if err != nil {
					m.agentOutput += styles.Error.Render("Error: "+err.Error()) + "\n"
				} else {
					m.agentOutput += styles.Success.Render("Agent complete (") +
						styles.Help.Render("iterations: "+itoa(result.Iterations)+")") + "\n"
					m.agentOutput += result.FinalResponse + "\n"
				}
				return tea.Msg("agent_done")
			}

		case "ctrl+n":
			if m.currentMod < len(m.book.Chapters[m.currentCh].Modules)-1 {
				m.currentMod++
			} else if m.currentCh < len(m.book.Chapters)-1 {
				m.currentCh++
				m.currentMod = 0
			}
			m.status = getStatus(m.book, m.currentCh, m.currentMod)
			m.loadCurrentModule()

		case "ctrl+p":
			if m.currentMod > 0 {
				m.currentMod--
			} else if m.currentCh > 0 {
				m.currentCh--
				m.currentMod = len(m.book.Chapters[m.currentCh].Modules) - 1
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
	help := styles.Help.Render("ctrl+a ask agent | ctrl+n next | ctrl+p prev | enter save | esc quit")

	outputBox := styles.AgentOutput.Width(m.width - 4).Render(m.agentOutput)
	statusBox := styles.StatusBar.Width(m.width - 4).Render(m.status)
	inputBox := styles.Input.Width(m.width - 4).Render("> " + m.writing.View())

	return lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		"",
		outputBox,
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

package styles

import "github.com/charmbracelet/lipgloss"

var (
	Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true).
		Padding(0, 1)

	StatusBar = lipgloss.NewStyle().
			Foreground(lipgloss.Color("250")).
			Background(lipgloss.Color("235")).
			Padding(0, 2).
			Width(60)

	AgentOutput = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Background(lipgloss.Color("236")).
			Padding(1, 2).
			Height(15).
			Width(60)

	Input = lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Background(lipgloss.Color("235")).
		Padding(0, 1)

	Success = lipgloss.NewStyle().
		Foreground(lipgloss.Color("82"))

	Error = lipgloss.NewStyle().
		Foreground(lipgloss.Color("204"))

	Help = lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Italic(true)
)

var (
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("86")).
			Padding(1)

	FocusedBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205")).
			BorderBackground(lipgloss.Color("236")).
			Padding(1)
)

func StatusText(chapterTitle, moduleTitle string) string {
	if moduleTitle == "" {
		return StatusBar.Render(chapterTitle)
	}
	return StatusBar.Render(chapterTitle + " → " + moduleTitle)
}

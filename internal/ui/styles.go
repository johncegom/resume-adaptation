package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	ColorPrimary = lipgloss.Color("#7C3AED")
	ColorAccent  = lipgloss.Color("#EC4899")
	ColorSuccess = lipgloss.Color("#10B981")
	ColorError   = lipgloss.Color("#EF4444")
	ColorMuted   = lipgloss.Color("#9CA3AF")
	ColorBorder  = lipgloss.Color("#4B5563")

	// Styles
	StyleTitle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			MarginLeft(1).
			MarginRight(1).
			Padding(0, 1)

	StyleHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F3F4F6")).
			Background(ColorPrimary).
			Bold(true).
			Padding(0, 1).
			MarginBottom(1)

	StyleSubTitle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Italic(true)

	StyleBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(1, 2).
			Margin(1)

	StyleText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F3F4F6"))

	StyleSuccess = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true)

	StyleError = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	StyleHighlight = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true)

	StyleMuted = lipgloss.NewStyle().
			Foreground(ColorMuted)
)

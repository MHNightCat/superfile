package components

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	minimumHeight   = 35
	minimumWidth    = 96
	bottomBarHeight = 13

	terminalTooSmall    lipgloss.Style
	terminalMinimumSize lipgloss.Style

	borderStyle lipgloss.Style
	cursorStyle lipgloss.Style
)

var (
	sideBarWidth    = 20
	sideBarTitle    lipgloss.Style
	sideBarItem     lipgloss.Style
	sideBarSelected lipgloss.Style
)

var (
	filePanelTopFolderIcon lipgloss.Style
	filePanelTopPath       lipgloss.Style
	filePanelItem          lipgloss.Style
	filePanelItemSelected  lipgloss.Style
)

func LoadThemeConfig() {
	terminalTooSmall = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.TerminalTooSmallError))
	terminalMinimumSize = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.TerminalSizeCurrect))

	borderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Border))
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Cursor)).Bold(true)

	sideBarTitle = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.SideBarTitle)).Bold(true)
	sideBarItem = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.SideBarItem))

	sideBarSelected = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.SideBarSelected)).Bold(true)

	filePanelTopFolderIcon = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.FilePanelTopFolderIcon)).Bold(true)
	filePanelTopPath = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.FilePanelTopPath)).Bold(true)
	filePanelItem = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.FilePanelItem))
	filePanelItemSelected = lipgloss.NewStyle().Foreground(lipgloss.Color(theme.FilePanelItemSelected))

}
func SideBarBoardStyle(height int, focus focusPanelType) lipgloss.Style {
	if focus == sideBarFocus {
		return lipgloss.NewStyle().
			BorderStyle(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color(theme.SideBarFocus)).
			Width(sideBarWidth).
			Height(height).Bold(true)
	} else {
		return lipgloss.NewStyle().
			BorderStyle(lipgloss.HiddenBorder()).
			Width(sideBarWidth).
			Height(height).Bold(true)
	}
}

func FilePanelBoardStyle(height int, width int, focusType filePanelFocusType, borderBottom string) lipgloss.Style {
	leftBorder := ""
	rightBorder := ""
	for i := 0; i < height; i++ {
		if i == 1 {
			leftBorder += "┣"
			rightBorder += "┫"
		} else {
			leftBorder += "┃"
			rightBorder += "┃"
		}
	}
	filePanelBottomBoard := lipgloss.Border{
		Top:         "━",
		Bottom:      borderBottom,
		Left:        leftBorder,
		Right:       rightBorder,
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┗",
		BottomRight: "┛",
	}
	return lipgloss.NewStyle().
		Border(filePanelBottomBoard, true, true, true, true).
		BorderForeground(lipgloss.Color(FilePanelFocusColor(focusType))).
		Width(width).
		Height(height)
}

func ProcsssBarBoarder(height int, width int, borderBottom string, focusType focusPanelType) lipgloss.Style {
	filePanelBottomBoard := lipgloss.Border{
		Top:         "━",
		Bottom:      borderBottom,
		Left:        "┃",
		Right:       "┃",
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┗",
		BottomRight: "┛",
	}
	if focusType == processBarFocus {
		return lipgloss.NewStyle().
			Border(filePanelBottomBoard, true, true, true, true).
			BorderForeground(lipgloss.Color(theme.BottomBarFocus)).
			Width(width).
			Height(height).Bold(true)
	} else {
		return lipgloss.NewStyle().
			Border(filePanelBottomBoard, true, true, true, true).
			BorderForeground(lipgloss.Color(theme.Border)).
			Width(width).
			Height(height).Bold(true)
	}
}

func FilePanelDividerStyle(focusType filePanelFocusType) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(FilePanelFocusColor(focusType))).Bold(true)
}

func TruncateText(text string, maxChars int) string {
	if len(text) <= maxChars {
		return text
	}
	return text[:maxChars-3] + "..."
}

func TruncateTextBeginning(text string, maxChars int) string {
	if len(text) <= maxChars {
		return text
	}
	runes := []rune(text)
	charsToKeep := maxChars - 3
	truncatedRunes := append([]rune("..."), runes[len(runes)-charsToKeep:]...)
	return string(truncatedRunes)
}

func PrettierName(name string, width int, isDir bool, isSelected bool) string {
	style := getElementIcon(name, isDir)
	if isSelected {
		return StringColorRender(style.color).Render(style.icon) + "  " + filePanelItemSelected.Render(TruncateText(name, width))
	} else {
		return StringColorRender(style.color).Render(style.icon) + "  " + filePanelItem.Render(TruncateText(name, width))
	}
}

// CHOOSE STYLE FUNCTION
func FilePanelFocusColor(focusType filePanelFocusType) string {
	if focusType == noneFocus {
		return theme.Border
	} else {
		return theme.FilePanelFocus
	}
}

func FilePanelBoard(focusType filePanelFocusType) lipgloss.Border {
	if focusType == noneFocus {
		return lipgloss.RoundedBorder()
	} else {
		return lipgloss.ThickBorder()
	}
}

func GenerateBottomBorder(countString string, width int) string {
	result := ""
	for i := 0; i < width-len(countString); i++ {
		result += "━"
	}
	return result + "┫" + countString + "┣"
}

func StringColorRender(color string) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(color))
}

// logger/logger.go
package logger

import (
    "os"
    "time"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/log"
)

var styles = log.DefaultStyles()

func init() {
    styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
        SetString("ERROR!!").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("204")).
        Foreground(lipgloss.Color("0"))

    styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
        SetString("INFO").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("34")).  
        Foreground(lipgloss.Color("15"))

    styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
        SetString("WARNING").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("214")). 
        Foreground(lipgloss.Color("0"))

    styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
        SetString("DEBUG").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("33")).  
        Foreground(lipgloss.Color("15"))
}

func NewLogger() *log.Logger {
    logger := log.New(os.Stderr)

    logger.SetStyles(styles)
    logger.SetReportTimestamp(true)
    logger.SetTimeFormat(time.Stamp)

    return logger
}

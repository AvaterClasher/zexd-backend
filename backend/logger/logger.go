// logger/logger.go
package logger

import (
    "os"
    "time"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/log"
)

var styles = log.DefaultStyles()
var stylesForHttp = log.DefaultStyles()

func init() {
    styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
        SetString("ERROR").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("204")).
        Foreground(lipgloss.Color("15"))

    styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
        SetString("INFO").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("34")).  
        Foreground(lipgloss.Color("15"))

    styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
        SetString("WARNING").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("214")). 
        Foreground(lipgloss.Color("15"))

    styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
        SetString("FATAL").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("#7F00FF")).
        Foreground(lipgloss.Color("15"))

    stylesForHttp.Levels[log.InfoLevel] = lipgloss.NewStyle().
        SetString("HTTP").
        Padding(0, 1, 0, 1).
        Background(lipgloss.Color("12")).  
        Foreground(lipgloss.Color("15"))
}

func NewLogger() *log.Logger {
    logger := log.New(os.Stderr)

    logger.SetStyles(styles)
    logger.SetReportTimestamp(true)
    logger.SetTimeFormat(time.Stamp)

    return logger
}

func NewLoggerforHttp() *log.Logger {
    logger := log.New(os.Stdout)

    logger.SetStyles(stylesForHttp)
    logger.SetReportTimestamp(true)
    logger.SetTimeFormat(time.Stamp)

    return logger
}

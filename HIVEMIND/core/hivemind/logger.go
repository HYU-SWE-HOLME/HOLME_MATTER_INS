package hivemind

import (
	"github.com/fatih/color"
	"time"
)

func SuccessfulLog(msg string) {
	title := color.New(color.FgBlack).Add(color.BgGreen).Add(color.Bold)
	_, err := title.Printf("SUCCESSFUL")
	if err != nil {
		return
	}
	timestamp := time.Unix(time.Now().Unix(), 0).String()
	color.Green("  %s  %s", msg, timestamp)
}

func InformativeLog(msg string) {
	title := color.New(color.FgBlack).Add(color.BgWhite).Add(color.Bold)
	_, err := title.Printf("INFO")
	if err != nil {
		return
	}
	timestamp := time.Unix(time.Now().Unix(), 0).String()
	color.White("  %s  %s", msg, timestamp)
}

func WarningLog(msg string) {
	title := color.New(color.FgBlack).Add(color.BgYellow).Add(color.Bold)
	_, err := title.Printf("WARN")
	if err != nil {
		return
	}
	timestamp := time.Unix(time.Now().Unix(), 0).String()
	color.Yellow("  %s  %s", msg, timestamp)
}

func ErrorLog(msg string) {
	title := color.New(color.FgBlack).Add(color.BgRed).Add(color.Bold)
	_, err := title.Printf("FATAL")
	if err != nil {
		return
	}
	timestamp := time.Unix(time.Now().Unix(), 0).String()
	color.Red("  %s  %s", msg, timestamp)
}

// Package gogger is a console logger designed for simplicity and great features at the same time!
package gogger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

const (
	// VerboseLevel represents the level at which Verbose logs can be displayed.
	VerboseLevel int = 0
	// InfoLevel represents the level at which Info logs can be displayed.
	InfoLevel int = 1
	// WarnLevel represents the level at which Warn logs can be displayed.
	WarnLevel int = 2
	// ErrorLevel represents the level at which Error logs can be displayed.
	ErrorLevel int = 3
	// FatalLevel represents the level at which Fatal logs can be displayed.
	FatalLevel int = 4
	
	// UKDateTimeFormat represents the time in the UK standard format.
	UKDateTimeFormat string = "02-01-2006 15:04:05"
	// USDateTimeFormat represents the time in the US standard format.
	USDateTimeFormat string = "01-02-2006 15:04:05"
)

var (
	// ProjectNameColor is the color used for the project name in logs.
	ProjectNameColor = color.FgGreen

	// LevelColors is a list of github.com/faith/color color values to use for each level
	LevelColors = [5]color.Attribute{color.FgMagenta,color.FgCyan,color.FgYellow,color.FgRed,color.FgRed}
	// LevelStrings represents the strings used in logs to represent that particular level.
	LevelStrings [5]string = [5]string{"verbose","info   ","warn   ","error  ","fatal  "}
)

// GoggerLogger is the logging handler used for printing messages to the console.
type GoggerLogger struct {
	ProjectName string
	Level int
	TimeFormat string
	Colorful bool
}

// SetLevel will change the current level of the GoggerLogger.
func (g *GoggerLogger) SetLevel(level int) error {
	var possibleLevels []int = []int{VerboseLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel}
	
	for _, possibleLevel := range possibleLevels {
		if possibleLevel == level {
			g.Level = level
			return nil
		} 
	}
	return fmt.Errorf("gogger: provided level (%d) is invalid", level)
}

func (g *GoggerLogger) fmtMsg(level int, message string) string {
	if g.Colorful {
		return fmt.Sprintf("%s %s [%s]: %s", time.Now().Format(g.TimeFormat), color.New(LevelColors[level]).Sprint(LevelStrings[level]), color.New(ProjectNameColor).Sprint(g.ProjectName), message)
	}

	return fmt.Sprintf("%s %s [%s]: %s", time.Now().Format(g.TimeFormat), LevelStrings[level], g.ProjectName, message)
}

// Verbose messages are at the lowest log level and represent very detailed information.
func (g *GoggerLogger) Verbose(message string) bool {
	const requiredLevel int = VerboseLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, message))
	return true
}

// Verbosef is like Verbose but with Sprintf.
func (g *GoggerLogger) Verbosef(message string, formatOptions ...interface{}) bool {
	const requiredLevel int = VerboseLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, fmt.Sprintf(message, formatOptions...)))
	return true
}

// Info logs are at the default logging level and represent informative information without being too specific.
func (g *GoggerLogger) Info(message string) bool {
	const requiredLevel int = InfoLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, message))
	return true
}

// Infof is like Info but with Sprintf.
func (g *GoggerLogger) Infof(message string, formatOptions ...interface{}) bool {
	const requiredLevel int = InfoLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, fmt.Sprintf(message, formatOptions...)))
	return true
}

// Warn messages at used to represent something bad happening without it being catastrophic.
func (g *GoggerLogger) Warn(message string) bool {
	const requiredLevel int = WarnLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, message))
	return true
}

// Warnf is like Warn but with Sprintf.
func (g *GoggerLogger) Warnf(message string, formatOptions ...interface{}) bool {
	const requiredLevel int = WarnLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, fmt.Sprintf(message, formatOptions...)))
	return true
}

// Error messages are used to represent an error that needs logging but was still handled.
func (g *GoggerLogger) Error(message string) bool {
	const requiredLevel int = ErrorLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, message))
	return true
}

// Errorf is like Error but with Sprintf.
func (g *GoggerLogger) Errorf(message string, formatOptions ...interface{}) bool {
	const requiredLevel int = ErrorLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, fmt.Sprintf(message, formatOptions...)))
	return true
}

// Fatal messages denote that something really bad happened and the program has shut down.
func (g *GoggerLogger) Fatal(message string) bool {
	const requiredLevel int = FatalLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, message))
	return true
}

// Fatalf is like Fatal but with Sprintf.
func (g *GoggerLogger) Fatalf(message string, formatOptions ...interface{}) bool {
	const requiredLevel int = FatalLevel
	
	if g.Level > requiredLevel {
		return false
	}

	fmt.Println(g.fmtMsg(requiredLevel, fmt.Sprintf(message, formatOptions...)))
	return true
}

// New initialises a new GoggerLogger.
func New(projectName string, timeFormat string, colorful bool) GoggerLogger {
	return GoggerLogger{
		ProjectName: projectName,
		TimeFormat: timeFormat,
		Level: InfoLevel,
		Colorful: colorful,
	}
}
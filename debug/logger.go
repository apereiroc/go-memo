package debug

import (
	"fmt"
	"log"
	"os"
	"path"
)

const debugFileName = "go-memo-debug.log"

// colour codes
const (
	// hex colour #E82424 → RGB(232, 36, 36)
	colourRed = "\033[38;2;232;36;36m"
	// hex colour #76946A: RGB(118, 148, 106)
	colourGreen = "\033[38;2;118;148;106m"
	// hex colour #A3D4D5:  RGB(163, 228, 229)
	colourBlue = "\033[38;2;163;228;229m"
	// hex colour #FF9E3B: RGB(255, 211, 60)
	colourYellow = "\033[38;2;255;211;60m"
	// RGB(243, 160, 132)
	colourOrange = "\033[38;2;243;160;132m"
	// RGB(130, 149, 205)
	colourViolet = "\033[38;2;130;149;205m"
	// reset
	colourReset = "\033[0m"
)

// format
const (
	textBold = "\033[1m"
)

const (
	colourDebug     = colourBlue
	colourInfo      = colourGreen
	colourWarn      = colourYellow
	colourError     = colourRed
	colourDebugInfo = colourViolet
)

var (
	debugEnabled bool = false
	debugFile    *os.File
	debugLogger  *log.Logger
)

// Initialise debug and open debug file to write debug information to it
func Start() {
	if debugEnabled {
		return
	}

	debugFullFileName := path.Join(os.TempDir(), debugFileName)

	var err error
	debugFile, err = os.OpenFile(debugFullFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open debug file:", err)
		return
	}

	infoStr := colourInfo + "INFO" + colourReset
	fileStr := colourYellow + debugFullFileName + colourReset
	fmt.Printf("[%s]: debug info will be saved to %s\n\n", infoStr, fileStr)

	debugLogger = log.New(debugFile, "", log.LstdFlags)
	debugEnabled = true

	debugLogger.Printf("--- Begin ---\n\n")
}

// Close debug file and set debug flag to false
func Stop() {
	if debugFile != nil {
		debugLogger.SetPrefix("\n")
		debugLogger.Printf("--- End ---\n\n")
		if err := debugFile.Close(); err != nil {
			log.Fatal(err)
		}
	}
	debugEnabled = false
}

// Core debug function
// Writes to the debugFile with a coloured message and a timestamp
func doLog(debugLevel string, colour string, msg string) {
	if !debugEnabled {
		return
	}

	// write useful information
	// file name, line number, function name
	fileName, lineNumber, functionName := trace()
	fileName = textBold + colourDebugInfo + fileName + colourReset
	lineString := textBold + colourDebugInfo + fmt.Sprintf("%d", lineNumber) + colourReset
	functionName = textBold + colourDebugInfo + functionName + colourReset
	prettyDebugInformation := fmt.Sprintf("%s:%s at function %s", fileName, lineString, functionName)
	debugLogger.SetPrefix("")
	debugLogger.Println(prettyDebugInformation)

	// set coloured debugLevel
	colouredLevel := colour + debugLevel + colourReset
	debugLogger.SetPrefix(fmt.Sprintf("[%s] ", colouredLevel))

	// write coloured msg
	colouredMsg := ": " + colour + msg + colourReset
	debugLogger.Println(colouredMsg)
}

func Info(msg string) {
	doLog("INFO", colourInfo, msg)
}

func Infof(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	Info(msg)
}

func Debug(msg string) {
	doLog("DEBUG", colourDebug, msg)
}

func Debugf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	Debug(msg)
}

func Warn(msg string) {
	doLog("WARN", colourWarn, msg)
}

func Warnf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	Warn(msg)
}

func Error(err error) {
	doLog("ERROR", colourError, err.Error())
}

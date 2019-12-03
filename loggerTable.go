package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"os"
)

type LogDetails struct {
	Loglevel   string `json:"level"`
	TimeStamp  string `json:"ts"`
	Msg        string `json:"msg"`
	Caller     string `json:"caller"`
	DevInfo    string `json:"deviceInfo"`
	StackTrace string `json:"stacktrace"`
}

type LgLevel struct {
	err   bool
	info  bool
	warn  bool
	debug bool
}

func readJSONFile(fileName string, filePath string, level LgLevel) {
	// File Reader
	file, err := os.Open(filePath + fileName)

	if err != nil {
		fmt.Print("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	_ = file.Close()

	// Initialize table format
	t := table.NewWriter()
	// you can also instantiate the object directly
	tTemp := table.Table{}
	tTemp.Render() // just to avoid the compile error of not using the object
	// HEADER fields to be added here
	t.AppendHeader(table.Row{"Level", "ts", "msg", "comment", "deviceInfo", "stacktrace"})

	t.SetAutoIndex(true)
	for _, eachline := range txtLines {
		var ld LogDetails
		_ = json.Unmarshal([]byte(eachline), &ld)
		
		if (level.err == true && "error" == ld.Loglevel) || (level.debug == true && "debug" == ld.Loglevel) || (level.info == true && "info" == ld.Loglevel) || (level.warn == true && "warn" == ld.Loglevel) {

			if ld.DevInfo != "" && ld.StackTrace == "" && strings.Contains(ld.Msg.pattern) {
				t.AppendRow(table.Row{ld.Loglevel, ld.TimeStamp, ld.Msg, ld.Caller, string(ld.DevInfo)})
			}
			if ld.StackTrace != "" && ld.DevInfo == "" && strings.Contains(ld.Msg.pattern) {
				t.AppendRow(table.Row{ld.Loglevel, ld.TimeStamp, ld.Msg, ld.Caller, "", string(ld.StackTrace)})
			}
			if ld.StackTrace != "" && ld.DevInfo != "" && strings.Contains(ld.Msg.pattern) {
				t.AppendRow(table.Row{ld.Loglevel, ld.TimeStamp, ld.Msg, ld.Caller, string(ld.DevInfo), string(ld.StackTrace)})
			}
			if ld.DevInfo == "" && ld.StackTrace == "" && strings.Contains(ld.Msg.pattern) {
				t.AppendRow(table.Row{ld.Loglevel, ld.TimeStamp, ld.Msg, ld.Caller})
			}
		}
	}
	t.SetStyle(table.StyleBold)
	colorBOnW := text.Colors{text.BgWhite, text.FgBlack}
	t.SetColorsHeader([]text.Colors{colorBOnW, colorBOnW, colorBOnW, colorBOnW, colorBOnW, colorBOnW})
	t.SetColors([]text.Colors{{text.FgGreen}, {text.FgHiBlue}, {text.FgHiWhite}, {text.FgHiYellow}, {text.FgCyan}, {text.FgHiMagenta}})
	t.SetColorsFooter([]text.Colors{{}, {}, colorBOnW, colorBOnW})
	t.SetAllowedColumnLengths([]int{0, 6, 30, 20, 30, 30})
	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateRows = true
	t.SetCaption(fileName)

	fmt.Println(t.Render())
	// To render into html enable this option
	//fmt.Println(t.RenderHTML())

}

func main() {
	fmt.Println("Starting-Analysig-JSON-file", )
	fileName := "sample.json"
	filePath := "/home/demo/"
	pattern := ""
	level := LgLevel{
		err:   true,
		info:  true,
		warn:  false,
		debug: true,
	}
	readJSONFile(fileName, filePath, level)
}


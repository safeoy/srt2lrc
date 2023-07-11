package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Subtitle struct {
	Index     int
	StartTime string
	EndTime   string
	Text      string
}

func parseSRT(data []byte) []Subtitle {
	lines := strings.Split(string(data), "\n")
	var subs []Subtitle
	var sub Subtitle
	for _, line := range lines {
		if line == "" {
			subs = append(subs, sub)
			sub = Subtitle{}
			continue
		}

		if sub.Index == 0 {
			fmt.Sscanf(line, "%d", &sub.Index)
			continue
		}

		if sub.StartTime == "" {
			times := strings.Split(line, " --> ")
			sub.StartTime = times[0]
			sub.EndTime = times[1]
			continue
		}

		sub.Text = line
	}
	return subs
}

func srtTimeToLRCTime(srtTime string) string {
	t, _ := time.Parse("15:04:05.000", srtTime)
	return fmt.Sprintf("[%02d:%02d.%02d]", t.Minute(), t.Second(), t.Nanosecond()/1e7)
}

func convertToLRC(subs []Subtitle) string {
	var lrc []string
	for _, sub := range subs {
		if sub.Index == 0 {
			continue
		}
		startTime := srtTimeToLRCTime(sub.StartTime)
		//endTime := srtTimeToLRCTime(sub.EndTime)
		//lrc = append(lrc, startTime+sub.Text, endTime+"\n")
		lrc = append(lrc, startTime+sub.Text)
	}
	return strings.Join(lrc, "\n")
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: srt2lrc input.srt output.lrc")
		os.Exit(1)
	}

	inputFile := args[1]
	outputFile := args[2]

	data, _ := ioutil.ReadFile(inputFile)
	subs := parseSRT(data)
	lrc := convertToLRC(subs)

	ioutil.WriteFile(outputFile, []byte(lrc), 0644)
}

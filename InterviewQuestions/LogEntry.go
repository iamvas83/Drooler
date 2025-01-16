package interviewquestions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Average response time

type LogEntry struct {
	RequestID    int
	Timestamp    string
	UserID       string
	Method       string
	Path         string
	ResponseTime int
	DBTime       int
}

func LogAnalysis() int {
	avrresponseTime := 0
	file, err := os.Open("test\\logs.txt")
	if err != nil {
		fmt.Print("error occured reading the file", err)
	}
	defer file.Close()
	var entries []LogEntry
	scanner := bufio.NewScanner(file)
	isHeader := true
	for scanner.Scan() {
		line := scanner.Text()

		if isHeader {
			isHeader = false
			headers := strings.Split(line, ",")
			fmt.Printf("Headers %v\n", headers)
			continue
		}
		fields := strings.Split(line, ",")

		if len(fields) < 7 {
			continue
		}

		responseTime, err := strconv.Atoi(strings.TrimSuffix(fields[5], "ms"))

		if err != nil {
			fmt.Print("Invalid response time", err)
			continue
		}

		entry := &LogEntry{
			RequestID:    parseInt(fields[0]),
			Timestamp:    fields[1],
			UserID:       fields[2],
			Method:       fields[3],
			Path:         fields[4],
			ResponseTime: responseTime,
			DBTime:       parseInt(strings.TrimSuffix(fields[6], "ms")),
		}

		entries = append(entries, *entry)

		total := 0

		for _, ent := range entries {
			total += ent.ResponseTime
		}

		avrresponseTime = total / len(entries)

	}
	fmt.Printf("Average response time in the logs : %d ", avrresponseTime)
	return avrresponseTime
}

func parseInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return val
}

package main

import (
	"bufio"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	t := os.Args[1]
	sec, unit, _ := humanize.ParseSI(t)
	if unit != "s" {
		return
	}
	ms := time.Duration(sec * 1000)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		os.Stdout.WriteString(scanner.Text() + "\n")
		time.Sleep(ms * time.Millisecond)
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}

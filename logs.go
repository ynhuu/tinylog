package tinylog

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func Init() *Logs {
	l := &Logs{nil, 0, "log", 5000}
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.SetOutput(l)
	return l
}

type Logs struct {
	fp         *os.File
	lines      int
	name       string
	totalLines int
}

func (l *Logs) SetName(name string) {
	l.name = name
}

func (l *Logs) SetLines(lines int) {
	l.totalLines = lines
}

func (l *Logs) refresh() {
	fp, _ := os.OpenFile(l.name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	scanner := bufio.NewScanner(fp)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	l.fp = fp
	l.lines = lines
}

func (l *Logs) Write(b []byte) (int, error) {
	if l.fp == nil {
		l.refresh()
	}
	if l.lines == l.totalLines {
		_ = l.Closer()
		_ = os.Rename(l.name, l.name+".0")
		l.fp, _ = os.OpenFile(l.name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		l.lines = 0
	} else {
		l.lines++
	}
	b = bytes.ReplaceAll(b, []byte("\n"), nil)
	b = bytes.ReplaceAll(b, []byte("  "), nil)
	b = append(b, '\n')
	_, _ = os.Stdout.Write(b)
	return l.fp.Write(b)
}

func (l *Logs) Closer() error {
	return l.fp.Close()
}

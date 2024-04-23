package tinylog

import (
	"log"
	"testing"
)

func TestLogs(t *testing.T) {
	elog := Init()
	elog.SetLines(1)
	elog.SetName("test")
	log.Println("1111")
}

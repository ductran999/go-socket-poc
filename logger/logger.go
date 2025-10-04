package logger

import (
	"log"
	"strings"
)

func Warn(message ...string) {
	log.Printf("[WARN] %s", strings.Join(message, " "))
}

func Error(message ...string) {
	log.Printf("[ERROR] %s", strings.Join(message, " "))
}

func Info(message ...string) {
	log.Printf("[INFO] %s", strings.Join(message, " "))
}

func Fatal(message ...string) {
	log.Fatalf("[FATAL] %s ", strings.Join(message, " "))
}

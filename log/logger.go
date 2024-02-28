package log

import (
	"log"
)

func DEBUG(msg string) {
	log.Println("[DEBUG] " + msg)
}

func INFO(msg string) {
	log.Println("[INFO] " + msg)
}

func WARN(msg string) {
	log.Println("[WARN] " + msg)
}

func ERROR(msg string) {
	log.Println("[ERROR] " + msg)
}

func FATAL(msg string) {
	log.Println("[FATAL] " + msg)
	panic(msg)
}

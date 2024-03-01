package log

import (
	"fmt"
	"log"
)

func DEBUG(format string, v ...any) {
	log.Printf("[DEBUG] "+format, v...)
}

func INFO(format string, v ...any) {
	log.Printf("[INFO] "+format, v...)
}

func WARN(format string, v ...any) {
	log.Printf("[WARN] "+format, v...)
}

func ERROR(format string, v ...any) {
	log.Printf("[ERROR] "+format, v...)
}

func FATAL(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	log.Println("[FATAL] " + msg)
	panic(msg)
}

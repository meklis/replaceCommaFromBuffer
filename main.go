package main

import (
	"github.com/atotto/clipboard"
	"time"
	"regexp"
	"strings"
	"log"
)

func main() {
	for {
		text, _ := clipboard.ReadAll()
		if match, _ := regexp.MatchString(`^([0-9]{1,3}),([0-9]{1,3}),([0-9]{1,3}),([0-9]{1,3})$`, strings.TrimSpace(text)); match {
			new := strings.Replace(strings.TrimSpace(text), ",", ".", -1)
			clipboard.WriteAll(new)
			log.Printf("Found regex, replaced from '%v' to '%v'", text, new)
		}
		time.Sleep(time.Millisecond * 30)
	}
}

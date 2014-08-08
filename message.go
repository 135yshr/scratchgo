package scratchgo

import (
	"regexp"
	_ "strconv"
	"strings"
)

const regexParam = `\s+(?:\"((?:\\"|[^\"])*)\"?|([^\"]\S*))`

var (
	types      = []string{"broadcast", "sensor-update"}
	trimValues = []string{" ", "\""}
)

type Message struct {
	Type      string // sensor-update | broadcast
	Variables map[string]string
}

func ParseMessage(message string) *Message {
	type_name := strings.Split(message, " ")[0]
	message = message[len(type_name):]

	var parse func(string) map[string]string
	if type_name == "broadcast" {
		parse = parseBroadcast
	} else if type_name == "sensor-update" {
		parse = parseSensorupdate
	}

	return &Message{type_name, parse(message)}
}

func parseBroadcast(message string) map[string]string {
	cmd := trim(message, trimValues)
	return map[string]string{"command": cmd}
}

func parseSensorupdate(message string) map[string]string {
	ret := make(map[string]string)
	words := regexp.MustCompile(regexParam).FindAllString(message, -1)

	var key string
	for index, word := range words {
		word := trim(word, trimValues)
		if index%2 == 0 {
			key = word
			continue
		}

		ret[key] = word
	}
	return ret
}

func trim(text string, trm []string) string {
	ret := text
	for _, t := range trm {
		ret = strings.Trim(ret, t)
	}
	return ret
}

func (self *Message) GetNames() []string {
	ret := make([]string, len(self.Variables))
	var index int = 0
	for k, _ := range self.Variables {
		ret[index] = k
		index++
	}
	return ret
}

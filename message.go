package scratchgo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const regexParam = `\s+(?:\"((?:\\"|[^\"])*)\"?|([^\"]\S*))`

var (
	types      = []string{"broadcast", "sensor-update"}
	trimValues = []string{" ", "\""}
)

type Message struct {
	Type string // sensor-update | broadcast
	m    map[string]interface{}
}

func ParseMessage(message string) (*Message, error) {
	var m map[string]interface{}
	type_name := strings.Split(message, " ")[0]
	message = message[len(type_name):]

	if type_name == "broadcast" {
		m = parseBroadcast(message)
	} else if type_name == "sensor-update" {
		m = parseSensorupdate(message)
	} else {
		return nil, fmt.Errorf("un supported type name.[%s]", type_name)
	}

	return &Message{type_name, m}, nil
}

func parseBroadcast(message string) map[string]interface{} {
	key := trim(message, trimValues)
	return map[string]interface{}{key: nil}
}

func parseSensorupdate(message string) map[string]interface{} {
	words := regexp.MustCompile(regexParam).FindAllString(message, -1)
	ret := make(map[string]interface{})

	var key string
	for index, word := range words {
		word := trim(word, trimValues)
		if index%2 == 0 {
			key = word
			continue
		}

		if value, err := strconv.Atoi(word); err == nil {
			ret[key] = value
		} else if value, err := strconv.ParseFloat(word, 32); err == nil {
			ret[key] = value
		} else if value, err := strconv.ParseBool(word); err == nil {
			ret[key] = value
		} else {
			ret[key] = word
		}
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
	ret := make([]string, len(self.m))
	var index int = 0
	for k, _ := range self.m {
		ret[index] = k
		index++
	}
	return ret
}

func (self *Message) Len() int {
	return len(self.m)
}

func (self *Message) Get(key string) interface{} {
	return self.m[key]
}

func (self *Message) Set(key string, value interface{}) {
	self.m[key] = value
}

func isScratchCommand(t string) bool {
	for _, name := range types {
		if t == name {
			return true
		}
	}
	return false
}

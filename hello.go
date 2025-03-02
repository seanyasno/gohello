package gohello

import "errors"

var Greetings = map[string]string{
	"Chinese":  "你好",
	"English":  "Hello",
	"French":   "Bonjour",
	"German":   "Hallo",
	"Japanese": "こんにちは",
	"Korean":   "안녕하세요",
	"Spanish":  "Hola",
	"Russian":  "Привет",
}

func Hello(language string) (greeting string, err error) {
	gr, ok := Greetings[language]

	if ok {
		return gr, nil
	}

	return "", errors.New("Language not found")
}

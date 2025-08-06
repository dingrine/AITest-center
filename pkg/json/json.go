package json

import (
	"github.com/CrazyThursdayV50/pkgo/json"
	jsoniter "github.com/json-iterator/go"
)

var JSON = json.JSON

func Init() {
	jsoncfg := jsoniter.Config{
		IndentionStep:           2,
		MarshalFloatWith6Digits: false,
		EscapeHTML:              true,
		SortMapKeys:             true,
		UseNumber:               true,
		DisallowUnknownFields:   false,
		CaseSensitive:           true,
	}
	json.Init(&jsoncfg)
}

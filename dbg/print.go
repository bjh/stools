package dbg

import (
	"encoding/json"
	"fmt"
	"strings"
)

func PrintAsciiLine(lineChar string) {
	fmt.Println(strings.Repeat(lineChar, 60))
}

func PrettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")

	fmt.Println(string(s))
}

func PrettyPrintRaw(i interface{}) []byte {
	s, _ := json.MarshalIndent(i, "", "  ")

	return s
}

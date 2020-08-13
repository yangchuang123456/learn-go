package json

import (
	"fmt"
	"testing"
)
type test_json struct {
	A int    `json:"a"`
	B string `json:"b"`
}

func Test_parse_json(t *testing.T) {
	tmp := make(map[string]int,0)
	tmp["323"] = 1
	k,ok := tmp["323"]
	fmt.Println("the k and ok is:",k,ok)
}

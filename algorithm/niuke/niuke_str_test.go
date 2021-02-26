package niuke

import (
	"log"
	"strconv"
	"testing"
)

func decode(s string) string {
	flag := 0
	for _, c := range s {
		if c == '[' {
			break
		}
		flag++
	}
	if flag == len(s) {
		return s
	}

	j, p := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			j = i
			continue
		}
		if s[i] == ']' {
			p = i
			break
		}
	}

	if j >= p {
		return ""
	}

	tmp := s[j : p+1]

	charPos := 0
	for index,v := range tmp{
		if v == '|'{
			charPos = index
			break
		}
	}

	number,_:=strconv.Atoi(tmp[1:charPos])
	tmpLen := len(tmp)

	var midStr string
	for k := 0; k < number; k++ {
		midStr += tmp[charPos+1 : tmpLen-1]
	}

	s = s[:j] + midStr + s[p+1:]
	return decode(s)
}

func Test_decodeStr(t *testing.T) {
	res := decode("HG[3|B]F")
	log.Println("the res is:", res)

	res = decode("HG[3|B[2|CA]]F")
	log.Println("the res is:", res)
}

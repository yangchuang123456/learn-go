package algorithm

import (
	"log"
	"testing"
)

func decode(s string) string{

	flag := 0
	for _,c:=range s{
		if c=='['{
			break
		}
		flag++
	}
	if flag == len(s){
		return s
	}


	j,p:=0,0
	for i:=0;i<len(s);i++ {
		if s[i] == '['{
			j=i
			continue
		}
		if s[i] == ']'{
			p = i
			break
		}
	}

	if j >= p{
		return ""
	}


	tmp := s[j:p+1]
	tmpLen := len(tmp)
	number := int(tmp[1])-0x30

	//log.Println("the tmp and number is:",tmp ,number)

	var midStr string
	for k:=0;k<number;k++{
		midStr += tmp[3:tmpLen-1]
	}

	s = s[:j]+midStr+s[p+1:]

	return decode(s)
}

func Test_decodeStr(t *testing.T){
	res := decode("HG[3|B]F")
	log.Println("the res is:",res)

	res = decode("HG[3|B[2|CA]]F")
	log.Println("the res is:",res)
}

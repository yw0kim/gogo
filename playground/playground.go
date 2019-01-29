package main

import "fmt"

var (
	start = rune(44032) // "가"
	end   = rune(55204)
)

func aaa(s string) bool {
	numEnds := 28
	result := false
	for i, r := range s {
		fmt.Println(i, ":", r)
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0

		}
	}
	return result
}

func main() {
	s := "GO 언어"

	aaa(s)
}

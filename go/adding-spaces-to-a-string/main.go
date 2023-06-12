package main

import (
	"fmt"
	"strings"
)

func main() {
	s, spaces := "LeetcodeHelpsMeLearn", []int{8, 13, 15}
	fmt.Println(addSpaces(s, spaces))
}

/*

このプログラムを高速化するために、文字列の結合にはstrings.Builderを使用し、文字列の結合を効率的に行うことができます。
文字列の結合にstrings.Builderが使用されるため、効率的な文字列の結合が行われます。
また、不要な文字列のコピーが削減されるため、パフォーマンスが向上することが期待されます。

*/

func addSpaces(s string, spaces []int) string {
	var sb strings.Builder
	spacesIndex := 0

	for i := 0; i < len(s); i++ {
		if spacesIndex < len(spaces) && i == spaces[spacesIndex] {
			spacesIndex++
			sb.WriteByte(' ')
		}
		sb.WriteByte(s[i])
	}

	for spacesIndex < len(spaces) {
		sb.WriteByte(' ')
		spacesIndex++
	}

	return sb.String()
}

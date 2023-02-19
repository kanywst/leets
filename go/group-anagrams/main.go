package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		ss := strings.Join(s, "")

		if _, ok := m[ss]; ok {
			m[ss] = append(m[ss], str)
		} else {
			m[ss] = []string{str}
		}
	}
	i := 0
	ans := make([][]string, len(m))
	for _, v := range m {
		ans[i] = v
		i += 1
	}
	return ans
}

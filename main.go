package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := code("aaabbbcccccaa")
	fmt.Print(result)
}

func code(s string) string {
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)] += 1
	}

	keys := make([]string, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		if k == "" {
			continue
		}
		sb.WriteString(k)
		sb.WriteString(strconv.Itoa(m[k]))
	}
	return sb.String()
}
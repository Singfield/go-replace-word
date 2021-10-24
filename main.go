package main

import (
	"fmt"
	"strings"
)

// ProcessLine - searches for old in line to replace it by new...
// it returns found=true, if pattern was found, res with the resulting string
// and occ with the number of occurrences of old
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)

	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, oldLower, newLower, -1)
		res = strings.Replace(res, old, new, -1)
	}

	return found, res, occ
}

func main() {
	found, res, occ := ProcessLine("Go is very good i like programming in go", "Go", "Python")

	fmt.Println(found, res, occ)
}

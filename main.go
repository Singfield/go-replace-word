package main

import (
	"bufio"
	"fmt"
	"os"
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

func FindReplaceFile(src string, dst string, old string, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()
	distFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer distFile.Close()
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(distFile)
	defer writer.Flush()

	old = old + " "
	new = new + " "
	scanner = bufio.NewScanner(srcFile)
	lineIdx := 1
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		fmt.Fprintf(writer, res)
		lineIdx++
	}

	return occ, lines, nil
}

func main() {
	//found, res, occ := ProcessLine("Go is very good i like programming in go", "Go", "Python")

	//fmt.Println(found, res, occ)

	old := "Go"
	new := "python"

	occ, lines, err := FindReplaceFile("wikigo.txt", "wikipython.txt", old, new)
	if err != nil {
		fmt.Println("Error wile executing function", err)
	}
	fmt.Println("===Summary===")
	defer fmt.Println("==End==")
	fmt.Printf("Number of occurences of %v : %v\n", old, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))
	fmt.Print("lines: [")
	lens := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < lens-1 {
			fmt.Printf("-")
		}
	}
	fmt.Println("]")
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the contacts function below.
 */
func contacts(queries [][]string) []int32 {
	store := NewTrie()

	result := []int32{}
	for i := 0; i < len(queries); i++ {
		if queries[i][0] == "add" {
			store.insert(queries[i][1])
		} else if queries[i][0] == "find" {
			result = append(result, store.find(queries[i][1]))
		}

	}
	return result

}

type TrieNode struct {
	size     int32
	children map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

func (t *Trie) insert(s string) {
	current := t.root
	for _, v := range s {
		if _, ok := current.children[v]; !ok {
			current.children[v] = NewTrieNode()
		}
		current = current.children[v]
		current.size++
	}
}

func (t *Trie) find(prefix string) int32 {
	curr := t.root

	for _, v := range prefix {
		var ok bool
		curr, ok = curr.children[v]
		if !ok {
			return 0
		}
	}

	return curr.size
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	queriesRows, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries [][]string
	for queriesRowItr := 0; queriesRowItr < int(queriesRows); queriesRowItr++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []string
		for _, queriesRowItem := range queriesRowTemp {
			queriesRow = append(queriesRow, queriesRowItem)
		}

		if len(queriesRow) != int(2) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := contacts(queries)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

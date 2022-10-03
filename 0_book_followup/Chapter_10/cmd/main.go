package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// This programming style is not recommended for Go ;-)
func main() {
	f, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	uniques := make(map[string]int)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 0 {
			uniques[fields[0]]++
		}
	}
	type freq struct {
		ip    string
		count int
	}
	freqs := make([]freq, 0, len(uniques))
	for k, v := range uniques {
		freqs = append(freqs, freq{
			ip:    k,
			count: v,
		})
	}
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})
	fmt.Printf("%-16s %s\n", "Address", "Count")
	for i, j := range freqs {
		if i > 9 {
			break
		}
		fmt.Printf("%-16s %d\n", j.ip, j.count)
	}
}

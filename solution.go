package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"time"
)


type Word struct {
	bytes   []byte
	counter int64
}

func Solution() {
	start := time.Now()
	file, err := os.Open("mobydick.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	if err != nil {
		panic(err)
	}
	var words []*Word

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	ch := make(chan []byte)
	go func(){
		for scanner.Scan() {
			var tempWord []byte
			b := scanner.Bytes()[0]
			if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {
				if b >= 65 && b <= 90 {
					b += 32
				}
				tempWord = append(tempWord, b)
			} else {
				if len(tempWord) > 0 {
					ch <- tempWord
					tempWord = []byte{}
				}
			}
		}
		close(ch)
	}()
	for data := range ch {
		isFound := false
		for _, w := range words {
			if bytes.Equal(w.bytes, data) {
				w.counter++
				isFound = true
				break
			}
			if !isFound {
				words = append(words, &Word{
					bytes:   data,
					counter: 1,
				})
			}
		}
		sort.Slice(words, func(i, j int) bool {
			return words[i].counter > words[j].counter
		})
		for _, w := range words[:20] {
			fmt.Println(w.counter, string(w.bytes))
		}
		fmt.Printf("Time taken:  %s\n", time.Since(start))
	}
}

func main() {
	Solution()
}
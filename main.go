package main

import (
	"bufio"
	"fmt"
	pos "github.com/kamildrazkiewicz/go-stanford-nlp"
	"log"
	"os"
)

func main() {
	var (
		tagger *pos.Tagger
		res    []*pos.Result
		err    error
	)

	if tagger, err = pos.NewTagger(
		"models/english-left3words-distsim.tagger", // path to model
		"ext/stanford-postagger.jar"); err != nil { // path to jar tagger file
		fmt.Print(err)
		return
	}

	lines, err := readLines("sentences.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		// fmt.Println(i, line, "We got it")
		if res, err = tagger.Tag(line); err != nil {
			fmt.Print(err)
			return
		}
		for _, r := range res {
			fmt.Println(r.Word, r.TAG, r.Description())
		}
	}

	// if err := writeLines(lines, "foo.out.txt"); err != nil {
	// 	log.Fatalf("writeLines: %s", err)
	// }

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// func writeLines(lines []string, path string) error {
//     file, err := os.Create(path)
//     if err != nil {
//         return err
//     }
//     defer file.Close()

//     w := bufio.NewWriter(file)
//     for _, line := range lines {
//         fmt.Fprintln(w, line)
//     }
//     return w.Flush()
// }

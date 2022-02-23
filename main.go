package main

import (
	"bufio"
	"fmt"
	pos "github.com/kamildrazkiewicz/go-stanford-nlp"
	"log"
	"os"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var adjectiveSlice []string
var countSlice []float64
var adjSlice []string

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
	for i, line := range lines {
		fmt.Println(i, line)
		if res, err = tagger.Tag(line); err != nil {
			fmt.Print(err)
			return
		}
		for _, r := range res {
			// fmt.Println(r.Word, r.TAG, r.Description())
			if r.TAG == "JJ" {
				// fmt.Println(r.Word)
				adjectiveSlice = append(adjectiveSlice, r.Word)
				fmt.Println(adjectiveSlice)
			}
		}

	}
	dup_map := dup_count(adjectiveSlice)
	for k, v := range dup_map {
		fmt.Printf("Item : %s , Count : %f\n", k, v)
		adjSlice = append(adjSlice, k)
		countSlice = append(countSlice, v)
		graph(adjSlice, countSlice)
	}

	// if err := writeLines(lines, "foo.out.txt"); err != nil {
	// 	log.Fatalf("writeLines: %s", err)
	// }

}

func graph(k []string, v []float64) {
	groupA := plotter.Values{}
	adjectives := []string{}

	groupA = append(groupA, v...)

	p := plot.New()

	p.Title.Text = "Frequency of Adjectives in a Data Set"
	p.Y.Label.Text = "Frequency"
	p.X.Label.Text = "Range of Adjectives"

	w := vg.Points(4)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = -w

	p.Add(barsA)
	p.Legend.Add("Adjective", barsA)

	p.Legend.Top = true
	adjectives = append(adjectives, k...)
	p.NominalX(adjectives...)

	if err := p.Save(25*vg.Inch, 15*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}

}

func dup_count(list []string) map[string]float64 {

	duplicate_frequency := make(map[string]float64)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
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

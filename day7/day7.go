package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Bag reprents a node in the graph with 2 different types of relationships
// innerBags represent a list of bags that thie bag contains
// outterBags represent a list of bags that contains this bag
type Bag struct {
	name       string
	innerBags  []*Bag
	outterBags []*Bag
	counts     []int
}

func (b *Bag) addInnerBag(bag *Bag, count int) {
	b.innerBags = append(b.innerBags, bag)
	b.counts = append(b.counts, count)
}

func (b *Bag) addOutterBag(bag *Bag) {
	b.outterBags = append(b.outterBags, bag)
}

func (b Bag) countInnerBags() (count int) {
	if len(b.innerBags) == 0 {
		return
	}

	for i, bagPtr := range b.innerBags {
		n := b.counts[i]
		count += n
		count += (n * bagPtr.countInnerBags())
	}
	return
}

func (b Bag) getAllOutterBagIDs() (output []string) {
	if len(b.outterBags) == 0 {
		return
	}

	for _, b := range b.outterBags {
		output = append(output, b.name)
		for _, name := range b.getAllOutterBagIDs() {
			output = append(output, name)
		}
	}
	return
}

// BagGraph represents a graph of bags and their relationships to each other
type BagGraph map[string]*Bag

func (m BagGraph) getBag(id string) *Bag {
	if _, prs := m[id]; !prs {
		m[id] = &Bag{name: id}
	}
	return m[id]
}

func (m BagGraph) countValidOutterBags(bagID string) int {
	uniqueBagNames := map[string]bool{}
	for _, name := range m.getBag(bagID).getAllOutterBagIDs() {
		uniqueBagNames[name] = true
	}
	return len(uniqueBagNames)
}

func (m BagGraph) countInnerBags(bagID string) int {
	return m.getBag(bagID).countInnerBags()
}

func parseInput(path string) BagGraph {
	fp, err := os.Open("input")
	check(err)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	// Parse rules into graph
	bg := BagGraph{}
	for scanner.Scan() {
		// words = ["xx xxx bags ", "n1 xx xxx bag/s, n2 xx xxx bag/s"]
		words := strings.Split(scanner.Text(), "contain ")
		outterBagName := words[0][:len(words[0])-6]
		outterBag := bg.getBag(outterBagName)

		if words[1] == "no other bags." {
			continue
		}

		// split inner bags into ["n1 xxx xxx bag/s", "n2 xxx xxx bag/s."]
		for _, word := range strings.Split(words[1], ", ") {
			// split each inner bag into something like ["3", "bright", "white", "bags.."]
			innerBagWords := strings.Split(word, " ")
			n, err := strconv.Atoi(innerBagWords[0])
			check(err)
			innerBagName := strings.Join(innerBagWords[1:3], " ")
			innerBag := bg.getBag(innerBagName)
			innerBag.addOutterBag(outterBag)
			outterBag.addInnerBag(innerBag, n)
		}
	}
	return bg
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	bags := parseInput("input")

	// Part 1
	fmt.Println(bags.countValidOutterBags("shiny gold"))

	// Part 2
	fmt.Println(bags.countInnerBags("shiny gold"))
}

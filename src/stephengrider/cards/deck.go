package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// di dalam Go, yg mempunyai receiver biasanya di namain dgn single letter contoh: a, b, c, d
func (d deck) print() { // deck is receiver
	for i, card := range d { // array d or cards bisa coba diganti dengan this or self
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{"Ace of Spades", "Two of Spades"}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // use underscore when u didnt use this value
		for _, value := range cardValues { // use underscore when u didnt use this value
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the err and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

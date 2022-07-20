package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Slice of strings
type deck []string

func new_deck() deck {
	result := deck{}
	suits := []string{"clubs", "diamonds", "hearts", "spades"}
	counts := []string{"one", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, count := range counts {
			result = append(result, count+" of "+suit)
		}
	}
	return result
}

func (xd deck) print_cards() {
	for i, card := range xd {
		fmt.Println(i, card)
	}
}

func deal(d deck, hand_size int) (deck, deck) {
	return d[:hand_size], d[hand_size:]
}

func (d deck) to_bytes() []byte {
	return []byte(strings.Join([]string(d), "\n"))
}

func (d deck) save_to_file(filename string) error {
	return ioutil.WriteFile(filename, d.to_bytes(), 0666)
}

func read_from_file(filename string) deck {
	d, err := ioutil.ReadFile(filename)
	if err == nil {
		return deck(strings.Split(string(d), "\n"))
	} else {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return nil
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano()) 
	r := rand.New(src)

	for i := range d {
		new_pos := r.Intn(len(d) - 1)
		d[i], d[new_pos] = d[new_pos], d[i]
	}
}

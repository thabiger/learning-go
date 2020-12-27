package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	switch card, suit := pickCard(); { //missing expression means true
	/*
	   i know this comparison is naive and wouldn't work
	   to actually compare the cards, but i leave it this
	   to be fixed in the future as i'll collect necessary
	   knowledge to do that
	*/
	case card > "10":
		fmt.Println("You've pick a high card:", card, "of", suit)
	case card < "5":
		fmt.Println("You've pick a low card:", card, "of", suit)
	default:
		fmt.Println("You've pick a card:", card, "of", suit)
	}

}

func pickCard() (string, string) {
	cards := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}

	rand.Seed(time.Now().Unix())

	return cards[rand.Intn(len(cards))], suits[rand.Intn(len(suits))]
}

package main

func main() {
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()

	/** Deck to string */
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards.txt")
	// cards := newDeckFromFile("my_card.txt")
	// cards.print()

	/** SHUFFLE */
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }

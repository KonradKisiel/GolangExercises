package main

//mport "fmt"

func main() {
	/*
		cards := deck{"Ace of Diamonds", newCard()}
		cards = append(cards, "Six of Spades")
	*/
	//iteration over slice
	//receiver function
	cards := newDeck()
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	/*
		hand, remainingDeck := deal(cards, 5)

		hand.print()
		fmt.Println("\n")
		remainingDeck.print()
	*/
}

/*
func newCard() string {
	return "Five of Diamonds"
}
*/

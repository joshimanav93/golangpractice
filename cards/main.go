package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// package main

// import "fmt"

// func main() {
// 	// cards := deck{"Ace of diamonds", newCard()}
// 	// cards = append(cards, "Six of spades")

// 	// cards.print()
// 	cards := newDeck()
// 	//fmt.Println(cards)
// 	cards.shuffle()
// 	cards.saveToFile("my_cards")
// 	cards.newDeckFromFile("my_cards")
// 	fmt.Println(cards.toString())
// 	//cards.print()
// 	// hand, remainingCards := deal(cards, 5)
// 	// hand.print()
// 	// remainingCards.print()
// }

// //t.Errorf("Expected deck length of 16, but got %v", len(d))
// // func newCard() string {
// // 	return "Ace of spades"
// // }

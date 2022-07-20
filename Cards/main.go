package main

func main() {

	//cards := new_deck()

	// //cards.print_cards()

	// a, b := deal(cards, 3)
	// a.print_cards()
	// fmt.Println("\n\n")
	// b.print_cards()
	//cards.save_to_file("Test.txt")
	new_cards := read_from_file("Test.txt")
	new_cards.print_cards()
}

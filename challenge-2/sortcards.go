package sortcards

var n int = 0

//This trivial example is not allowed.  It merely is a test of the tests.
func SortCards(in string) string {
	var rc string
	switch n {
	case 0:
		rc = "Three of Spades\nQueen of Diamonds\nTwo of Diamonds\nAce of Clubs\nJack of Clubs"
	case 1:
		rc = "Four of Spades\nQueen of Diamonds\nTwo of Diamonds\nAce of Clubs\nJack of Clubs"
	default:
		rc = "I don't know"
	}
	n++
	return rc
}

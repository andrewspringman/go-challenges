package sortcards

import "testing"

var tests = []struct {
	input  string
	output string
}{
	{
		"Three of Spades\nJack of Clubs\nAce of Clubs\nTwo of Diamonds\nQueen of Diamonds",
		"Three of Spades\nQueen of Diamonds\nTwo of Diamonds\nAce of Clubs\nJack of Clubs",
	},
	{
		"Four of Spades\nJack of Clubs\nAce of Clubs\nTwo of Diamonds\nQueen of Diamonds",
		"Four of Spades\nQueen of Diamonds\nTwo of Diamonds\nAce of Clubs\nJack of Clubs",
	},
}

func TestSortCards(t *testing.T) {
	for _, test := range tests {
		out := SortCards(test.input)
		if out != test.output {
			t.Errorf("Expected:\n\n%s\n\nbut received:\n\n%s\n\n", test.output, out)
		}
	}
}

func BenchmarkSortCards(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			SortCards(test.input)
		}
	}
}

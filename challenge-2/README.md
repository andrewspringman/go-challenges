# Challenge 2 - Sort Cards
Demonstrates mastery of the string manipulation and custom sorting

Create a function with the signature
```
func SortCards(in string) string
```
that when given a newline separated list of card names, returns a sorted newline separate list of the cards (highest value to lowest).  Card names are in the following format (where capitalization and spacing matters).
```
Three of Spades
Jack of Clubs
Ace of Clubs
Two of Diamonds
Queen of Diamons
```
Ace is high.  The suit order from highest to lowest is Spades, Hearts, Diamonds, Clubs.
Two is called Two, Jack is called Jack, and there is no Joker.

For example, given the list above, the result should be.
```
Three of Spades
Queen of Diamons
Two of Diamonds
Ace of Clubs
Jack of Clubs
```
The trivial example in sortcards.go is not allowed.  You've actually got to do the sort.

### Run tests with benchmarks

```
go test -bench .
```

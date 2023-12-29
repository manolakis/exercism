package blackjack

const (
	STAND = "S"
	HIT   = "H"
	SPLIT = "P"
	WIN   = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int

	switch card {
	case "ace":
		value = 11
	case "ten", "jack", "queen", "king":
		value = 10
	case "nine":
		value = 9
	case "eight":
		value = 8
	case "seven":
		value = 7
	case "six":
		value = 6
	case "five":
		value = 5
	case "four":
		value = 4
	case "three":
		value = 3
	case "two":
		value = 2
	case "one":
		value = 1
	default:
		value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return SPLIT
	}

	dealerCardValue := ParseCard(dealerCard)

	switch analyzeCards(card1, card2) {
	case "blackjack":
		if dealerCardValue < 10 {
			return WIN
		} else {
			return STAND
		}
	case "17+":
		return STAND
	case "12+":
		if dealerCardValue >= 7 {
			return HIT
		}
		return STAND
	default:
		return HIT
	}

}

func analyzeCards(card1, card2 string) string {
	sum := ParseCard(card1) + ParseCard(card2)

	if sum == 21 {
		return "blackjack"
	}

	if sum >= 17 {
		return "17+"
	}

	if sum >= 12 {
		return "12+"
	}

	return "other"
}

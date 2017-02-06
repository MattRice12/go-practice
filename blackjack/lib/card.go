package lib

import "strconv"

// Card holds attributes for a card
type Card struct {
	Suit  string
	Value string
}

// FixFace sets 11+ cards to J, Q, K, A
func FixFace(v int) string {
	if v == 14 {
		return "A"
	} else if v == 13 {
		return "K"
	} else if v == 12 {
		return "Q"
	} else if v == 11 {
		return "J"
	}
	return strconv.Itoa(v)
}

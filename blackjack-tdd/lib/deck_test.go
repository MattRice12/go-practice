package lib

import "testing"

func TestCreateDeck(t *testing.T) {
	d := CreateDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 cards; got %d cards.", len(d))
	}
	if d[0].Suit != "Hearts" || d[0].Value == 0 {
		t.Errorf("Expected 2 of Hearts; got %d of %v", d[0].Value, d[0].Suit)
	}
}

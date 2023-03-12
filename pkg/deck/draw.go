package deck

import "fmt"

func (d *Deck) DrawCard(count int) ([]Card, error) {
	if len(d.Cards) == 0 {
		return nil, fmt.Errorf("sorry! the deck is empty")
	}
	if len(d.Cards) < count {
		return nil, fmt.Errorf("insufficient cards! cannot draw %d cards from a deck of %d cards", count, len(d.Cards))
	}

	cards := d.Cards[:count]
	d.Cards = d.Cards[count:]

	return cards, nil
}

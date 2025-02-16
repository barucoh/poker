package poker

import (
	"fmt"

	"github.com/barucoh/poker/random"
)

const MaxRetries = 5

var fullDeck *Deck

func init() {
	fullDeck = &Deck{initializeFullCards()}
}

type Deck struct {
	cards []Card
}

func NewDeck() (*Deck, error) {
	deck := &Deck{}
	var err error

	for i := 0; i < MaxRetries; i++ {
		err = deck.Shuffle()
		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to shuffle deck. Error: %w", err)
	}

	return deck, nil
}

func (deck *Deck) Shuffle() error {
	deck.cards = make([]Card, len(fullDeck.cards))
	copy(deck.cards, fullDeck.cards)

	err := random.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})

	if err != nil {
		return fmt.Errorf("error in deck generation: %w", err)
	}

	return nil
}

func (deck *Deck) Draw(n int) []Card {
	cards := make([]Card, n)
	copy(cards, deck.cards[:n])
	deck.cards = deck.cards[n:]
	return cards
}

func (deck *Deck) Empty() bool {
	return len(deck.cards) == 0
}

func initializeFullCards() []Card {
	var cards []Card

	for _, rank := range strRanks {
		for suit := range charSuitToIntSuit {
			cards = append(cards, NewCard(string(rank)+string(suit)))
		}
	}

	return cards
}

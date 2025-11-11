/*
 * Poker Straight Flush Probability Calculator
 * 
 * This program uses Monte Carlo simulation to calculate the probability
 * of getting a straight flush in 7-card stud poker.
 * 
 * A straight flush is 5 consecutive cards of the same suit.
 * The program shuffles a deck, deals multiple 7-card hands, and checks
 * each hand for straight flushes. Results are compared against published
 * poker probability tables (approximately 0.0279% or about 1 in 3,590).
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit int

const (
	club Suit = iota
	diamond
	heart
	spade
)

type Card struct {
	s   Suit
	pip int
}

// shuffle randomizes the deck using Fisher-Yates algorithm
func shuffle(d *[52]Card) {
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

// isStraightFlush checks if a 7-card hand contains a straight flush
// First checks for flush (5+ cards of same suit), then checks for straight
func isStraightFlush(h []Card) bool {
	// Group cards by suit
	clubCards := []int{}
	diamondCards := []int{}
	heartCards := []int{}
	spadeCards := []int{}

	for _, card := range h {
		switch card.s {
		case club:
			clubCards = append(clubCards, card.pip)
		case diamond:
			diamondCards = append(diamondCards, card.pip)
		case heart:
			heartCards = append(heartCards, card.pip)
		case spade:
			spadeCards = append(spadeCards, card.pip)
		}
	}

	// Check each suit that has 5 or more cards for a straight
	if len(clubCards) >= 5 && hasStraight(clubCards) {
		return true
	}
	if len(diamondCards) >= 5 && hasStraight(diamondCards) {
		return true
	}
	if len(heartCards) >= 5 && hasStraight(heartCards) {
		return true
	}
	if len(spadeCards) >= 5 && hasStraight(spadeCards) {
		return true
	}

	return false
}

// hasStraight checks if a set of pip values contains a 5-card straight
// Handles special case where Ace can be low (A-2-3-4-5) or high (10-J-Q-K-A)
func hasStraight(pips []int) bool {
	if len(pips) < 5 {
		return false
	}

	// Create a boolean array to mark which cards we have (1-13)
	present := make([]bool, 14) // Index 0 unused, 1-13 for cards
	for _, pip := range pips {
		present[pip] = true
	}

	// Check for 5 consecutive cards (including Ace-low straight)
	// Ace-low straight: A-2-3-4-5
	if present[1] && present[2] && present[3] && present[4] && present[5] {
		return true
	}

	// Check for regular straights (2-3-4-5-6 through 10-J-Q-K-A)
	for i := 2; i <= 9; i++ {
		if present[i] && present[i+1] && present[i+2] && present[i+3] && present[i+4] {
			return true
		}
	}

	// Ace-high straight: 10-J-Q-K-A
	if present[10] && present[11] && present[12] && present[13] && present[1] {
		return true
	}

	return false
}

// printCard displays a card in human-readable format
func (c Card) printCard() {
	var suitName, pipName string

	switch c.s {
	case club:
		suitName = "Clubs"
	case diamond:
		suitName = "Diamonds"
	case heart:
		suitName = "Hearts"
	case spade:
		suitName = "Spades"
	}

	switch c.pip {
	case 1:
		pipName = "Ace"
	case 2:
		pipName = "Two"
	case 3:
		pipName = "Three"
	case 4:
		pipName = "Four"
	case 5:
		pipName = "Five"
	case 6:
		pipName = "Six"
	case 7:
		pipName = "Seven"
	case 8:
		pipName = "Eight"
	case 9:
		pipName = "Nine"
	case 10:
		pipName = "Ten"
	case 11:
		pipName = "Jack"
	case 12:
		pipName = "Queen"
	case 13:
		pipName = "King"
	}

	fmt.Printf("%s of %s  ", pipName, suitName)
}

// initializeDeck creates a standard 52-card deck
func initializeDeck() [52]Card {
	var deck [52]Card

	// Set pip values (1-13, where 1=Ace, 11=Jack, 12=Queen, 13=King)
	for i := 0; i < 52; i++ {
		deck[i].pip = (i % 13) + 1
	}

	// Set suits (13 cards per suit)
	for i := 0; i < 13; i++ {
		deck[i].s = club
	}
	for i := 13; i < 26; i++ {
		deck[i].s = diamond
	}
	for i := 26; i < 39; i++ {
		deck[i].s = heart
	}
	for i := 39; i < 52; i++ {
		deck[i].s = spade
	}

	return deck
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Poker Straight Flush Probability Calculator ===")
	fmt.Println("Calculating probability of straight flush in 7-card stud poker")
	fmt.Println()

	var totalTrials int
	fmt.Print("Enter number of trials (recommend 100000 or more): ")
	fmt.Scanf("%d", &totalTrials)

	deck := initializeDeck()
	straightFlushCount := 0

	// We can deal 7 hands per shuffle (52 / 7 = 7 with 3 cards left)
	// Why not 8 hands? Because 8 * 7 = 56 > 52 (not enough cards in deck)
	handsPerShuffle := 7

	fmt.Println("\nRunning simulation...")

	// Run trials
	for trial := 0; trial < totalTrials; trial++ {
		shuffle(&deck)

		// Deal 7 hands from the shuffled deck
		for handNum := 0; handNum < handsPerShuffle; handNum++ {
			hand := make([]Card, 7)
			startIdx := handNum * 7

			// Deal 7 cards for this hand
			for i := 0; i < 7; i++ {
				hand[i] = deck[startIdx+i]
			}

			// Check for straight flush
			if isStraightFlush(hand) {
				straightFlushCount++

				// Print every 100th straight flush found
				if straightFlushCount%100 == 0 {
					fmt.Printf("\nStraight Flush #%d found:\n", straightFlushCount)
					for _, card := range hand {
						card.printCard()
					}
					fmt.Println()
				}
			}
		}

		// Progress indicator for long runs
		if (trial+1)%10000 == 0 {
			fmt.Printf("Completed %d trials...\n", (trial+1)*handsPerShuffle)
		}
	}

	// Calculate results
	totalHands := totalTrials * handsPerShuffle
	probability := float64(straightFlushCount) / float64(totalHands) * 100

	fmt.Println("\n=== Results ===")
	fmt.Printf("Total hands dealt: %d\n", totalHands)
	fmt.Printf("Straight flushes found: %d\n", straightFlushCount)
	fmt.Printf("Probability: %.4f%%\n", probability)
	fmt.Printf("Odds: approximately 1 in %.0f\n", float64(totalHands)/float64(straightFlushCount))
	fmt.Println("\nPublished probability: ~0.0279% (approximately 1 in 3,590)")
}
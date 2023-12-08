package main

import (
	"log"
	"sort"
	"strings"
)

func day7b(c chan string) {
	var total int64
	hands := initializeDay7(c)
	for h := range hands {
		cards := []rune(hands[h].cards)
		sort.Slice(cards, func(i, j int) bool {
			return cards[i] < cards[j]
		})
		// 5 of a kind
		if cards[0] == cards[1] && cards[1] == cards[2] && cards[2] == cards[3] && cards[3] == cards[4] {
			hands[h].cards = "7" + hands[h].cards
			continue
		}

		// 4 of a kind
		if cards[1] == cards[2] && cards[2] == cards[3] && (cards[3] == cards[4] || cards[0] == cards[1]) {
			if strings.Contains(string(cards), "J") {
				hands[h].cards = "7" + hands[h].cards
			} else {
				hands[h].cards = "6" + hands[h].cards
			}
			continue
		}
		// full house
		if (cards[0] == cards[1] && cards[1] == cards[2] && cards[3] == cards[4]) || (cards[0] == cards[1] && cards[2] == cards[3] && cards[3] == cards[4]) {
			if strings.Contains(string(cards), "J") {
				hands[h].cards = "7" + hands[h].cards
			} else {
				hands[h].cards = "5" + hands[h].cards
			}
			continue
		}
		// 3 of a kind
		if (cards[0] == cards[1] && cards[1] == cards[2]) || (cards[2] == cards[3] && cards[3] == cards[4]) || (cards[1] == cards[2] && cards[2] == cards[3]) {
			if strings.Contains(string(cards), "JJJ") {
				hands[h].cards = "6" + hands[h].cards
			} else if strings.Contains(string(cards), "J") {
				hands[h].cards = "6" + hands[h].cards
			} else {
				hands[h].cards = "4" + hands[h].cards
			}
			continue
		}

		pairs := 0
		for i := 1; i < len(cards); i++ {
			if cards[i] == cards[i-1] {
				pairs++
			}
		}
		switch pairs {
		// two pairs
		case 2:
			if strings.Contains(string(cards), "JJ") {
				hands[h].cards = "6" + hands[h].cards
			} else if strings.Contains(string(cards), "J") {
				hands[h].cards = "5" + hands[h].cards
			} else {
				hands[h].cards = "3" + hands[h].cards
			}
		// one pair
		case 1:
			if strings.Contains(string(cards), "J") {
				hands[h].cards = "4" + hands[h].cards
			} else {
				hands[h].cards = "2" + hands[h].cards
			}

		// High card
		case 0:
			if strings.Contains(string(cards), "J") {
				hands[h].cards = "2" + hands[h].cards
			} else {
				hands[h].cards = "1" + hands[h].cards
			}
		}
	}
	for h := range hands {
		cardsStr := strings.ReplaceAll(hands[h].cards, "A", "E")
		cardsStr = strings.ReplaceAll(cardsStr, "K", "D")
		cardsStr = strings.ReplaceAll(cardsStr, "Q", "C")
		cardsStr = strings.ReplaceAll(cardsStr, "J", "0")
		cardsStr = strings.ReplaceAll(cardsStr, "T", "A")
		hands[h].cards = cardsStr
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].cards < hands[j].cards
	})
	for i, hand := range hands {
		log.Println(hand.cards)
		total += int64(hand.bid * (i + 1))
	}
	log.Printf("7B Total: %d", total)

}

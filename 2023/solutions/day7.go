package year2023

import (
	"bartwestenenk/aoc/utils"
	"fmt"
	"golang.org/x/exp/slices"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func GetDay7() [2]func(input string) int {
	return [2]func(input string) int{
		challenge7a,
		challenge7b,
	}
}

type hand struct {
	cards    string
	cardsMap map[string]int
	kinds    []int
	bid      int
}

func (h hand) is5ofaKind() bool {
	return len(h.kinds) == 1
}

func (h hand) is4ofaKind() bool {
	return len(h.kinds) == 2 && h.kinds[1] == 1 && h.kinds[0] == 4
}

func (h hand) isFullHouse() bool {
	return len(h.kinds) == 2 && h.kinds[1] == 2 && h.kinds[0] == 3
}

func (h hand) is3ofaKind() bool {
	return len(h.kinds) == 3 && h.kinds[2] == 1 && h.kinds[1] == 1 && h.kinds[0] == 3
}

func (h hand) is2pair() bool {
	return len(h.kinds) == 3 && h.kinds[2] == 1 && h.kinds[1] == 2 && h.kinds[0] == 2
}

func (h hand) ispair() bool {
	return len(h.kinds) == 4 && h.kinds[3] == 1 && h.kinds[2] == 1 && h.kinds[1] == 1 && h.kinds[0] == 2
}

func (h hand) isHighcard() bool {
	return len(h.kinds) == 5
}

func (h hand) getComboScore() int {
	switch {
	case h.is5ofaKind():
		return 7
	case h.is4ofaKind():
		return 6
	case h.isFullHouse():
		return 5
	case h.is3ofaKind():
		return 4
	case h.is2pair():
		return 3
	case h.ispair():
		return 2
	case h.isHighcard():
		return 1
	}
	return 0
}

func (h1 hand) Less(h2 hand) bool {
	comboScore1 := h1.getComboScore()
	comboScore2 := h2.getComboScore()

	// Decide first based on the comboScore
	if comboScore1 != comboScore2 {
		return comboScore1 < comboScore2
	}

	// Decide on based on the card values
	for i := 0; i < len(h1.cards); i++ {
		c1 := h1.cards[i]
		c2 := h2.cards[i]

		// If they are the same, check the next card
		if c1 == c2 {
			continue
		}

		// If one is digit and the other is a letter, letter wins
		if unicode.IsDigit(rune(c1)) == unicode.IsLetter(rune(c2)) {
			return !unicode.IsLetter(rune(c1))
		}

		// Use int value to compare if they are both digits
		if unicode.IsDigit(rune(c1)) {
			i1, _ := strconv.Atoi(string(c1))
			i2, _ := strconv.Atoi(string(c2))

			return i1 < i2
		}

		convMap := map[rune]int{
			'A': 5,
			'K': 4,
			'Q': 3,
			'J': 2,
			'T': 1,
		}

		// Use letter conversion and compare
		if unicode.IsLetter(rune(c1)) {
			return convMap[rune(c1)] < convMap[rune(c2)]
		}
	}
	return false
}

func challenge7a(input string) int {
	lines := utils.SplitLines(input)
	hands := []hand{}
	answer := 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		bid, _ := strconv.Atoi(splitLine[1])
		cardsMap := make(map[string]int)
		var kinds []int
		for _, amount := range cardsMap {
			kinds = append(kinds, amount)
		}
		slices.Sort(kinds)
		slices.Reverse(kinds)

		hands = append(hands, hand{
			cards:    splitLine[0],
			cardsMap: cardsMap,
			kinds:    kinds,
			bid:      bid,
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Less(hands[j])
	})

	for _, h := range hands {
		fmt.Printf("%v: %v - %v\n", h.cards, h.getComboScore(), h.kinds)
	}

	for i, h := range hands {
		answer += (i + 1) * h.bid
	}

	return answer
}

func challenge7b(input string) int {
	answer := 0
	return answer
}

package main

import (
	"aoc2023/pkg/files"
	"slices"
	"strconv"
	"strings"
)

const Data = "data/day07"

var card_values = map[rune]int{
	'*': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

const (
	HighCard     = 1
	OnePair      = 2
	TwoPairs     = 3
	ThreeOfAKind = 4
	FullHouse    = 5
	FourOfAKind  = 6
	FiveOfAKind  = 7
)

type Hand struct {
	cards   []rune
	bid     int
	scoring int
}

type CardCount struct {
	card  rune
	value int
	count int
}

func getCardCounts(cards []rune) []CardCount {
	tmp := make(map[rune]int)
	for _, c := range cards {
		tmp[c]++
	}
	res := make([]CardCount, 0)
	for k, v := range tmp {
		res = append(res, CardCount{card: k, count: v, value: card_values[k]})
	}
	return res
}

func getCardScoring(cards []rune) (int, []rune) {
	scoring := HighCard
	counts := getCardCounts(cards)
	jokers := 0
	for _, count := range counts {
		if count.card == '*' {
			jokers = count.count
			continue
		}
		if count.count == 5 {
			scoring = FiveOfAKind
		} else if count.count == 4 {
			scoring = FourOfAKind
		} else if count.count == 3 {
			if scoring == OnePair {
				scoring = FullHouse
			} else {
				scoring = ThreeOfAKind
			}
		} else if count.count == 2 {
			if scoring == OnePair {
				scoring = TwoPairs
			} else if scoring == ThreeOfAKind {
				scoring = FullHouse
			} else {
				scoring = OnePair
			}
		}
	}

	if jokers > 0 {
		switch scoring {
		case HighCard:
			switch jokers {
			case 1:
				scoring = OnePair
			case 2:
				scoring = ThreeOfAKind
			case 3:
				scoring = FourOfAKind
			case 4, 5:
				scoring = FiveOfAKind
			}
		case OnePair:
			switch jokers {
			case 1:
				scoring = ThreeOfAKind
			case 2:
				scoring = FourOfAKind
			case 3, 4, 5:
				scoring = FiveOfAKind
			}
		case TwoPairs:
			scoring = FullHouse
		case ThreeOfAKind:
			if jokers == 1 {
				scoring = FourOfAKind
			} else {
				scoring = FiveOfAKind
			}
		case FourOfAKind:
			scoring = FiveOfAKind
		}
	}
	slices.SortFunc(counts, func(a, b CardCount) int {
		if a.count == b.count {
			return b.value - a.value
		}
		return b.count - a.count
	})
	/*
		// I outsmarted myself… why is Camel Cards not proper Poker?
		hand := make([]rune, 0)
		for _, count := range counts {
			for i := 0; i < count.count; i++ {
				hand = append(hand, count.card)
			}
		}
		return scoring, hand
	*/
	return scoring, cards
}

func solve1(file string) int {
	res := 0
	hands := make([]Hand, 0)
	if err := files.ReadLines(file, func(line string) {
		tmp := strings.SplitN(line, " ", 2)
		_b, _ := strconv.Atoi(tmp[1])
		score, hand := getCardScoring([]rune(tmp[0]))
		hands = append(hands, Hand{cards: hand, bid: _b, scoring: score})
	}); err != nil {
		panic(err)
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.scoring == b.scoring {
			for i := 0; i < len(a.cards); i++ {
				if a.cards[i] != b.cards[i] {
					return card_values[a.cards[i]] - card_values[b.cards[i]]
				}
			}
		}
		return a.scoring - b.scoring
	})
	for i, hand := range hands {
		res += (i + 1) * hand.bid
		// println(fmt.Sprintf("%d: %s", hand.scoring, string(hand.cards)))
	}
	return res
}

func solve2(file string) int {
	res := 0
	hands := make([]Hand, 0)
	if err := files.ReadLines(file, func(line string) {
		tmp := strings.SplitN(line, " ", 2)
		_b, _ := strconv.Atoi(tmp[1])
		score, hand := getCardScoring([]rune(strings.ReplaceAll(tmp[0], "J", "*")))
		hands = append(hands, Hand{cards: hand, bid: _b, scoring: score})
	}); err != nil {
		panic(err)
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.scoring == b.scoring {
			for i := 0; i < len(a.cards); i++ {
				if a.cards[i] != b.cards[i] {
					return card_values[a.cards[i]] - card_values[b.cards[i]]
				}
			}
		}
		return a.scoring - b.scoring
	})
	for i, hand := range hands {
		res += (i + 1) * hand.bid
		// println(fmt.Sprintf("%d: %s", hand.scoring, string(hand.cards)))
	}
	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 1: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}

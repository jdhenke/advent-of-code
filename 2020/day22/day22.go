package day22

import (
	"io"
	"regexp"
	"sort"
	"strconv"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 22: Crab Combat ---
It only takes a few hours of sailing the ocean on a raft for boredom to sink
in. Fortunately, you brought a small deck of space cards! You'd like to play a
game of Combat, and there's even an opponent available: a small crab that
climbed aboard your raft before you left.

Fortunately, it doesn't take long to teach the crab the rules.

Before the game starts, split the cards so each player has their own deck (your
puzzle input). Then, the game consists of a series of rounds: both players draw
their top card, and the player with the higher-valued card wins the round. The
winner keeps both cards, placing them on the bottom of their own deck so that
the winner's card is above the other card. If this causes a player to have all
of the cards, they win, and the game ends.

For example, consider the following starting decks:

    Player 1:
    9
    2
    6
    3
    1

    Player 2:
    5
    8
    4
    7
    10

This arrangement means that player 1's deck contains 5 cards, with 9 on top and
1 on the bottom; player 2's deck also contains 5 cards, with 5 on top and 10 on
the bottom.

The first round begins with both players drawing the top card of their decks: 9
and 5. Player 1 has the higher card, so both cards move to the bottom of player
1's deck such that 9 is above 5. In total, it takes 29 rounds before a player
has all of the cards:

    -- Round 1 --
    Player 1's deck: 9, 2, 6, 3, 1
    Player 2's deck: 5, 8, 4, 7, 10
    Player 1 plays: 9
    Player 2 plays: 5
    Player 1 wins the round!

    -- Round 2 --
    Player 1's deck: 2, 6, 3, 1, 9, 5
    Player 2's deck: 8, 4, 7, 10
    Player 1 plays: 2
    Player 2 plays: 8
    Player 2 wins the round!

    -- Round 3 --
    Player 1's deck: 6, 3, 1, 9, 5
    Player 2's deck: 4, 7, 10, 8, 2
    Player 1 plays: 6
    Player 2 plays: 4
    Player 1 wins the round!

    -- Round 4 --
    Player 1's deck: 3, 1, 9, 5, 6, 4
    Player 2's deck: 7, 10, 8, 2
    Player 1 plays: 3
    Player 2 plays: 7
    Player 2 wins the round!

    -- Round 5 --
    Player 1's deck: 1, 9, 5, 6, 4
    Player 2's deck: 10, 8, 2, 7, 3
    Player 1 plays: 1
    Player 2 plays: 10
    Player 2 wins the round!

    ...several more rounds pass...

    -- Round 27 --
    Player 1's deck: 5, 4, 1
    Player 2's deck: 8, 9, 7, 3, 2, 10, 6
    Player 1 plays: 5
    Player 2 plays: 8
    Player 2 wins the round!

    -- Round 28 --
    Player 1's deck: 4, 1
    Player 2's deck: 9, 7, 3, 2, 10, 6, 8, 5
    Player 1 plays: 4
    Player 2 plays: 9
    Player 2 wins the round!

    -- Round 29 --
    Player 1's deck: 1
    Player 2's deck: 7, 3, 2, 10, 6, 8, 5, 9, 4
    Player 1 plays: 1
    Player 2 plays: 7
    Player 2 wins the round!

    == Post-game results ==
    Player 1's deck:
    Player 2's deck: 3, 2, 10, 6, 8, 5, 9, 4, 7, 1

Once the game ends, you can calculate the winning player's score. The bottom
card in their deck is worth the value of the card multiplied by 1, the
second-from-the-bottom card is worth the value of the card multiplied by 2, and
so on. With 10 cards, the top card is worth the value on the card multiplied by
10. In this example, the winning player's score is:

       3 * 10
    +  2 *  9
    + 10 *  8
    +  6 *  7
    +  8 *  6
    +  5 *  5
    +  9 *  4
    +  4 *  3
    +  7 *  2
    +  1 *  1
    = 306

So, once the game ends, the winning player's score is 306.

Play the small crab in a game of Combat using the two decks you just dealt.
What is the winning player's score?
*/
func Part1(r io.Reader) (answer int, err error) {
	return day22(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day22(r)
}

var re = regexp.MustCompile(`Player (\d+):`)

func day22(r io.Reader) (answer int, err error) {
	// parse cards
	hands := make(map[string][]int)
	{
		var player string
		if err := input.ForEachLine(r, func(line string) error {
			if match := re.FindStringSubmatch(line); match != nil {
				player = match[1]
				return nil
			}
			if line == "" {
				return nil
			}
			x, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			hands[player] = append(hands[player], x)
			return nil
		}); err != nil {
			return 0, err
		}
	}

	// play
	for {
		round := make(map[string]int)
		for player, hand := range hands {
			if len(hand) == 0 {
				continue
			}
			round[player] = hand[0]
			hands[player] = hand[1:]
		}
		winner, cards := judge(round)
		hands[winner] = append(hands[winner], cards...)
		if winningHand, over := isOver(hands); over {
			return score(winningHand), nil
		}
	}
}

// returns cards in descending sorted order
func judge(round map[string]int) (winner string, cards []int) {
	var maxCard int
	for player, card := range round {
		cards = append(cards, card)
		if card > maxCard {
			maxCard = card
			winner = player
		}
	}
	// reverse sort
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	return winner, cards
}

func isOver(hands map[string][]int) (winningHand []int, over bool) {
	numWithCards := 0
	for _, cards := range hands {
		if len(cards) > 0 {
			numWithCards++
			winningHand = cards
		}
	}
	if numWithCards > 1 {
		return nil, false
	}
	return winningHand, true
}

func score(hand []int) int {
	out := 0
	for i := 1; len(hand)-i >= 0; i++ {
		out += i * hand[len(hand)-i]
	}
	return out
}

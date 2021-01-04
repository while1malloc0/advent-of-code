package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

func getScore(hand []int) int {
	var score int
	for i := range hand {
		pos := len(hand) - 1 - i
		mult := i + 1
		score += hand[pos] * mult
	}

	return score
}

func parseDeck(in string) ([]int, []int) {
	halves := strings.Split(in, "\n\n")
	decks := make([][]int, 2)
	for i := range halves {
		half := halves[i]
		r := strings.NewReader(half)
		s := bufio.NewScanner(r)
		deck := []int{}
		for s.Scan() {
			line := s.Text()
			if strings.HasPrefix(line, "Player") {
				continue
			}
			card, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			deck = append(deck, card)
		}
		decks[i] = deck
	}
	return decks[0], decks[1]
}

func playTurn(playerOne, playerTwo []int) ([]int, []int) {
	// One player already won, nothing to do
	if len(playerOne) == 0 || len(playerTwo) == 0 {
		return playerOne, playerTwo
	}

	playerOneCard := playerOne[0]
	playerTwoCard := playerTwo[0]

	playerOne = playerOne[1:]
	playerTwo = playerTwo[1:]

	if playerOneCard > playerTwoCard {
		playerOne = append(playerOne, playerOneCard)
		playerOne = append(playerOne, playerTwoCard)
	} else {
		playerTwo = append(playerTwo, playerTwoCard)
		playerTwo = append(playerTwo, playerOneCard)
	}

	return playerOne, playerTwo
}

func playGame(in string) int {
	playerOneDeck, playerTwoDeck := parseDeck(in)

	for len(playerOneDeck) > 0 && len(playerTwoDeck) > 0 {
		playerOneDeck, playerTwoDeck = playTurn(playerOneDeck, playerTwoDeck)
	}

	var winningPlayerDeck []int
	if len(playerOneDeck) == 0 {
		winningPlayerDeck = playerTwoDeck
	} else {
		winningPlayerDeck = playerOneDeck
	}

	winningScore := getScore(winningPlayerDeck)
	return winningScore
}

var playerOneMemo = make(map[string]struct{})
var playerTwoMemo = make(map[string]struct{})

func makeHash(in []int) string {
	hash := ""
	for i := range in {
		val := strconv.Itoa(in[i])
		hash += val
	}
	return hash
}

func playTurnRecursive(playerOne, playerTwo []int) ([]int, []int) {
	if len(playerOne) == 0 || len(playerTwo) == 0 {
		return playerOne, playerTwo
	}

	playerOneHash := makeHash(playerOne)
	playerTwoHash := makeHash(playerTwo)
	if _, ok := playerOneMemo[playerOneHash]; ok {
		if _, ok := playerTwoMemo[playerTwoHash]; ok {
			playerTwo = []int{}
			return playerOne, playerTwo
		}
	}

	playerOneMemo[playerOneHash] = struct{}{}
	playerTwoMemo[playerTwoHash] = struct{}{}

	playerOneCard := playerOne[0]
	playerOne = playerOne[1:]

	playerTwoCard := playerTwo[0]
	playerTwo = playerTwo[1:]

	if len(playerOne) >= playerOneCard && len(playerTwo) >= playerTwoCard {
		var playerOneCopy []int
		for i := 0; i < playerOneCard; i++ {
			playerOneCopy = append(playerOneCopy, playerOne[i])
		}

		var playerTwoCopy []int
		for i := 0; i < playerTwoCard; i++ {
			playerTwoCopy = append(playerTwoCopy, playerTwo[i])
		}

		for len(playerOneCopy) > 0 && len(playerTwoCopy) > 0 {
			// playerOneHash := makeHash(playerOneCopy)
			// playerTwoHash := makeHash(playerTwoCopy)
			// if _, ok := playerOneMemo[playerOneHash]; ok {
			// 	if _, ok := playerTwoMemo[playerTwoHash]; ok {
			// 		playerTwoCopy = []int{}
			// 		break
			// 	}
			// }
			// playerOneMemo[playerOneHash] = struct{}{}
			// playerTwoMemo[playerTwoHash] = struct{}{}
			playerOneCopy, playerTwoCopy = playTurnRecursive(playerOneCopy, playerTwoCopy)
		}

		if len(playerOneCopy) > 0 {
			playerOne = append(playerOne, playerOneCard, playerTwoCard)
		} else {
			playerTwo = append(playerTwo, playerTwoCard, playerOneCard)
		}

		return playerOne, playerTwo
	}

	if playerOneCard > playerTwoCard {
		playerOne = append(playerOne, playerOneCard, playerTwoCard)
	} else {
		playerTwo = append(playerTwo, playerTwoCard, playerOneCard)
	}

	return playerOne, playerTwo
}

func playGameRecursive(in string) int {
	playerOneDeck, playerTwoDeck := parseDeck(in)

	for len(playerOneDeck) > 0 && len(playerTwoDeck) > 0 {
		playerOneDeck, playerTwoDeck = playTurnRecursive(playerOneDeck, playerTwoDeck)
	}

	var winningPlayerDeck []int
	if len(playerOneDeck) == 0 {
		winningPlayerDeck = playerTwoDeck
	} else {
		winningPlayerDeck = playerOneDeck
	}
	winningScore := getScore(winningPlayerDeck)
	return winningScore
}

func main() {
	partOneFunc := func() error {
		input, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		winningScore := playGame(string(input))
		fmt.Println(winningScore)
		return nil
	}

	partTwoFunc := func() error {
		input, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		winningScore := playGameRecursive(string(input))
		fmt.Println(winningScore)
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}

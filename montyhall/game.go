package main

import "math/rand"

// GameState represents game state with randomly located goats(false) and car(true)
type GameState [3]bool

// OneGameStats is a stats from one game
type OneGameStats struct {
	State         GameState
	PFirstChoise  byte
	PSecondChoise byte
	GOpenedDoor   byte
	PIsWin        bool
}

// GameStats is a global stats from all games
type GameStats []OneGameStats

// gameNewState returns new game state (random state with one car and two goats)
func gameNewState() GameState {
	var gameState GameState
	randIndex := rand.Intn(3)
	gameState[randIndex] = true
	return gameState
}

// gameOpenGoatDoor returns a number of door that contains a goat, but not a user's choice
func (s *GameState) gameOpenGoatDoor(playersChoice byte) byte {
	var res byte
	if !s[playersChoice] {
		switch playersChoice {
		case 0:
			if s[1] {
				res = 2
			} else {
				res = 1
			}
		case 1:
			if s[0] {
				res = 2
			} else {
				res = 0
			}
		case 2:
			if s[1] {
				res = 0
			} else {
				res = 1
			}
		}
	} else {
		for {
			randIndex := rand.Intn(3)
			if byte(randIndex) != playersChoice {
				return byte(randIndex)
			}
		}
	}
	return res
}

// gameIsPlayerWin checks is player win
func (s *GameState) gameIsPlayerWin(choise byte) bool {
	if s[choise] {
		return true
	} else {
		return false
	}
}

// gameWinProbability calculates wins probability
func gameWinProbability(rounds, wins int) float64 {
	w := float64(wins)
	r := float64(rounds)
	return w / r
}

// gameWinsRounds returns number of rounds and wins from stats
func (stats GameStats) gameWinsRounds() (int, int) {
	var wins int = 0
	rounds := len(stats)
	for _, s := range stats {
		if s.PIsWin {
			wins += 1
		}
	}
	return rounds, wins
}

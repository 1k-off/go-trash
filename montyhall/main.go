package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var stats GameStats
	rounds := 20
	gamblingComission := true

	for i := 1; i <= rounds; i++ {
		stats = playGame(stats)
	}

	playedRounds, wins := stats.gameWinsRounds()
	winProbability := gameWinProbability(playedRounds, wins)

	fmt.Println("WIN probability: ", winProbability)

	if gamblingComission {
		statsJson, _ := json.Marshal(stats)
		fmt.Println(string(statsJson))
	}
}

// playGame simulates full gameplay for one round with all game and player choices and returns stats for all previous and current rounds
func playGame(stats GameStats) GameStats {
	var currentStats OneGameStats
	s := gameNewState()
	playersChoice := playerFirstDecision()
	goatDoor := s.gameOpenGoatDoor(playersChoice)
	playersNewChoise := playerFinalDecisionChange(playersChoice, goatDoor)

	currentStats.State = s
	currentStats.PFirstChoise = playersChoice
	currentStats.GOpenedDoor = goatDoor
	currentStats.PSecondChoise = playersNewChoise
	currentStats.PIsWin = s.gameIsPlayerWin(playersNewChoise)

	return append(stats, currentStats)
}

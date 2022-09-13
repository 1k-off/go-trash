package main

import "math/rand"

// playerFirstDecision returns random number behind 0-2 and represents random player's choice from 3 doors
func playerFirstDecision() byte {
	return byte(rand.Intn(3))
}

// playerFinalDecisionChange returns number that is not first player's choice and is not a door that game opened
func playerFinalDecisionChange(previousDecision byte, openedDoor byte) byte {
	var res byte
	switch previousDecision {
	case 0:
		if openedDoor == 1 {
			res = 2
		} else {
			res = 1
		}
	case 1:
		if openedDoor == 0 {
			res = 2
		} else {
			res = 0
		}
	case 2:
		if openedDoor == 0 {
			res = 1
		} else {
			res = 0
		}
	}
	return res
}

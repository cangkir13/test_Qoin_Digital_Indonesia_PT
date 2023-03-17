package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	name  string
	dices []int
	score int
}

func (p *Player) RollDices() {
	for i := range p.dices {
		p.dices[i] = rand.Intn(6) + 1
	}
}

func (p *Player) EvaluateDices(neighbor *Player) {
	for i := 0; i < len(p.dices); i++ {
		switch p.dices[i] {
		case 1:
			neighbor.AddDice(1)
			p.dices[i] = 0
		case 6:
			p.score++
			p.dices[i] = 0
		}
	}
	p.dices = removeZeros(p.dices)
}

func (p *Player) AddDice(dice int) {
	p.dices = append(p.dices, dice)
}

func (p *Player) HasDices() bool {
	return len(p.dices) > 0
}

func removeZeros(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v != 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numPlayers := 3
	numDices := 4

	players := make([]*Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = &Player{
			name:  fmt.Sprintf("Pemain #%d", i+1),
			dices: make([]int, numDices),
		}
	}

	round := 1
	for {
		fmt.Printf("Giliran %d lempar dadu:\n", round)
		anyoneHasDices := false
		for i := range players {
			if players[i].HasDices() {
				anyoneHasDices = true
				players[i].RollDices()
				fmt.Printf("%s (%d): %v\n", players[i].name, players[i].score, players[i].dices)
			}
		}
		if !anyoneHasDices {
			break
		}

		for i := range players {
			if !players[i].HasDices() {
				continue
			}
			neighbor := players[(i+1)%numPlayers]
			players[i].EvaluateDices(neighbor)
			fmt.Printf("%s (%d): %v\n", players[i].name, players[i].score, players[i].dices)
		}

		round++
	}

	fmt.Println("==================")
	maxScore := -1
	var winner *Player
	for _, p := range players {
		if p.score > maxScore {
			maxScore = p.score
			winner = p
		}
		fmt.Printf("%s (%d): %v\n", p.name, p.score, p.dices)
	}
	fmt.Printf("==================\nGame dimenangkan oleh %s karena memiliki poin lebih banyak dari pemain lainnya.", winner.name)
}

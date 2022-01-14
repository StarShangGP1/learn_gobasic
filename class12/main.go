package main

import (
	"fmt"
	"math/rand"
)

func main() {
	vg := &voteGame{students: []*student{
		{name: "1"},
		{name: "2"},
		{name: "3"},
		{name: "4"},
		{name: "5"},
		{name: "6"},
		{name: "7"},
		{name: "8"},
		{name: "9"},
		{name: "10"},
	}}
	leader := vg.goRun()
	fmt.Println(leader)
}

type student struct {
	name     string
	agree    int
	disagree int
}

type voteGame struct {
	students []*student
}

type Leader student

func (g *voteGame) goRun() *Leader {
	//randInt := rand.Int()
	//p := g.students[randInt%len(g.students)]
	for _, item := range g.students {
		randInt := rand.Int()
		if randInt%2 == 0 {
			item.voteA(g.students[randInt%len(g.students)])
		} else {
			item.voteD(g.students[randInt%len(g.students)])
		}
	}
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 {
		return (*Leader)(g.students[maxScoreIndex])
	}
	return nil
}

func (std *student) voteA(target *student) {
	target.agree++
}

func (std *student) voteD(target *student) {
	target.disagree++
}

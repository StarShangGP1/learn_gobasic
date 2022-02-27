package main

import (
	"fmt"
	"learn_go/homework/class18/fatRateRank"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var personFatRate = map[string]float64{}
var lock sync.Mutex

func fatRateRankLock() {
	PersonNum := 1000
	var Persons []Person
	rand.Seed(time.Now().Unix())

	for i := 0; i < PersonNum; i++ {
		Persons = append(Persons, Person{
			name:    fmt.Sprintf("Person%d", i),
			fatRate: fatRateRank.RandFR(0.1, 0.4, 1),
		})
	}

	for _, runnerItem := range Persons {
		runnerItem.register(runnerItem.name, runnerItem.fatRate)
	}

	for {
		for i := 0; i < PersonNum; i++ {
			wg := sync.WaitGroup{}
			wg.Add(1)
			Persons[i].fatRate = fatRateRank.RandFR(0.1, 0.4, 1)
			go func(name string, fatRate float64, wg *sync.WaitGroup) {
				defer wg.Done()
				update(Persons[i].name, Persons[i].fatRate)
			}(Persons[i].name, Persons[i].fatRate, &wg)
			wg.Wait()
			rank, _ := getRand(Persons[i].name)
			fmt.Printf("%s的排名是%d\n", Persons[i].name, rank)
		}
	}
}

func update(name string, rate float64) {
	lock.Lock()
	personFatRate[name] = rate
	lock.Unlock()
}

func (p Person) register(name string, fatRate float64) {
	lock.Lock()
	personFatRate[name] = fatRate
	lock.Unlock()
}

func getRand(name string) (rank int, fataRate float64) {
	fatRate2PersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
		rankArr = append(rankArr, frItem)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		_names := fatRate2PersonMap[frItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fataRate = frItem
				return
			}
		}
	}
	return 0, 0
}

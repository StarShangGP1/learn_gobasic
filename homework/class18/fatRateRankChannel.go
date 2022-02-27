package main

import (
	"fmt"
	"learn_go/homework/class18/fatRateRank"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Person struct {
	name    string
	fatRate float64
	rank    int
}

func fatRateRankChannel() {
	for {
		PersonCh := make(chan Person, 1000)
		rankCh := make(chan []Person, 1000)
		var PersonSlice []Person
		PersonCounter := 1000
		wg := sync.WaitGroup{}
		wg.Add(PersonCounter)
		rand.Seed(time.Now().Unix())

		for i := 0; i < PersonCounter; i++ {
			go func(i int, wg *sync.WaitGroup) {
				defer wg.Done()
				var Person = Person{
					name:    fmt.Sprintf("Person%d", i),
					fatRate: fatRateRank.RandFR(0.1, 0.4, 1),
				}

				PersonCh <- Person

				getRandChan(Person.name, rankCh)
			}(i, &wg)
		}

		finishedFileCount := 0
		for Person := range PersonCh {
			finishedFileCount++
			PersonSlice = append(PersonSlice, Person)
			if finishedFileCount == PersonCounter {
				close(PersonCh)
			}
		}

		sort.Slice(PersonSlice, func(i, j int) bool {
			return PersonSlice[i].fatRate < PersonSlice[j].fatRate
		})

		for i, _ := range PersonSlice {
			PersonSlice[i].rank = i + 1
		}
		for i := 0; i < PersonCounter; i++ {
			rankCh <- PersonSlice
		}

		wg.Wait()
	}
}

func getRandChan(name string, rank chan []Person) {
	PersonSlice := <-rank
	for _, Person := range PersonSlice {
		if Person.name == name {
			fmt.Printf("%s的体脂排名为%d\n", name, Person.rank)
		}
	}
}

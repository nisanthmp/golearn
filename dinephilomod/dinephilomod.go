package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var chPermit, chDone chan int

type chopS struct{ sync.Mutex }

var chopSs [5]*chopS

type philo struct {
	num         int
	philoChopSs [2]*chopS
}

var philos [5]*philo

func main() {
	chPermit = make(chan int)
	chDone = make(chan int)

	for i := 0; i < 5; i++ {
		chopSs[i] = new(chopS)
	}
	for i := 0; i < 5; i++ {
		philos[i] = newPhilo(i)
	}
	go host()

	wg.Add(15)
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			go philos[i].eat()
		}
	}
	wg.Wait()
}

func host() {
	chPermit <- 1
	for i := 0; i < 14; i++ {
		chPermit <- 1
		<-chDone
	}
	<-chDone
}

func newPhilo(i int) *philo {
	var p philo
	p.num = i + 1
	p.philoChopSs[0] = chopSs[i]
	p.philoChopSs[1] = chopSs[(i+1)%5]
	return &p
}

func (p philo) eat() {
	//for i := 0; i < 3; i++ {
	<-chPermit

	p.philoChopSs[0].Lock()
	p.philoChopSs[1].Lock()

	fmt.Println("starting to eat", p.num)
	fmt.Println("finishing eating", p.num)

	p.philoChopSs[1].Unlock()
	p.philoChopSs[0].Unlock()

	chDone <- 1
	//}
	wg.Done()
}

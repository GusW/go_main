package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var philosEatingMut sync.Mutex
var philosEating int = 0

func incrementMutTarget(){
	philosEatingMut.Lock()
	philosEating++
	philosEatingMut.Unlock()
}

func decrementMutTarget(){
	philosEatingMut.Lock()
	philosEating--
	philosEatingMut.Unlock()
}

var ph1ETMut, ph2ETMut, ph3ETMut, ph4ETMut, ph5ETMut sync.Mutex
var ph1ET, ph2ET, ph3ET, ph4ET, ph5ET int = 0,0,0,0,0

func incrementPhiloEatTimes(philoIdx int){
	switch philoIdx {
	case 0:
		ph1ETMut.Lock()
		ph1ET++
		ph1ETMut.Unlock()
	case 1:
		ph2ETMut.Lock()
		ph2ET++
		ph2ETMut.Unlock()
	case 2:
		ph3ETMut.Lock()
		ph3ET++
		ph3ETMut.Unlock()
	case 3:
		ph4ETMut.Lock()
		ph4ET++
		ph4ETMut.Unlock()
	case 4:
		ph5ETMut.Lock()
		ph5ET++
		ph5ETMut.Unlock()
	}
}

func getPhiloEatTimes(philoIdx int)int{
	switch philoIdx {
	case 0:
		return ph1ET
	case 1:
		return ph2ET
	case 2:
		return ph3ET
	case 3:
		return ph4ET
	case 4:
		return ph5ET
	default:
		return 0
	}
}

var donePhiloIdxs = []int{}
var donePhilosMut sync.Mutex

func addDonePhilo(philoIdx int){
	donePhilosMut.Lock()
	donePhiloIdxs=append(donePhiloIdxs, philoIdx)
	donePhilosMut.Unlock()
	wg.Done()
}

type boolChan chan bool

type ChopS struct{ sync.Mutex }

type Philo struct {
	number int
	leftCS, rightCS  *ChopS
}

func (p Philo) eat(canEat boolChan, philoIdx int) {
	for ce := range canEat{
		if ce{
			p.leftCS.Lock()
			p.rightCS.Lock()
			if getPhiloEatTimes(philoIdx) < 3{
				fmt.Println("start to eat", p.number)
				incrementMutTarget()
				incrementPhiloEatTimes(philoIdx)
				fmt.Println("finishing eating", p.number)
				if getPhiloEatTimes(philoIdx) == 3{
					addDonePhilo(philoIdx)
				}
				decrementMutTarget()
			}
			p.leftCS.Unlock()
			p.rightCS.Unlock()
		}
	}
}

type Host struct{}

func (h Host) handleDiningPhilos(philos []*Philo){
	wg.Add(len(philos))
	canEatChans := make([]boolChan, len(philos))

	for idx := range philos{
		canEatChans[idx] = make(boolChan)
		go philos[idx].eat(canEatChans[idx], idx)
	}

	for len(donePhiloIdxs) < len(philos){
		for idx := range philos{
			if philosEating < 2 {
				canEatChans[idx] <- true
			}
		}
	}
	wg.Done()
	wg.Wait()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < cap(CSticks); i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < cap(philos); i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	wg.Add(1)
	host := Host{}
	go host.handleDiningPhilos(philos)
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

*/

var wg sync.WaitGroup

type ChopStick struct {
	sync.Mutex
}

type Philospher struct {
	id int
	leftChopStick *ChopStick
	rightChopStick *ChopStick
	eatRequest chan int
	startFlag chan bool
	updateEatFinish chan int
	eatCount int
}

/*
	Eat method checks the count when the count reaches value of 3 it quits the goroutine and decrements waitgroup
	Sends the professor id to the eatRequest channel to request eating
	Reads the startEat channel if the value is true it proceeds for eating and also increments the eat counter which helps in quitting when it reaches 3
	Once eating is finished it updates the updateEatFinish channel
 */
func (p Philospher) eat(){
	for {
		if p.eatCount >= 3{
			p.updateEatFinish <- p.id
			fmt.Println("calling done")
			wg.Done()
			return
		}
		p.eatRequest <- p.id
		if startEat :=<- p.startFlag; startEat {
			p.leftChopStick.Lock()
			p.rightChopStick.Lock()
			fmt.Println("starting to eat :",p.id)
			fmt.Println("finishing eating :",p.id)
			p.eatCount++
			p.updateEatFinish <- p.id
			p.rightChopStick.Unlock()
			p.leftChopStick.Unlock()
		}
	}

}

/*
	Host goroutine approves the request by checking the profDineMap,
	if map length is lesser than 2 it updates the startEat channel to true else it updates to false
	Also, it listens to updateEatFinish channel when updateEat is received it deletes the value from Map
	When it receives any data in abortchannel it quit the host goroutine
 */
func host(eatRequest chan int, startEat chan bool, updateEatFinish chan int, abortChannel chan string){
	profDineMap := make(map[int]string)
	for{
		select {
			case req_id := <- eatRequest:
				fmt.Println("Eat Request from professor id :", req_id)
				if len(profDineMap) < 2 {
					profDineMap[req_id] = string(req_id)
					startEat <- true
				}else{
					startEat <- false
				}
			case del_id := <- updateEatFinish:
				fmt.Println("Update Finish request:", del_id)
				delete(profDineMap,del_id)
			case <- abortChannel:
				fmt.Println("Abort request:")
				return
		}
	}
}

func main(){
	CStick :=make([]*ChopStick,5)
	PhilSlice := make([]*Philospher, 5)
	eatRequest := make(chan int)
	startEat := make(chan bool)
	updateEatFinish := make(chan int)
	abortChannel := make(chan string)

	//Initializing the Chopstick
	for i:=0;i<5;i++{
		CStick[i] = new(ChopStick)
	}
	//Initializing the Philospher
	for i:=0;i<5;i++{
		PhilSlice[i] = &Philospher{
			i,
			CStick[i],
			CStick[(i+1)%5],
			eatRequest,
			startEat,
			updateEatFinish,
			0,
		}
	}
	wg.Add(5)
	// Host go routine with all the channels
	go host(eatRequest, startEat,updateEatFinish,abortChannel)
	for i:=0; i< 5; i++{
		go PhilSlice[i].eat()
	}
	wg.Wait()
	abortChannel <- "Abort"
}
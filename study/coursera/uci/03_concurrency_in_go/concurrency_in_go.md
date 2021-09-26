## Concurrency

- Parallel Execution => execute exactly at the same time

  - need replicated hardware
  - better throughput
  - some tasks must be performed sequentially
    - wash dish, dry dish

- Von Neumann bottleneck

  - cache access time => 1 clock cycle
  - main memory access time => ~100 clock cycles
  - increased cache capacity

- Moore's law

  - transistor density would double every two years but:
    - more transitors => more power consumption
    - more power => higher temperature
    - `power wall`
    - `P = a * CFV2`
      - a => percentage of time switching
      - C => capacitance (related to size)
      - F => clock frequency
      - V => voltage swing (from low to high)

- Dennard Scaling

  - Voltage should scale with transistor size
  - Keeps power (hence temperature) low
  - Problem => `voltage can't go too low`
    - must stay above voltage threshold
    - noise more significant in small voltage swings (ranges)
  - Leakage power
    - Transistors conductors go smaller forcing insulators to be thicker

- Multi-Core systems
  - Parallel execution needed to exploit multi-core

---

## Concurrent Execution

- Tasks are concurrent when start and end time overlap
- Can be executed in either different or same hardware
- Tasks alternate processor time
  - Only one task actually executed at time

---

## Concurrent Programming

- Programmer determines which tasks can be executed in parallel
- If they will be executed with parallelism or concurrency depends on `hardware mapping`
  - operating system
  - Go runtime schedule
  - `not in control of the programmer`
- Hardware Mapping in Go
  - Programmers not aware of underlying architecture

---

## Hiding Latency

- Concurrency improves performance even w/o parallelism
  - Tasks must periodically wait for something i.e wait for memory

---

## Processes

- Memory (virtual address space)
  - Code
  - Stack
  - Heap
  - Shared libs
- Registers
  - Program counter
  - data regs

---

## Operating System

- Executes several processes concurrently
- Process scheduling
  - process switched quickly: 20 ms
- Impression of parallelism
- Must allocate fair amount of resources to each process

---

## Context Switch

- Control flow changes from one process to the other
  - and back and forth...
- Process "context" `state` must be swapped
- Threads vs. Processes
  - Threads are leightweight and share context inside a process
  - Because of shared state, the context switch is faster than processing

---

## Gorountine

- Threads in Go
  - The OS schedules a main thread
  - Inside of that main thread Go creates several Goroutines
- Go Runtime Scheduler
  - Schedule Goroutines inside an OS Thread

---

## Interleavings

- Order os execution of tasks in different concurrent processes is unknown
  - Non deterministic, can change at every run

---

## Race Conditions

- Outcome that depends on the non-deterministic interleaving
- Goroutines exchange state via communication

---

## Create Goroutine

- One Goroutine is created always to `main` func
-     a := 1
      go foo()
      b := 2
- When the main Goroutine is complete all other Goroutines exit

---

## Synchronization

- Use of global event which will be shared between different Goroutines
- Necessary 'evil' => breaks the idea of concurrency and parallelism by
  - ordering
  - thus, sequential

---

## Wait Groups

- Type of sync
  - `sync.WaitGroup`
  - Incrementes an internal counter for every Goroutine waited
  - Decrements the internal counter for every Goroutine completed
-     func foo(wg *sync.WaitGroup){
        fmt.Println("Foo Goroutine")
        wg.Done() // decrements the Counter
      }

      func main(){
        var wg sync.WaitGroup
        wg.Add(1) // increments the Counter
        go foo(&wg)
        wg.Wait() // wait to counter == 0
        fmt.Println("Main Goroutine")
      }

---

## Communications

- Needs to send data from main Goroutine to sub-Goroutines
- Needs to send results from sub-Goroutines to main Goroutine
- Channels => used to communication between Goroutines

  - typed
  -     myChannel := make(chan int)
        myChannel <- 3 // Send data on the channel
        myInt := <- myChannel // receive data from a channel
  -     func prod(x int, y int, c chan int) {
          c <- x * y
        }

        func main(){
          c := make(chan int)
          go prod(1, 2, c)
          go prod(3, 4, c)

          firstProd := <- c // not necessarialy from 1 and 2
          secondProd := <- c // not necessarialy from 3 and 4

          fmt.Println(firstProd * secondProd)
        }

  - Default channels are `unbuffered` => cannot hold data in transit
    - Sending blocks until data is received
    - Receiving blocks until data is sent
  - Channel communication is synchronous
    - Blocking is the same as waiting for communication
    -     c <- 3
          ...
          <- c // discarded data, used only for synchorization and not for data exchange (same as sync.WaitGroup)
  - `Buffered` channels

    - `Capacity` => the number of objects it can hold in transit
    -     c := make(chan, int, 3)
    - Sending only blocks if the `buffer is full`
    - Recieving only blocks if the `buffer is empty`

  - Iterating through a channel => continuely reading from channel

    -     for i := range myChannel {
            fmt.Println(i)
          }

          close(myChannel) // close the channel if iterating on range

  - Receiving from multiple Goroutines => `sequentially`
    -     a := <- c1
          b := <- c2
          fmt.Println(a*b)
  - Select statement => fist come, first served
    - Need data from first finished channel
    -     select {
            case a = <- c1:
              fmt.Prinln(a)
            case b = <- c2:
              fmt.Prinln(b)
          }
  - Select `Send` or `Receive` => whichever executes first
    -     select {
            case a = <- inchan:
              fmt.Prinln("Received from inchan")
            case outchan <- b:
              fmt.Prinln("Sent b to outchan")
          }
  - Select with `Abort Channel`
    -     for {
            select {
            case a = <- inchan:
              fmt.Prinln("Received from inchan")
            case <- abort:
              return // exits infinite loop
          }
  - Select with non blockers `default`
    -     select {
            case a = <- c1:
              fmt.Prinln(a)
            case b = <- c2:
              fmt.Prinln(b)
            default:
              fmt.Println("not blocked")
          }

---

## Goroutines sharing variables

- Concurrency `safe` => function `does not interfere` with other Goroutines
- `Bad` variable sharing

  -     i := 0
        var wg sync.WaitGroup

        func inc() {
          i += 1
          wg.Done()
        }

        func main() {
          wg.Add(2)
          go inc()
          go inc()
          wg.Wait()
          fmt.Print(i)
        }

- Concurrency is at the `machine code level`, not source code level

  -     read i
        increment
        write i

- Correct Sharing => `Mutex` Mutual Exclusion

  - Do not let 2 Goroutines write to a shared variable at the same time
  - Restrict possible interleavings
  - Writting to shared varibles should be `mutually exclusive`

- `sync.Mutex`

  - Binary semaphore
    - `Flag up` => shared variable in use => `Lock()`
    - `Flag down` => shared variable available => `Unlock()`
  - `Lock()` blocks subsequential Goroutines until `Unlock()`
  -     i := 0
        var mut sync.Mutex

        func inc(){
          mut.Lock()
          i += 1
          mut.Unlock()
        }

---

## Sync Package

- `sync.Once`

  - Function is executed only once does not matter how many Goroutines call it
  - All calls to once.Do() blocks until the 1st returns
  -     var wg sync.WaitGroup
        var on sync.Once

        func setUp(){
          fmt.Println('init')
        }

        func doStuff(){
          on.Do(setUp)
          fmt.Println('hello')
          wg.Done()
        }

        func main(){
          wg.Add(2)
          go doStuff()
          go doStuff()
          wg.Wait()
        }

- `Deadlock` => Circular sync dependencies

  -     func whatever(c1, c2 chan int) {
          <- c1    // read from 1st channel
          c2 <- 1  // write to 2nd channel
          wg.Done()
        }

        func main(){
          c1 := make(chan, int)
          c2 := make(chan, int)
          wg.Add(2)
          go whatever(c1, c2)
          go whatever(c2, c1)
          wg.Wait()
        }

---

## Dining Philosophers Problem

- 5 philosophers on a table
- 1 chopstick is placed between each adjacent pair
- To eat rice you need 2 chopsticks, the left and right one
- When 1 philosopher eats, its neighbors cannot

  - Each chopstick is a `mutex`
  - Each philosopher is associated with a Goroutine and 2 chopsticks
  -      type ChopS struct{ sync.Mutex }

         type Philo struct { leftCS, rightCS *ChopS }

         func (p Philo) eat(){
           for {
             p.LeftCS.Lock()
             p.RightCS.Lock()

             fmt.Println('eating')

             p.LeftCS.Unlock()
             p.RightCS.Unlock()
           }
         }

         func main(){
           CSticks := make([]*ChopS, 5)
           for i := 0; i < cap(CSticks); i++{
             CSticks[i] = new(ChopS)
           }

           philos := make([]*Philo, 5)
           for i := 0; i < cap(philos); i++{
             philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
           }

           for i := 0; i < len(philos); i++{
             go philos[i].eat()
           }
         }

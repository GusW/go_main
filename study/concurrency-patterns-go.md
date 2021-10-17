## Concurrency Patterns in Go

---

https://www.youtube.com/watch?v=YEKjSzIwAdA

---

### Communicating Sequential Processes (CSP) - 1978

- Each process is built for sequential execution
- Data is communicated between processes via channels
- No shared state
- Scale by adding more of the same

---

### Go's concurrency toolset

- goroutines
- channels
- select
- sync package

---

### Channels

- "bucket chain"
- Sender, buffer, Receiver
- buffer optional
- Blocking => Deadlocks, impossibility to scale
  - Unbuffered channel - No sender - No receiver
  - Buffered channel - No sender - Sending > than buffer size
- Closing channels

  - Sends a special "closed" message
  - Receiver gets message and stop listening
  - If Sender tries to send more: `panic`
  - Closing twice => sends a new message on a closed channel: `panic`
  -     c := make(chan int)

        close(c)

        fmt.Println(<-c) // 0, false

    - 0 => zero value of int
    - false => "no more data"

---

### Select

- Switch-alike statement on channel operations
- Order of cases does not matter
- Default case
- First non-blocking case is chosen (send and/or receive)
- Making channels non-blocking:
-     func TryRecieve(c <- chan int, duration time.Duration) (data int, more, ok bool) {
          select {
          case data, more = <- c:
              return data, more, true
          case <- time.After(duration): return 0, true, false
          }
      }
- Channels are streams of data

  - Fan-out: 1 -> select -> N
  - Funnel: N -> select -> 1
  - Turnout: N -> select -> N
  -     func Fanout(In <-chan int, OutA, OutB chan int) {
            for data := range In{ // Receive until closed
                select { // Send to 1st non-blocking channel
                case OutA <- data: ...
                case OutB <- data: ...
                }
            }
        }

        func Turnout(InA, InB <- chan int, OutA, OutB chan int) {
        ...
        for {
            select {    // Receive to 1st non-blocking channel
            case data, more = <- InA: ...
            case data, more = <- InB: ...
            case <- Quit:
                close(InA)  // anti-pattern: receiver is closing sender?
                close(InB)

                Fanout(InA, OutA, OutB) // flush the remaining data
                Fanout(InB, OutA, OutB)
                return
            }

            if !more {
                ...
                return
            }

            select {    // Send to 1st non-blocking channel
            case OutA <- data: ...
            case OutB <- data: ...
            }
        }

- Channel fails
  - Can create deadlocks
  - Pass around copies, impacting performance
  - Passing pointers => likely to create race conditions
  - Shared caches and registries => 'sharing' state

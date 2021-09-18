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

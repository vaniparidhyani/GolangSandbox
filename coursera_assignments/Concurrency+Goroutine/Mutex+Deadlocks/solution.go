package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number int
	leftChopstick, rightChopstick *Chopstick
}

type Host struct {
        // count the no. of philo eating concurrently
	count int
}

func (host Host) askHost() bool {
	// host allows no more than 2 philosophers to eat concurrently
	if host.count < 2 {
		host.count++
		return true
	}
	return false
}

func (host Host) leaveHost() {
	// decrement count when done eating
	host.count--
}

func main() {

	// initialized 5 chopsticks
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// initialized 5 philos
	philosophers := make([]Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
	}

	host := Host{0}
	wg := sync.WaitGroup{}

	// since each philo eats 3 times hence calling 5*3 eat goroutine
	wg.Add(15)
	for i := 0; i < 5; i++ {
		// nested loop for each philo
		for j := 0; j < 3; j++ {
			// call goroutine
			go philosophers[i].eat(&host, &wg)
		}
	}
	wg.Wait()
}

func (philosopher Philosopher) eat(host *Host, wg *sync.WaitGroup) {
	defer wg.Done()
	// have each eat() take a different amount of time to demonstrate interleaving of output. Add randomness
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	host_free := false
	// wait for host to free up
	for !host_free {
		host_free = host.askHost()
	}
	
	
	philosopher.rightChopstick.Lock()
	philosopher.leftChopstick.Lock()
	fmt.Printf("starting to eat \t%d\n", philosopher.number)
	philosopher.leftChopstick.Unlock()
	philosopher.rightChopstick.Unlock()
	host.leaveHost()
	fmt.Printf("finishing eating \t%d\n", philosopher.number)
}

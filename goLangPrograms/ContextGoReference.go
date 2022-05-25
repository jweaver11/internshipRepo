// A context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
// Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context

package main

import (
	"context"
	"fmt"
	"time"
)

type Context interface {
	Deadline(deadline time.Time, ok bool)
	Err() error
	Value(key any) any
}

const shortDuration = 1 * time.Millisecond //Sets 'shortDuration' as a const variable of 1 x Millisecond

func main() {
	d := time.Now().Add(shortDuration) //Sets 'd' as a variable that adds the currently measured time plus the milliseconds

	ctx1, cancel1 := context.WithCancel(context.Background())
	//WithCancel returns a copy of parent with a new Done channel. Closes the context's Done channel when

	ctx2, cancel2 := context.WithDeadline(context.Background(), d)
	//WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than 'd'.

	ctx3, cancel3 := context.WithTimeout(context.Background(), shortDuration)
	//WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).

	//Below function used for 'WithCancel' function
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int) //declares 'dst' as variable chan int
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	defer cancel1() //Delay execution of function until above function returns

	//Below function used for 'WithCancel' function
	for n := range gen(ctx1) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	defer cancel2() //Delay execution of function until above function returns

	//Below function used for 'WithDeadline' function
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx2.Done():
		fmt.Println(ctx2.Err())
	}

	defer cancel3() //Delay execution of function until above function returns

	//Below function used for 'WithTimeout' function
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx3.Done():
		fmt.Println(ctx3.Err()) // prints "context deadline exceeded"
	}
}

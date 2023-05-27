package main

import (
	"fmt"
	"sync"
)

func hello(str string, wg *sync.WaitGroup) {
	fmt.Println(str, " 😀😀")
	wg.Done() //decreases the wait by 1
}
func main() {
	wgp := new(sync.WaitGroup)
	go hello("roshin", wgp)
	go hello("james", wgp)
	wgp.Add(2)       //2 wait is added as 2 functions to get executed
	defer wgp.Wait() //waits until the waiting counter becomes 0...which is updated by done() on hello() execution completion
	fmt.Println("\nExecuted functions completely")
}

/*
Alternative to Sleep()------
Waitgroup is a feature of sync package which allows to wait before termination of the program
until the internal counter of waitgroup becomes 0.

SYNTAX----variable_name:=new(sync.WaitGroup)

Internal counter value can be manipulated using the following functions----->>>
Add(int)	---- It increases WaitGroup counter by given integer value.
Done()  ---  It decreases WaitGroup counter by 1, we will use it to indicate termination of a goroutine.
Wait() --- It Blocks the execution until it’s internal counter becomes 0.
*/

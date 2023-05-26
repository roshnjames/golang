package main

import (
	"fmt"
	"time"
)

func sent(c chan<- string, str string) {
	c <- str
}

func recieve(c <-chan string) {
	ans := <-c
	fmt.Println("On Reciever side-----", ans)

}
func main() {
	ch := make(chan string)
	go sent(ch, "roshin james")
	go recieve(ch)
	time.Sleep((time.Second))
}

//chan<-    indicates that value is put into the channel
//<-chan    indicates that value is only read from the channel specified

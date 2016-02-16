//Example from Safari Intro to Go Programming videos by John Graham-Cumming
//https://www.safaribooksonline.com/library/view/introduction-to-go/9781491913871/
package main

import (
	"fmt"
)

func emit(strChan chan string, doneChan chan bool) {
	words := []string{"this", "that", "the other"}

	defer close(strChan)
	defer close(doneChan)
	i := 0

	//Keep putting this, that, the other on the words channel
	//until we get somethign on the quit channel.
	//Close both channels and return
	for {
		select {
		case strChan <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-doneChan:
			return
		}
	}

}

func main() {
	strChan := make(chan string)
	doneCh := make(chan bool)

	go emit(strChan, doneCh)

	//read from the strChan 100 times
	for i := 0; i < 100; i++ {
		fmt.Printf("%s \n", <-strChan)
	}

	//Send true on the doneCh channel
	doneCh <- true
}

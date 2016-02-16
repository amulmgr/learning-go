//Example from Safari Intro to Go Programming videos by John Graham-Cumming
//https://www.safaribooksonline.com/library/view/introduction-to-go/9781491913871/
package main

import (
	"fmt"
	"time"
)

func emit(strChan chan string) {
	words := []string{"this", "that", "the other"}
	defer close(strChan)

	i := 0

	t := time.NewTimer(3 * time.Second)

	//Keep putting words in the strChan channel.
	//In 3 sedonds, the timer's channel will be notified
	for {
		select {
		case strChan <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-t.C:
			return
		}
	}

}

func main() {
	strChan := make(chan string)

	//Puts some words in the strChan channel
	go emit(strChan)

	//Range over strChan channel. As words are received from the strChan channel,
	//print them
	for w := range strChan {
		fmt.Printf("%s \n", w)
	}
	//Should finish printing words in 3 seconds because the timer channel in the
	//emit goroutine will be notified.
}

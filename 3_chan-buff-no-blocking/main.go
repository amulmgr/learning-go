package main

import (
	"fmt"
)

func main(){

	// Buffered Channel Interactions : 
	// ---------------------------------
	// - Sending to a buffered channel blocks when the buffer is full. 
	// - Receives block when the buffer is empty.

	//Create a buffered channel of capacity is 2. 
	bufferedChan := make(chan int, 2)

	//Send two elements on the channel... 
	bufferedChan <- 1;
	bufferedChan <- 2;

	//Nothing is processed yet. This counter will be incremented every time we
	//receive and process something from our channel
	count := 1

	//We range over the channel, Whenever an item becomes available
	//on the chanel, print it.
	for channel := range bufferedChan{		
		fmt.Println(channel)
		//When counter reaches 2, close our channel, This will 
		//enable us to exit this for loop and avoid a deadlock
		if count == 2 {			
			close(bufferedChan) //Close the channel when it reaches capacity
		}
		count ++
	}
	
	//We should reach the end of the program without a deadlock
}
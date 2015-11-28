package main

import(
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

	//We range over the channel, Whenever an item becomes available
	//on the chanel, print it.

	//We didn't close the channel before ranging over it. 
	//We also don't conditionally close the channel once buffer
	//is empty inside the loop
	for channel := range bufferedChan{	
		//We're processing  two elements we recently put on the channel.
		//Receives block when the buffer is empty.
		//The channel is not closed. All threads are asleep
		//We're in a deadlock		
		fmt.Println(channel)
	}

	close(bufferedChan)	
}
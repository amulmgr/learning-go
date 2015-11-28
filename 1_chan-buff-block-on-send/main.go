package main

func main(){
	// Buffered Channel Interactions : 
	// ---------------------------------
	// - Sending to a buffered channel blocks when the buffer is full. 
	// - Receives block when the buffer is empty.

	
	//Create a buffered channel of capacity is 2. 
	var bufferedChan chan int
	bufferedChan = make(chan int, 2)

	//Send two elements on the channel... 
	bufferedChan <- 1;
	bufferedChan <- 2;

	//The channel is full, it started blocking when it reached full capacity
	//This should create a deadlock
	bufferedChan <- 3;

}
// Select is used to block thread until it receives value from one of the channel.

package main

import "fmt"

func main() {
	myChannel := make(chan string)
	myAnotherChannel := make(chan string)

	go func() { // Thread1 Running concurrently
		myChannel <- "Ganesh"
	}()

	go func() { // Thread2 Running concurrently.
		myAnotherChannel <- "Kachare"
	}()
	// select statement block main thread until it receives value either from mychannel or myanotherchannel.
	// What when both values available at same time. -> Random will be selected.
	select {
	case myChannelValue := <-myChannel:
		fmt.Println(myChannelValue)
	case myAnotherChannelValue := <-myAnotherChannel:
		fmt.Println(myAnotherChannelValue)
	}

}

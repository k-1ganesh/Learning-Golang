// Channels are used to send values from one thread to another.
// Go routines communicates through channels
// make(chan data_type) 

package main  
import "fmt" 

func main() {
	myChannel := make(chan string) 

	go func(){
		myChannel <- "Ganesh" // "Ganesh" is send to mychannel
	}() 
	data := <- myChannel // "Ganesh" is read from mychannel
    
	// Here main thread waits till it recives data from the mychannel.
	fmt.Println(data)   

}
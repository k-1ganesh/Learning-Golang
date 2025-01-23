package main 
import "fmt" 

// Its a kind of dictonary in python.

func main() {
   
	// Ways to create map. 

	mp1 := make(map[string]int) // map with key as string and value as int 
    mp2 := map[string]int{ // Initializing the map. 
		"age":20,
		"id":1,
	}
	mp3 := map[int]int{} // Creating empty map. 
    

	fmt.Println(mp1 , mp3) 
	// Accessing the elements of map. 
	fmt.Println(mp2["age"]) 


	// How to check key exist in map 
	// Can be done using ,ok 

	_, ok := mp2["name"]  // _ used just to ignore value.
	if ok {
		fmt.Println( "Value found")
	} else {
		fmt.Println("Value not found")
	}

	// Deleting the element from map. 

	delete(mp2,"age") 

	// Iterating over the map 
	for key,value:= range mp2 {
		fmt.Println(key , value)
	}

}
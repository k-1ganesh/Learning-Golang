// Modules are the dependency management tool of golang. 
// Modules let you create go project anywhere and manages the dependency.
// mod.go -> This file has the information of what dependencies project require.
// sum.go -> This file has the hash code of information of dependencies. (imported)
// Dependencies are the external packages required for the go code to run . 
// go mod init Module_Name :- This is used to create module. Note (Give global path if you want others to use your module by importing.)
// go mod tidy :- to remove unused dependencies and to add required dependencies.
// go get :- To download dependencies (go mod tidy can also do this) 


package main 
import (
	"fmt"
	"github.com/fatih/color" // This is dependency.
)

func main() {
	fmt.Println("Ha ha")
	color.Cyan("This is cyan color")
}
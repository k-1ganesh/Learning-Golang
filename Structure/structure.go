// Structures are used to create the custom datatype which can store the multiple values.

package main 
import "fmt"

type student struct { // This is a syntax to create struct.
	age int 
	name string
}

type college struct { // Nested structure.
	teacherAge int 
	teacherName string 
    assignedStudent student
}

type university struct { // Embedded structure.
    prn int 
	college   // all the fields of college are in university now.
}

func main() {
     stud1  := student{} 
	 stud1.age = 22 
	 stud1.name = "Ganesh" 
	 
	 fmt.Println(stud1) 

	 stud2 := student{22,"Sanjeevani"} // In this way structure can be initialised.
	 fmt.Println(stud2)   

	 col1 := college{} 
	 col1.teacherAge = 30 
	 col1.teacherName = "Gaurav"
	 col1.assignedStudent.age = 22
	 col1.assignedStudent.name = "Ganesh"
	 fmt.Println(col1) 

	 un1 := university{} 
	 un1.prn = 1234
	 un1.teacherAge = 15
	 fmt.Println(un1) 

}
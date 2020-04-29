package main

// Student holds information about student
type Student struct {
	Name          string
	StudentNumber int
}

//main will print out Student property values
func main() {
	var student = Student{"Steve Savage", 3454556}

	println("Student: ", student.Name, "| Number: ", student.StudentNumber)
}

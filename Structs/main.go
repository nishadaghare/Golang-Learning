package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"` //struct tags
}

type Address struct {
	City  string
	State string
}

type Employee struct {
	Name    string
	Age     int
	Address Address //Nested Struct
}

func updateAge(s Student) {
	s.Age = 30
}

func updateAgePtr(s *Student) {
	s.Age = 30
}

func main() {
	s1 := Student{
		Name:  "Nishada",
		Age:   20,
		Grade: "A",
	}
	fmt.Println(s1)
	s2 := Student{"Vinay", 22, "B"} //Without field names (order matters)
	fmt.Println(s2)

	var s3 Student
	fmt.Println(s3) // { 0 } → all fields take zero values

	s4 := new(Student) // returns *Student (pointer)
	s4.Name = "Eve"
	fmt.Println(s4.Name)

	s := Student{Name: "Tom", Age: 20}
	s.Age = 21 //modified
	fmt.Println(s)

	students := []Student{ //Slice of Structs
		{Name: "Rutuja", Age: 20, Grade: "A"},
		{Name: "Revati", Age: 21, Grade: "B"},
	}

	for _, s := range students {
		fmt.Println(s.Name, s.Grade)
	}

	//Anonymous Struct
	person := struct {
		Name string
		City string
	}{
		Name: "Varsha",
		City: "Pune",
	}

	fmt.Println(person.Name, person.City)

	e1 := Employee{
		Name: "Radha",
		Age:  20,
		Address: Address{ //Nested Struct
			City:  "Pune",
			State: "Maharashtra",
		},
	}

	fmt.Println(e1)

	st := Student{Name: "Maya", Age: 20}
	updateAge(st)
	fmt.Println(st.Age) // 20 (not changed)

	updateAgePtr(&st)
	fmt.Println(st.Age)

	u := Student{"Grishma", 25, "A"}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))

}

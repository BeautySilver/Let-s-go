package main

import "fmt"

type Worker struct {
	company string
	salary int
	Person struct{
     	name string
     	age int
	 }
}

type Developer struct {
	SpecialField string `example:"test"`
	level string
	project string
	worker Worker
}

type QA struct {
	SpecialField string `example:"test"`
	worker Worker
}

type ProjectManager struct {
	SpecialField string `example:"test"`
	worker Worker
}

type Analyst struct {
	SpecialField string `example:"test"`
	worker Worker
}

type Designer struct {
	SpecialField string `example:"test"`
	worker Worker
}

func main(){
	 mike:= Developer{
		 SpecialField: "test",
		 level:        "middle",
		 project:      "Abstraction",
		 worker:       Worker{company:"Guru", salary:10000, Person: struct {
			 name string
			 age  int
		 }{name:"Mike" , age:23 },
		 },
	 }

	 fmt.Println(mike)
}
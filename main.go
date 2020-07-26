package main

import (
	"fmt"
	"time"
)

type Person struct{
	name string
	age int
	birthday time.Time
}


type Worker struct {
	company string
	salary int
	person *Person

}

type Developer struct {
	SpecialField string `example:"test"`
	level string
	project string
	worker *Worker
}

type QA struct {
	SpecialField string `example:"test"`
	worker *Worker
}

type ProjectManager struct {
	SpecialField string `example:"test"`
	worker *Worker
}

type Analyst struct {
	SpecialField string `example:"test"`
	worker *Worker
}

type Designer struct {
	SpecialField string `example:"test"`
	worker *Worker
}

func (r *Person) IsAdult() string{
	if (r.age>=18){
		return "yes"
	}else{
		return "no"
	}
}

func (r *Person) HappyBirthday() string {
	if (time.Now().Day()==r.birthday.Day()  && time.Now().Month()==r.birthday.Month()) {
		return "Happy Birthday"
	}else{
		return "Ooops"
	}
}

func main(){
	mike:=new(Developer)
	mike.SpecialField= "test"

	pts := &mike

     var alex = QA{"test",
		  &Worker{company:"Wix"},
	 }

    pts2 := &alex

    artem := new(Worker)
    artem.person= &Person{birthday:time.Date(2001, 07, 26,0,0,0,0,time.UTC)}
    fmt.Println(artem.person.HappyBirthday())

    fmt.Println(pts2)
    fmt.Println(mike)
	fmt.Println(*pts)
}
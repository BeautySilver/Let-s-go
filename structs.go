
package main

import (
	"fmt"
	"time"
)
type Employee interface{
	GetInfo()
}
type Person struct{
	name string
	age int
	birthday time.Time
}


type Worker struct {
	company string
	rank string
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

func (d Developer) GetInfo(){
	fmt.Printf(" Name: %v \n Rank: %v \n Company: %v \n", d.worker.person.name, d.worker.rank, d.worker.company)
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

func describe (e Employee){
	fmt.Printf("%T \n", e)
}
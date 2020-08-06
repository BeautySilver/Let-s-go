package main

func main(){

	var mike Employee = Developer{
		worker: &Worker{
			company:"Test",
			rank:"Test",
			person: &Person{
				name:"Test"},
		},
	}

	workers := make(map[int]Employee)
	workers[1]=mike
	for _, worker := range workers{
		worker.GetInfo()
		describe(worker)
	}

}
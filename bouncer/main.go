package main

import "fmt"

// Person is person
type Person struct {
	Name string
	Age  int
}

func main() {
	peopleList := createPeople()
	for _, person := range peopleList {
		bounce := checkAge(person)
		fmt.Println(person.Name, ":", bounce)
	}
}

func createPeople() []Person {
	people := []Person{}
	names := []string{"Matt", "Dane", "Justin"}
	ages := []int{29, 55, 12}
	for i := range names {
		person := Person{names[i], ages[i]}
		people = append(people, person)
	}
	return people
}

func checkAge(person Person) string {
	if person.Age < 21 {
		return "You can't drink yet"
	}
	return "Cheers!"
}

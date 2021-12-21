package main

import "fmt"

func main() {

	personIds := []PersonId{{
		IdType:   111,
		IdNumber: "111111111111111111",
	}, {
		IdType:   222,
		IdNumber: "222222222222222222",
	}}

	persons := make([]Person, 0)

	for _, pi := range personIds {
		person := convertPersonIdToPerson(pi)
		if person != nil {
			persons = append(persons, *person)
		}
	}

	fmt.Println(persons)
}

func convertPersonIdToPerson(personId PersonId) *Person {
	person := &Person{
		id: personId.IdType,
		no: personId.IdNumber,
	}

	return person
}

type PersonId struct {
	IdType   int
	IdNumber string
}

type Person struct {
	id int
	no string
}

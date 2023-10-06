package main

import (
	"flag"
	"fmt"
)

// Person represents an individual in the family tree.
type Person struct {
	Name         string
	Gender       string
	Relationship string
	Children     []*Person
}

// FamilyTree represents the entire family tree.
type FamilyTree struct {
	People map[string]*Person
}

// NewFamilyTree initializes a new family tree.
func NewFamilyTree() *FamilyTree {
	return &FamilyTree{
		People: make(map[string]*Person),
	}
}

// AddPerson adds a new person to the family tree.
func (ft *FamilyTree) AddPerson(name, gender string) {
	person := &Person{
		Name:   name,
		Gender: gender,
	}
	ft.People[name] = person
}

// AddRelationship adds a relationship to a person in the family tree.
func (ft *FamilyTree) AddRelationship(name, relationship string) {
	person, ok := ft.People[name]
	if !ok {
		fmt.Println("Error: Person not found")
		return
	}
	person.Relationship = relationship
}

// ConnectPeople connects two people in the family tree with a specified relationship.
func (ft *FamilyTree) ConnectPeople(name1, relationship, name2 string) {
	person1, ok1 := ft.People[name1]
	person2, ok2 := ft.People[name2]

	if !ok1 || !ok2 {
		fmt.Println("Error: One or both persons not found")
		return
	}

	switch relationship {
	case "father":
		person1.Children = append(person1.Children, person2)
	case "son", "daughter":
		person2.Children = append(person2.Children, person1)
	default:
		fmt.Println("Error: Invalid relationship")
	}
}

// CountPeople counts the number of people with a specific relationship to a given person.
func (ft *FamilyTree) CountPeople(name, relationship string) int {
	person, ok := ft.People[name]
	if !ok {
		return 0
	}

	count := 0
	for _, child := range person.Children {
		if child.Gender == relationship {
			count++
		}
	}

	return count
}

// GetFather returns the name of the father of a given person.
func (ft *FamilyTree) GetFather(name string) string {
	person, ok := ft.People[name]
	if !ok {
		return "Not found"
	}

	for _, child := range ft.People {
		for _, c := range child.Children {
			if c == person && child.Relationship == "father" {
				return child.Name
			}
		}
	}

	return "Not found"
}

func main() {
	familyTree := NewFamilyTree()

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addPerson := addCommand.String("person", "", "Name of the person to add")
	addRelationship := addCommand.String("relationship", "", "Relationship to add")

	connectCommand := flag.NewFlagSet("connect", flag.ExitOnError)
	connectName1 := connectCommand.String("name1", "", "Name of the first person to connect")
	connectRelationship := connectCommand.String("relationship", "", "Relationship to add")
	connectName2 := connectCommand.String("name2", "", "Name of the second person to connect")

	countCommand := flag.NewFlagSet("count", flag.ExitOnError)
	countType := countCommand.String("type", "", "Type of relationship to count")
	countName := countCommand.String("name", "", "Name of the person")

	fatherCommand := flag.NewFlagSet("father", flag.ExitOnError)
	fatherName := fatherCommand.String("name", "", "Name of the person")

	if len(flag.Args()) < 1 {
		fmt.Println("Error: Missing command")
		return
	}

	command := flag.Arg(0)
	switch command {
	case "add":
		addCommand.Parse(flag.Args()[1:])
		if *addPerson != "" {
			familyTree.AddPerson(*addPerson, "unknown")
		}
		if *addRelationship != "" {
			familyTree.AddRelationship(*addPerson, *addRelationship)
		}
	case "connect":
		connectCommand.Parse(flag.Args()[1:])
		familyTree.ConnectPeople(*connectName1, *connectRelationship, *connectName2)
	case "count":
		countCommand.Parse(flag.Args()[1:])
		fmt.Printf("Count: %d\n", familyTree.CountPeople(*countName, *countType))
	case "father":
		fatherCommand.Parse(flag.Args()[1:])
		fmt.Printf("Father: %s\n", familyTree.GetFather(*fatherName))
	default:
		fmt.Println("Error: Unknown command")
		return
	}
}

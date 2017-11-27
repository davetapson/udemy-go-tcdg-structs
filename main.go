package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}
type address struct
{
	line1 string
	line2 string
}
type person struct {
	firstName string
	lastName string
	contact contactInfo
	address
}


func main(){

	// instantiate 1
	alex := person{"Alex", "Anderson", contactInfo{"a@b.com", 12345}, address{"1", "2"}}
	fmt.Println(alex)

	alex.firstName = "Bob"
	fmt.Println(alex)

	// instantiate 2
	var fred person
	fmt.Println(fred)
	fmt.Printf("%+v\n",fred)

	fmt.Printf("%v:%v",fred.firstName, fred.lastName)
	fmt.Println()
	fred.firstName = "Fred"
	fred.lastName = "Frondes"
	fmt.Println(fred)
	fmt.Println()
	fmt.Printf("%+v\n",fred)

	fred.contact.email = "b@e.com"
	fred.contact.zipCode = 11111
	fmt.Printf("%+v\n", fred)

	// instantiate 3
	joe := person{"Joe",
		"Jones",
		contactInfo{"z@z.com",22222,},
							address{"line 1","line 2",},
							}

	fmt.Printf("%+v\n", joe)

	// or
	joe.print()


	joePointer := &joe
	joePointer.updateName("Joejoe")
	joe.print()

	//or
	joe.updateName("Joey")
	joe.print()
}

func (pointerToPerson *person) updateName(firstName string){
	(*pointerToPerson).firstName = firstName
}

func (p person) print(){
	fmt.Printf("%+v\n", p)
}
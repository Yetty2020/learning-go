package main

import "fmt"


type speaker interface{
	speak() string
}

type dog struct{
	name string

}
type cat struct{
	name string

}

func (d dog) speak() string{
	return fmt.Sprintf("%s says woof", d.name)
}


func (c cat) speak() string{
	return fmt.Sprintf("%s says woof", c.name)
}

func introduce(s speaker){
	fmt.Println(s.speak())
}

func main() {
	introduce(dog{name: "Rex"})
	// now the cat
}
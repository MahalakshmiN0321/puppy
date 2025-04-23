package puppy

import (
	"fmt"

	"github.com/MahalakshmiN0321/dog"
)
func Bark() string{
	return "woof"
}

func Barks() string{
	return "woof woof woof!"
}

func BigBark() string{
	return dog.WhenGrowUp(Bark())

}

func BigBarks() string{
	return dog.WhenGrowUp(Barks())
	
}

func from10(){
	fmt.Println("I am from version v1.0.0")
}
func from11(){
	fmt.Println("I am from version v1.2.0")
}
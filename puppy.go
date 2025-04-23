package puppy
import (
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
	return dog.WhenGrowUp(Bark())
	
}
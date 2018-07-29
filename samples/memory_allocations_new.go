package samples

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func Run4() {
	pointer := NewZeroValuesSimpleStructure()
	// passing a reference explicitly
	fmt.Println(&pointer)

	// passing a value explicitly
	fmt.Println(*pointer)

	// this is the same as using new keyword
	var declaredStructure SimpleStructure
	spew.Dump(declaredStructure)
	//(samples.SimpleStructure) {
	//Name: (string) "",
	//		Surname: (string) ""
	//}

}

type SimpleStructure struct {
	Name    string
	Surname string
}

func NewZeroValuesSimpleStructure() *SimpleStructure {
	return new(SimpleStructure)
}

// it will work like NewZeroValuesSimpleStructure
func NewEmptySimpleStructure() *SimpleStructure {
	return &SimpleStructure{}
}

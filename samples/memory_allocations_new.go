package samples

import "fmt"

func Run4() {
	pointer := NewZeroValuesSimpleStructure()
	// passing a reference explicitly
	fmt.Println(&pointer)

	// passing a value explicitly
	fmt.Println(*pointer)
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

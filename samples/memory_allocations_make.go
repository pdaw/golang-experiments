package samples

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"reflect"
)

func Run5() {
	// final effect will be the same
	newSliceWithMake := make([]int, 0)
	newSliceWithMake = append(newSliceWithMake, 10)
	fmt.Println(newSliceWithMake)

	// final effect will be the same
	newSliceWithNew := *new([]int)
	newSliceWithNew = append(newSliceWithNew, 10)
	fmt.Println(newSliceWithNew)

	var declaredSlice []int
	// ([]int) <nil>
	spew.Dump(declaredSlice) // == nil
	// ([]int) <nil>
	spew.Dump(*new([]int)) // == nil

	// (bool) true
	spew.Dump(reflect.DeepEqual([]int{}, make([]int, 0)))

	// ([]int) {}
	spew.Dump(make([]int, 0)) // it's not nil, it's actually an empty slice
}

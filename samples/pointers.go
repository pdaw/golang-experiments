package samples

type Structure struct{}

// EXAMPLES FOR FUNCTIONS

func AcceptOnlyStructure(s Structure)           {}
func AcceptOnlyPointerToStructure(s *Structure) {}

func Run6Functions() {
	AcceptOnlyStructure(Structure{})

	// this cannot be compiled
	// AcceptOnlyStructure(&Structure{})

	AcceptOnlyPointerToStructure(&Structure{})

	// this cannot be compiled
	//AcceptOnlyPointerToStructure(Structure{})
}

// EXAMPLES FOR METHODS

func (s Structure) AcceptStructureOrPointerToStructure()        {}
func (s *Structure) AcceptStructureOrPointerToStructureAsWell() {}

func Run6Methods() {
	s := Structure{}

	s.AcceptStructureOrPointerToStructure()
	(&s).AcceptStructureOrPointerToStructure()

	s.AcceptStructureOrPointerToStructureAsWell()
	(&s).AcceptStructureOrPointerToStructureAsWell()
}

// EXAMPLES FOR INTERFACES

type Interface interface {
	Method(s Structure)
}

type StructureForInterface struct{}

func (si StructureForInterface) Method(s Structure) {}

func Run6InterfacesInterchangeability() {
	var i Interface
	si := StructureForInterface{}

	// either StructureForInterface or a pointer of StructureForInterface may be used as the Interface
	i = &si
	i.Method(Structure{})
	i = si
	i.Method(Structure{})
}

// ANOTHER EXAMPLE FOR INTERFACES

type InterfaceB interface {
	MethodB(s Structure)
}

type StructureForInterfaceB struct{}

func (si *StructureForInterfaceB) MethodB(s Structure) {}

func Run6InterfacesLackOfInterchangeabilityForPointers() {
	var i InterfaceB
	si := StructureForInterfaceB{}

	i = &si
	i.MethodB(Structure{})
	// this cannot be compiled
	// i = si
	// i.MethodB(Structure{})
}

func Run6() {
	Run6Functions()
	Run6Methods()
	Run6InterfacesInterchangeability()
	Run6InterfacesLackOfInterchangeabilityForPointers()
}

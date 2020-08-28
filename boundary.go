package main

type Boundary int

type BoundaryType int

const (
	BoundaryTypeUpper BoundaryType = iota
	BoundaryTypeExact
	BoundaryTypeLower
)

func BoundaryUpper(x int) Boundary {
	return Boundary(x << 2) - 1
}

func BoundaryExact(x int) Boundary {
	return Boundary(x << 2)	
}

func BoundaryLower(x int) Boundary {
	return Boundary(x << 2) + 1
}

func (b Boundary) ForceUpper() Boundary {
	return b.ForceExact() - 1
}

func (b Boundary) ForceExact() Boundary {
	return (b+1) & ^3
}

func (b Boundary) ForceLower() Boundary {
	return b.ForceExact() + 1
}

func (b Boundary) Value() int {
	return int(b.ForceExact() / 4)
}

func (b Boundary) IsUpper() bool {
	return (b%4 + 4) % 4 == 3
}

func (b Boundary) IsExact() bool {
	return (b%4 + 4) % 4 == 0
}

func (b Boundary) IsLower() bool {
	return (b%4 + 4) % 4 == 1
}

func (b Boundary) Type() BoundaryType {
	switch (b%4 + 4) % 4 {
	case 3:
		return BoundaryTypeUpper
	case 0:
		return BoundaryTypeExact
	case 1:
		return BoundaryTypeLower
	default:
		panic("invlid boundary")
	}
}
package ocp

type Size int
type Color int

const (
	Red Color = iota
	Blue
	Green
)

const (
	Small Size = iota
	Medium
	Large
)

type Product struct {
	Name  string
	Size  Size
	Color Color
}

var SampleProducts = []Product{
	{Name: "apple", Size: Small, Color: Green},
	{Name: "tree", Size: Large, Color: Green},
	{Name: "shirt", Size: Medium, Color: Blue},
}

type Filter struct{}

// it becomes difficult to add different filters in this way
// we are modifying the filter type if we add new methods to it
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	var result []*Product
	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

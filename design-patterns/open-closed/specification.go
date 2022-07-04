package ocp

type Specification interface {
	isSatisfied(*Product) bool
}

type ColorSpecification struct {
	Color Color
}

func (c ColorSpecification) isSatisfied(p *Product) bool {
	return c.Color == p.Color
}

type MultipleSpecification struct {
	Size  Size
	Color Color
}

func (m MultipleSpecification) isSatisfied(p *Product) bool {
	return m.Color == p.Color && m.Size == p.Size
}

// since we have a single filter having it standalone makes sense
func GeneralFilter(products []Product, spec Specification) []*Product {
	var result []*Product
	for i := range products {
		if spec.isSatisfied(&products[i]) {
			result = append(result, &products[i])
		}
	}
	return result
}

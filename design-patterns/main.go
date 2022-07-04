package main

import (
	// "example/design-patterns/single-responsibility"
	"example/design-patterns/open-closed"
	"fmt"
)

func ocp_execute() {
	myFilter := ocp.Filter{}
	products := ocp.SampleProducts
	for _, val := range myFilter.FilterByColor(products, ocp.Green) {
		fmt.Println(val.Name)
	}
}

func specification_execute() {
	// greenSpec := ocp.ColorSpecification{Color: ocp.Green}
	largeGreenSpec := ocp.MultipleSpecification{Color: ocp.Green, Size: ocp.Large}

	products := ocp.SampleProducts
	for _, val := range ocp.GeneralFilter(products, largeGreenSpec) {
		fmt.Println(val.Name)
	}
}

func main() {
	fmt.Println("Main function was run")
	specification_execute()
}

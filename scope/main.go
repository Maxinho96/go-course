package main

import (
	"scope/packageone"
)

var myVar = "This is a package level variable in main"

func main() {
	blockVar := "This is a block level variable"

	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}

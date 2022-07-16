package packageone

import "fmt"

var PackageVar = "This is a package level variable in packageone"

func PrintMe(strings ...string) {
	for _, s := range strings {
		fmt.Println(s)
	}
}

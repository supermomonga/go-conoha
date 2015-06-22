package main

import (
	"./identity"
	"fmt"
)

func main() {
	versions := identity.GetVersions()
	fmt.Println(versions)
}

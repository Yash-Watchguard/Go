package main

import (
	"fmt"
)

func main() {
	mp := make(map[string][]string)

	mp["bond_james"] = []string{"shaken", "notstirred", "martinis"}
	mp["bond_yash"] = []string{"shaken", "notstirred", "martinis"}
	mp["bond_manav"] = []string{"shaken", "notstirred", "martinis"}
	fmt.Println("before deletion\n")
	fmt.Printf("%#v", mp)

	delete(mp, "bond_manav")

	fmt.Println(fmt.Println("afterhhh"))
	fmt.Printf("%#v", mp)

	// for key, val := range mp {
	// 	fmt.Printf("%s have given choices\n", key)
	// 	for _, v := range val {
	// 		fmt.Printf("1.%s\n", v)
	// 	}
	// 	fmt.Println("\n")
	// }

}

package main

import "fmt"

func slice() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// Length
	fmt.Println(len(loons))

	// 0 indexing loons[1] = daffy
	// slices
	fmt.Println(loons[1:]) // [daffy, taz]

	// ========= Array loop =========

	// traditional for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// Single value range
	for i := range loons {
		fmt.Println(loons[i])
	}

	// Double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	// Ignore index using value range
	for _, name := range loons { // unused vars are compiratin errorrs, _ is needed
		fmt.Println(name)
	}
}

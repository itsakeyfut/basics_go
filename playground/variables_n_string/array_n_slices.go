package variables_n_string

import "fmt"

func ArrayAndSlices() {
	var ageArr1 [3]int = [3]int{20, 25, 30}
	var ageArr2 = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ageArr1, len(ageArr1))
	fmt.Println(ageArr2, len(ageArr2))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 1999
	scores = append(scores, 3294)
	fmt.Println(scores)

	// slice ranges
	r1 := names[1:3]
	r2 := names[2:]
	r3 := names[:3]
	fmt.Println(r1, r2, r3)
}
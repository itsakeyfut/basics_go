package variables_n_string

import (
	"fmt"
	"sort"
	"strings"
)

func Stl() {
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Index(greeting, "t"))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	ages := []int{46, 43, 65, 23, 2, 34, 12, 56, 27}

	sort.Ints(ages)
	fmt.Println(ages)

	idx := sort.SearchInts(ages, 43)
	fmt.Println(idx)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}
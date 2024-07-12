package main

import "fmt"

func intersection(set_1, set_2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	for key := range set_1 {
		if set_2[key] {
			result[key] = true
		}
	}

	return result
}

func main() {
	set_1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set_2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	intersected := intersection(set_1, set_2)

	fmt.Println(intersected)
}

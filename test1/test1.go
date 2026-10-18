package test1

import "fmt"

func Test1() {

	a := []int{1}     // slice len 1 cap 1
	b := append(a, 2) // new b 1, 2 len 2 cap 2
	b[0] = 0          // b 0,2 len 2, cap 2

	fmt.Println(a) // 1
	fmt.Println(b) // 0, 2

	c := append(b, 3) // new c 0, 2, 3 len 3, cap 4
	d := append(c, 4) // d 0, 2, 3, 4 len cap 4

	c[0] = 1 // 1, 2, 3

	fmt.Println(b) // 0, 2
	fmt.Println(c) // 1, 2, 3
	fmt.Println(d) // 1, 2, 3, 4

}

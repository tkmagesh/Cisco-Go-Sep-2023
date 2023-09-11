/* complex type */

package main

import "fmt"

func main() {
	var c1 complex128
	c1 = 4 + 5i
	fmt.Println(c1)

	c2 := 9 + 7i
	fmt.Println(c2)

	c3 := c1 + c2
	fmt.Println(c3)
}

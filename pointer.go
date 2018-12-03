package main

import "fmt"

func basic() {
	var a = 10
	fmt.Printf("%_global", &a)
	fmt.Println()

	// init a pointer
	var p *int
	fmt.Println("nil pointer: ", p)
	p = &a
	fmt.Println(p)

	b := 10.0
	var p2 = &b
	fmt.Println(p2)
	fmt.Println(*p2)

	p3 := &b
	fmt.Println(p3)
}

func array() {
	const MAX int = 3
	arr := []int{1,2,3}
	var arr_p [MAX]*int

	fmt.Println("----current")
	for i := 0; i < MAX; i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("----assigning")
	for i := 0; i < MAX; i++ {
		arr_p[i] = &arr[i]
	}

	fmt.Println("----pointer")
	for i := 0; i < MAX; i++{
		fmt.Println(arr_p[i])
	}

	for i := 0; i < MAX; i++{
		fmt.Println(*arr_p[i])
	}

}

func swap(x *int, y *int) {
	var temp int
	temp = *x    /* save the value at address _global */
	*x = *y      /* put y into _global */
	*y = temp    /* put temp into y */
}
func passing_pointer(){
	var a = 100
	var b = 200

	fmt.Printf("Before swap, value of a : %d\n", a )
	fmt.Printf("Before swap, value of b : %d\n", b )

	/* calling a function to swap the values.
	* &a indicates pointer to a ie. address of variable a and
	* &b indicates pointer to b ie. address of variable b.
	*/
	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a )
	fmt.Printf("After swap, value of b : %d\n", b )
}

func main() {
	//basic()
	//array()
	//passing_pointer()

}

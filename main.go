package main

import (
	"fmt"
)

func main() {

	fmt.Println("fix me plz")

	fmt.Println("hello world!!")

	ten := 103_500_999

	fll := 23.0e23

	andIntZ := 10
	var anIntS = 10
	var anInt int = 10
	fmt.Println(anInt, anIntS, andIntZ)

	fmt.Println(ten)

	fmt.Println(fll)

	var grt = "Greatings \n and \"Salutations\"!"
	var grt2 = `Greatings
	and "Salutations!`

	fmt.Println(grt)
	fmt.Println(grt2)

	var x = 10
	x *= 2

	fmt.Println(x)

	t, y, i := 10, 20, 30

	fmt.Println(t, y, i)

	q, w, e := 10, "hey", true

	fmt.Println(q, w, e)

	var (
		o        = 20
		p  int   = 30
		kl int64 = 8478475
		ui string
		uy string = "lalalala"
	)

	fmt.Println(o, p, kl, ui, uy)

	const xpto = 100

	fmt.Println(xpto)

	var _0 int = 45
	var π = 3.14

	fmt.Println(_0)
	fmt.Println(π)

	a, å := 10, 20

	fmt.Println(a, å)

	var xx = [...]int{1, 2, 3}
	yy := [3]int{1, 2, 3}
	yx := [...]int64{4, 5, 56, 7, 4, 4, 34, 3, 4, 34, 34, 3434}

	fmt.Println(xx)

	fmt.Println(yy)
	fmt.Println(yx)

	var matrix [2][3]int
	matrix[0][0] = 10
	matrix[0][1] = 10
	matrix[0][2] = 10
	matrix[1][0] = 1000
	matrix[1][1] = 1000
	matrix[1][2] = 1000
	fmt.Println(matrix)

	var xslice []int = nil
	var slice = []int{1, 2, 3, 4, 5, 6}
	var sliceRange = []int{1, 5: 5, 6, 20: 20, 100}
	var sliceMatrix [][]int

	fmt.Println("SLICES SECTION!")

	fmt.Println(xslice == nil)
	fmt.Println(slice)
	fmt.Println(sliceRange)
	fmt.Println(sliceMatrix)

	fmt.Println(len(slice))
	slice = append(slice, 10, 45, 56)
	slice = append(slice, []int{10, 45, 56}...)
	fmt.Println(len(slice))
	fmt.Println(slice)
	fmt.Println(len(sliceMatrix))
	fmt.Println(len(sliceRange))

	fmt.Println(cap(slice))
	slice = append(slice, []int{10, 45, 56}...)
	slice = append(slice, []int{10, 45, 56, 10, 45, 56, 10, 45, 56, 10, 45, 56, 10, 45, 56, 10, 45, 56, 10, 45, 56, 10, 45, 56}...)
	slice = append(slice, []int{10, 45, 56}...)
	slice = append(slice, []int{10, 45, 56}...)
	fmt.Println(cap(slice))

	var correctSlice = make([]int, 10, 200)
	correctSlice = append(correctSlice,
		[]int{1, 1, 1, 1, 1, 11, 1, 1}...)
	fmt.Println(len(correctSlice), cap(correctSlice), correctSlice)

	var s string = "Hello 😆"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s, s2, s3, s4)

	var ar rune = 'a'
	var as string = string(ar)
	var bb byte = 'a'
	var as2 string = string(bb)
	fmt.Println(ar, as, bb, as2)

}

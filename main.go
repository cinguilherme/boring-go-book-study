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

}

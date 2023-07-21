package main

/* type calculator interface {
	sum() int
	mul() int
	countList() int
} */

type sliceCalculator struct {
	x []int
}

func (sc sliceCalculator) sum() int {
	value := 0
	for i := 0; i < len(sc.x); i++ {
		value += sc.x[i]
	}
	return value
}

func callSliceCalculator1() {
	a := sliceCalculator{
		x: []int{1, 2, 3, 4, 5},
	}
	print(a.sum())
}

/* func callSliceCalculator2(c calculator) {
	print(c.sum())
} */

func main() {
	callSliceCalculator1()
	//callSliceCalculator2(sliceCalculator{x : []int{2,3,4}})
}

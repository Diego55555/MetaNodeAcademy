package main

import "fmt"

func main() {
	num := int64(20)

	fmt.Println("num:", num)
	modifyInt(&num)
	fmt.Println("num:", num)

	numSplice := []int64{1, 2, 3, 4, 5}
	fmt.Println("numSplice:", numSplice)
	modifyIntSplice(&numSplice)
	fmt.Println("numSplice:", numSplice)
}

func modifyInt(num *int64) {
	*num += 10
}

func modifyIntSplice(numSplice *[]int64) {
	for i := 0; i < len(*numSplice); i++ {
		(*numSplice)[i] *= 2
	}
}

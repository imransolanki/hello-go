package main

import(
	"fmt"
	"math"
	"errors"
)

func main() {
/*	fmt.Println("Hello, World!") */
/*	x := 5
	y := 10
	sum := x + y
	fmt.Println(sum)

	if sum > 5 {
		fmt.Println("sum > 5")
	} else if sum < 5 {
		fmt.Println("sum < 5")
	} else {
		fmt.Println("Unknown")
	}*/

/*	var a[5] int
	a[0] = 1
	a[1] = 7
	fmt.Println(a)
*/

/*	a := [5]int{1,2,3,4,5}
	fmt.Println(a)

	slice := []int{11,22,33}
	fmt.Println(slice)

	slice = append(slice, 44)
	fmt.Println(slice)
*/

/*	vertices := make(map[string]int)
	vertices["one"] = 1
	vertices["two"] = 2
	fmt.Println(vertices)
	fmt.Println(vertices["two"])
	delete(vertices, "two")
	fmt.Println(vertices)*/

/*	for i:=0 ; i<5 ; i++ {
		fmt.Println(i)
	}
*/

/*	i:=0
	for i < 5 {
		fmt.Println(i)
		i++
	}
*/

/*	arr := []string{"imran","fatima","faizan"}
	for index,value := range arr {
		fmt.Println("index:",index,"value:",value)
	}
*/

/*	vertices := make(map[string]int)
	vertices["one"] = 1
	vertices["two"] = 2
	vertices["three"] = 3

	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}*/

/*	result := sum(10,20)
	fmt.Println(result)*/

/*	result, error := sqrt(-4)
	if error == nil {
		fmt.Println(result)		
	} else {
		fmt.Println(error)
	}*/

/*	imran := person {
		name: "Imran",
		age: 30,
	}

	fmt.Println(imran)
	fmt.Println(imran.age)
*/

/*	number := 1214
	inc(&number)
	fmt.Println(number)*/
}


func inc(number *int) {
	*number++
}

type person struct {
	name string
	age int
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative")
	} 

	return math.Sqrt(x), nil
}

func sum(x int, y int) int {
	return x + y
}
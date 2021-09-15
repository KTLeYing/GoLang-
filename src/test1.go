package main

import "fmt"

/*const(
	a string = "abc"
	b int = len(a)
	c = unsafe.Sizeof(a)
)*/

/*const (
	a = iota
	b
	c
	d = "ha"
	e
	f = 100
	g
	h = iota
	i
	j = false
	k = 5.5
	l = iota
	m
	n
)*/

/*const(
	i = 1 << iota
	j = 3 << iota
	k
	l
)
*/
func main() {
	/*const LENGTH int = 10
	const WIDTH = 5
	var area int
	var area1 = 8
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println("面积为: ", area)
	println(a, b, c)
	println(a, b)
	println(area1)
	fmt.Println("面积: ", area)
	fmt.Println(area1)*/

	//fmt.Println(a, b ,c)

	//fmt.Println(i, j, k, l)

	/*var a = 21
	var b = 10*/
	//var c int

	/*c = a + b
	fmt.Println(c)

	c = a - b
	fmt.Printf("%d", c)
	fmt.Println()

	c = a * b
	fmt.Println(c)

	c = a / b
	fmt.Println(c)

	c = a % b
	fmt.Println(c)

	a++
	fmt.Println(a)

	a--
	fmt.Println(a)*/

	/*if(a > b){
		fmt.Println("a > b")
	}
	if(a < b) {
		fmt.Println("a < b")
	}
	if (a == b) {
		fmt.Println("a == b")
	}
	*/

	/*var grade = "B"
	var mark = 60

	switch mark {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50, 60, 70: grade = "C"
	default: grade = "D"

	}

	switch {
	case grade == "A":
		fmt.Println("成绩优秀")
	case grade == "B", grade == "C":
		fmt.Println("成绩良好")
	case grade == "D":
		fmt.Println("成绩合格")
	case grade == "F":
		fmt.Println("成绩不及格")
	default:
		fmt.Println("成绩差")
	}*/

	/*var i, j int
	for i = 0; i < 10; i++ {
		for j = 0; j <= i; j++ {
			fmt.Println(i, j)
		}
		if i == 5{
			break
		}
	}*/

	/*var a  = 100
	var b = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("最大值是：%d",  ret)*/

    /*a, b := swap("Google", "Runboot")
	fmt.Println(a, b);*/

	/*var a = 100
	var b = 200
	fmt.Printf("a交换前：%d\n", a)
	fmt.Printf("b交换前：%d\n", b)

	swap(&a, &b)
	fmt.Printf("a交换后：%d\n", a)
	fmt.Printf("b交换后: %d\n", b)*/

	/*getSquare := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquare(9))*/


	/*nextNum := getSequence()

	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())

	nextNum1 := getSequence()
	fmt.Println(nextNum1())
	fmt.Println(nextNum1())
	fmt.Println(nextNum1())*/

  /*  var c1 Circle
	c1.radius = 10.0
	fmt.Println("圆的面积为：" , c1.getArea())
	*/

	var i, j, k int
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance1 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j = 0; j < 5; j++ {
		fmt.Printf("balance1[%d] = %f\n", j, balance1)
	}

	balance2 := [5]float32{1:2.0, 4:8.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance2[%d] = %f\n", k, balance2)
	}
}

/*type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}*/


/*func getSequence() func() int{
	i := 0
	return func() int {
		i += 1
		return i
	}
}
*/
/*func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}*/

/*func swap(x, y string) (string, string)  {
	return y, x
}
*/
/*func max(a int, b int) int {
	var result int
	if a > b {
		result = a
	}else {
		result = b
	}
	return result
}
*/

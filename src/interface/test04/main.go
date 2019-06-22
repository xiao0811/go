package main

import (
	"fmt"
	"math"
)

// 定义一个接口
type Shape interface {
	peri() float64
	area() float64
}

// 定义实现类: 三角形
type Triangle struct {
	a, b, c float64
}

type Circle struct {
	radius float64 // 半径
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2                                     // 半周长
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c)) // 三角形面积: 海伦公式
	return s
}

// 测试函数
func testShape(s Shape)  {
	fmt.Printf("周长: %f, 面积: %f\n", s.peri(), s.area())
}

func main() {
	/**
	多态:一个事物的多种形态
	go语言: 通过接口模拟多态性
	一个实现类的对象:
		看作是一个实现类类型：能够访问实现类中的方法和属性
		还可以看作是对应的接口类型：只能够访问接口中定义的方法

	接口的用法:
		用法一：一个函数如果接收接口类型作为参数，那么实际上可以传入该接口的任意实现类对象作为参数。
	    用法二：定义一个类型为接口类型，那么实际上可以赋值任意实现类的对象。
			如果定义了一个接口类型的容器，实际上该容器中可以存储任意的实现类对象。
	 */

	t1 := Triangle{3, 4, 5}
	//fmt.Println(t1.peri())
	//fmt.Println(t1.area())
	//fmt.Println(t1.a, t1.b, t1.c)

	var s1 Shape
	s1 = t1
	fmt.Println(s1.area())
	fmt.Println(s1.peri())

	c1 := Circle{4}
	fmt.Println(c1.area())
	fmt.Println(c1.peri())

	var s2 Shape = Circle{5}
	fmt.Println(s2.peri(), s2.area())

	testShape(s2)

	arr := [4]Shape{t1, s1, c1, s2}
	fmt.Println(arr)
}

package main

import "fmt"

func main() {
	addressOfStruct()

	//li := Lili{}
	//li.fmtPointer()
	//Lili{}.fmtPointer()// compile error

}

// Q1
type Struct1 struct {
	name     string
	position int
}

func addressOfStruct() {
	s := Struct1{
		name:     "s",
		position: 10,
	}

	fmt.Println("position:=", s.position)
	fmt.Println("position adress=", &s.position)
	fmt.Println("position adress=", (&s).position)
}

// Q2
type littleGirl struct {
	Name string
	Age  int
}

type girl *littleGirl

//func(this girl) changeName(name string){ // 指针不允许出现在接受器中
//	this.Name = name
//}

// Q3
type Lili struct{
	Name string
}

func (Lili *Lili) fmtPointer(){
	fmt.Println("poniter")
}

func (Lili Lili) fmtReference(){
	fmt.Println("reference")
}



package main

import (
	"errors"
	"fmt"
)

type Stack struct {

	Maxsize int
	Top int
	Array [5]int
}

func (s *Stack)Push(val int) (err error) {
	if s.Maxsize -1 == s.Top {
		fmt.Println("栈已满")
		return errors.New("栈已满")
	}
	s.Top++
	s.Array[s.Top] = val
	return nil
}

func (s *Stack)Pop() (val int,err error) {
	if s.Top < 0 {
		fmt.Println("栈已空")
		return -1,errors.New("栈已空")
	}
	val = s.Array[s.Top]
	s.Top--
	if s.Top < -1 {
		s.Top = -1
	}
	return val,nil
}

func (s *Stack)Show()  {
	for i:=s.Top;i>=0;i-- {
		fmt.Println(s.Array[i])
	}
}

func main() {
	s :=&Stack{
		Maxsize:5,
		Top:-1,
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	s.Show()

	val ,_:=s.Pop()
	 fmt.Println(val)
	val ,_=s.Pop()
	fmt.Println(val)
	val ,_=s.Pop()
	val ,_=s.Pop()
	val ,_=s.Pop()
	val ,_=s.Pop()
	val ,_=s.Pop()
	val ,_=s.Pop()

	s.Show()

	s.Push(10)
	s.Push(44)
	s.Show()

}
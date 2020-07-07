package stack

import "testing"

func TestNewArrayStack(t *testing.T) {
	s := NewArrayStack(2)
	newStackTest(s, t)
}

func TestArrayStackPushOne(t *testing.T) {
	s := NewArrayStack(2)
	pushOneTest(s, t)
}

func TestArrayStackPushTwo(t *testing.T) {
	s := NewArrayStack(2)
	pushTwoTest(s, t)
}

func TestArrayStackPushThree(t *testing.T) {
	s := NewArrayStack(2)
	pushThreeTest(s, t)
}

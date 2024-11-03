package stack

import (
	"testing"
)

func TestNewListStack(t *testing.T) {
	s := NewListStack()
	newStackTest(s, t)
}

func TestListStackPushOne(t *testing.T) {
	s := NewListStack()
	pushOneTest(s, t)
}

func TestListStackPushTwo(t *testing.T) {
	s := NewListStack()
	pushTwoTest(s, t)
}

func TestListStackLock(t *testing.T) {
	s := NewListStack()
	lockTest(s, t)
}

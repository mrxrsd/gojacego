package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	if s.Len() != 0 {
		t.Errorf("lenght should be 0")
	}

	s.Push("1")
	s.Push("2")

	if s.Len() != 2 {
		t.Errorf("lenght should be 2")
	}

	top := s.Peek()

	if top.(string) != "2" {
		t.Errorf("value should be 2")
	}

	if s.Len() != 2 {
		t.Errorf("lenght should be 2")
	}

	item := s.Pop()
	if item.(string) != "2" {
		t.Errorf("value should be 2")
	}

	if s.Len() != 1 {
		t.Errorf("lenght should be 1")
	}

	item2 := s.Pop()
	if item2.(string) != "1" {
		t.Errorf("value should be 1")
	}

	nilItem := s.Pop()
	if nilItem != nil {
		t.Errorf("value should be nil")
	}

}

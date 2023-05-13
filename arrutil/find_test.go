package arrutil

import "testing"

func TestFindPositive(t *testing.T) {
	src := []int{1, 2, 3}
	expected := 2
	fn := func(val interface{}) bool {
		return val == expected
	}
	result, err := Find[int](src, fn)

	if err != nil {
		t.Errorf("Expected nil, but got error: %s", err.Error())
	}

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindNegative(t *testing.T) {
	src := []int{}
	fn := func(val interface{}) bool {
		return val == 2
	}
	_, err := Find[int](src, fn)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

// 测试 "Find" 函数参数是struct类型的切片
func TestFindStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	src := []Person{
		{"Jack", 18},
		{"Tom", 19},
	}
	expected := Person{"Tom", 19}
	fn := func(val interface{}) bool {
		return val.(Person).Name == expected.Name
	}
	result, err := Find[Person](src, fn)

	if err != nil {
		t.Errorf("Expected nil, but got error: %s", err.Error())
	}

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

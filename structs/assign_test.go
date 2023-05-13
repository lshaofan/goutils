package structs

import "testing"

func TestAssign(t *testing.T) {
	// Positive Test Case
	type SourceStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	type TargetStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	source := SourceStruct{
		Name: "John",
		Age:  20,
	}
	target := &TargetStruct{}
	err := Assign(source, target)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if target.Name != source.Name || target.Age != source.Age {
		t.Errorf("Expected target to be equal to source, got %v", target)
	}
	// Negative Test Case
	type SourceStruct2 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	type TargetStruct2 struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}
	source2 := SourceStruct2{
		Name: "John",
		Age:  20,
	}
	target2 := make(map[string]interface{})
	err = Assign(source2, target2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

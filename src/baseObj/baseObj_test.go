package baseObj

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestObjectToMap(t *testing.T) {
	t.Run("Valid Object", func(t *testing.T) {
		validObj := &TestStruct{ID: 1, Name: "John"}
		result, err := ObjectToMap(validObj)
		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}
		expectedResult := map[string]interface{}{"id": 1, "name": "John"}
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected %v, but got %v", expectedResult, result)
		}
	})
	t.Run("Invalid Object (Nil Pointer)", func(t *testing.T) {
		invalidObj := (*TestStruct)(nil)
		_, err := ObjectToMap(invalidObj)
		if err == nil {
			t.Errorf("Expected an error but got nil")
		}
	})
}

func TestMapToObject(t *testing.T) {
	data := map[string]interface{}{"id": 2, "name": "Alice"}
	result, err := MapToObject[TestStruct](data)
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	expectedResult := TestStruct{ID: 2, Name: "Alice"}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestIsRegexOkay(t *testing.T) {
	t.Run("Valid Input", func(t *testing.T) {
		input := "example123"
		pattern := `^[a-zA-Z0-9]+$`
		result := IsRegexOkay(input, pattern)
		if !result {
			t.Errorf("Expected true but got false")
		}
	})
	t.Run("Invalid Input (Contains Space)", func(t *testing.T) {
		invalidInput := "example 123"
		pattern := `^[a-zA-Z0-9]+$`
		result := IsRegexOkay(invalidInput, pattern)
		if result {
			t.Errorf("Expected false but got true")
		}
	})
}

func TestIsStringSizeOkay(t *testing.T) {
	t.Run("Valid Input (Within Size Limits)", func(t *testing.T) {
		input := "example"
		minSize := 3
		maxSize := 10
		result := IsStringSizeOkay(input, minSize, maxSize)
		if !result {
			t.Errorf("Expected true but got false")
		}
	})
	t.Run("Invalid Input (Below Minimum Size)", func(t *testing.T) {
		invalidInput := "ex"
		minSize := 3
		maxSize := 10
		result := IsStringSizeOkay(invalidInput, minSize, maxSize)
		if result {
			t.Errorf("Expected false but got true")
		}
	})
	t.Run("Invalid Input (Exceeds Maximum Size)", func(t *testing.T) {
		invalidInput := "thisisaverylongstring"
		minSize := 3
		maxSize := 10
		result := IsStringSizeOkay(invalidInput, minSize, maxSize)
		if result {
			t.Errorf("Expected false but got true")
		}
	})
}

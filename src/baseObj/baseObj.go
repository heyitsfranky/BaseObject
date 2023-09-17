package baseObj

import (
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
)

func ObjectToMap(obj interface{}) (map[string]interface{}, error) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr || objValue.IsNil() {
		return nil, errors.New("input must be a non-nil pointer to a struct")
	}
	objValue = objValue.Elem()
	objType := objValue.Type()
	result := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldName := field.Tag.Get("json")
		if fieldName == "" {
			fieldName = field.Name
		}
		fieldValue := objValue.Field(i).Interface()
		result[fieldName] = fieldValue
	}
	return result, nil
}

func MapToObject[T any](data map[string]interface{}) (T, error) {
	var elem T
	jsonData, err := json.Marshal(data)
	if err != nil {
		return elem, err
	}
	err = json.Unmarshal(jsonData, &elem)
	if err != nil {
		return elem, err
	}
	return elem, nil
}

func IsRegexOkay(input string, pattern string) bool {
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(input)
}

func IsStringSizeOkay(input string, minSize int, maxSize int) bool {
	return len(input) > minSize && len(input) < maxSize
}

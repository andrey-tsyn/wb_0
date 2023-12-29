package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func FillStructFromJSON(jsonData []byte, obj *interface{}) error {
	var m map[string]interface{}
	var missingFields []string
	hasMissing := false

	if err := json.Unmarshal(jsonData, &m); err != nil {
		return err
	}

	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		if value, ok := m[jsonTag]; ok {
			fieldValue := objValue.Field(i)
			if fieldValue.CanSet() {
				fieldValue.Set(reflect.ValueOf(value))
			}
		} else {
			hasMissing = true
			missingFields = append(missingFields, jsonTag)
		}
	}

	if hasMissing {
		return errors.New(fmt.Sprintf("Missing : %s", strings.Join(missingFields, ", ")))
	}

	return nil
}

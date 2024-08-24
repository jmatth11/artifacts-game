package helpers

import (
	"fmt"
	"reflect"
)

func MapFromStruct(v interface{}) map[string]string {
  t := reflect.TypeOf(v)
  tv := reflect.ValueOf(v)
  N := t.NumField()
  result := make(map[string]string)
  for i := 0; i < N; i++ {
    field := t.Field(i)
    fieldValue := tv.Field(i)
    jsonName, ok := field.Tag.Lookup("json")
    if ok {
      fieldValueStr := fmt.Sprintf("%v", fieldValue)
      if fieldValueStr != "" {
        result[jsonName] = fieldValueStr
      }
    }
  }
  return result
}

func JsonTagValue(v interface{}, fieldName string) (string, bool) {
  t := reflect.TypeOf(v)
  sf, ok := t.FieldByName(fieldName)
  if !ok {
      return "", false
  }
  return sf.Tag.Lookup("json")
}

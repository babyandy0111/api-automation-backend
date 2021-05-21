package helper

import (
	"reflect"
	"strconv"
	"time"
)

/*
This function will help you to convert your object from struct to map[string]interface{} based on your JSON tag in your structs.
Example how to use posted in sample_test.go file.
*/
func StructToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()

		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				v, ok := field.(time.Time)
				if ok {
					res[tag] = v
				} else {
					res[tag] = StructToMap(field)
				}
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

func MapStrStr2Struct(item interface{}, m map[string]string) error {
	reflectType := reflect.TypeOf(item)

	reflectValue := reflect.Indirect(reflect.ValueOf(item))

	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}

	for i, n := 0, reflectType.NumField(); i < n; i++ {
		tag := reflectType.Field(i).Tag.Get("json")

		if tag == "" || tag == "-" {
			tag = reflectType.Field(i).Name
		}
		switch reflectType.Field(i).Type.Kind() {
		// string to int64
		case reflect.Int64:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					v, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).SetInt(v)
				}
			}
		// string to int
		case reflect.Int:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					v, err := strconv.Atoi(value)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).SetInt(int64(v))
				}
			}
		// string to float32
		case reflect.Float32:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					v, err := strconv.ParseFloat(value, 32)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).SetFloat(v)
				}
			}
		// string to float64
		case reflect.Float64:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					v, err := strconv.ParseFloat(value, 64)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).SetFloat(v)
				}
			}
		// string to string
		case reflect.String:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					reflectValue.FieldByName(reflectType.Field(i).Name).SetString(value)
				}
			}
		// string to time
		case reflect.Struct:
			if _, ok := reflectValue.FieldByName(reflectType.Field(i).Name).Interface().(time.Time); ok {
				if value, ok := m[tag]; ok {
					moment, err := time.Parse(time.RFC3339, value)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).Set(reflect.ValueOf(moment))
				}
			}
		// string to bool
		case reflect.Bool:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					v, err := strconv.ParseBool(value)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).SetBool(v)
				}
			}
		// string to slice
		case reflect.Slice:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					reflectValue.FieldByName(reflectType.Field(i).Name).SetBytes([]byte(value))
				}
			}
		}
	}

	return nil
}

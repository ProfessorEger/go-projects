package main

import (
	"encoding/json"
	"os"
	"reflect"
)

func load_constants_from_file(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data map[string]interface{}
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	for key, value := range data {
		set_constant_value(key, value)
	}

	return nil
}

func set_constant_value(name string, value interface{}) {
	constant := reflect.ValueOf(&name).Elem()

	switch constant.Kind() {
	case reflect.String:
		constant.SetString(value.(string))
	case reflect.Int32:
		constant.SetInt(int64(value.(float64)))
	}
}

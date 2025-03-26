package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type NoRequiredTagSample struct {
	Name        string
	Description string
}

func IsValid(data any) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func PackageReflect() {
	sample := Sample{"Dims"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")
	max := structField.Tag.Get("max")

	fmt.Println(sampleType.NumField())
	fmt.Println(structField.Name)
	fmt.Println(required)
	fmt.Println(max)

	fmt.Println(IsValid(sample))
	sample.Name = ""
	fmt.Println(IsValid(sample))

	sample2 := NoRequiredTagSample{"", ""}
	fmt.Println(IsValid(sample2))
}

package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (nf *notFoundError) Error() string {
	return nf.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "Dims" {
		return &notFoundError{"not found"}
	}

	return nil
}

func CustomError() {
	err := SaveData("Dimss", nil)

	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}

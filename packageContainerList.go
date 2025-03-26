package main

import (
	"container/list"
	"fmt"
)

func PackageContainerList() {
	data := list.New()
	data.PushBack("Dims")
	data.PushBack("Raka")
	data.PushBack("D")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}

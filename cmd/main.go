package main

import (
	"fmt"

	auth "github.com/roughyy/Jubelio-Autoinput/services/auth"
	category "github.com/roughyy/Jubelio-Autoinput/services/category"
)

func main() {
	result, err := auth.GetToken()
	if err != nil {
		panic(err)
	}

	// categoryInfo, err := category.GetCategoryInfo(1, result)
	// if err != nil {
	// 	panic(err)
	// }	
	// fmt.Println(categoryInfo)

	allCategory, err := category.GetAllCategory(result)
	if err != nil {	
		panic(err)
	}
	fmt.Println(allCategory)
}
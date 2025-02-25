package category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetCategoryInfo(categoryId int, token string) (result string, err error) {
	url := "https://api2.jubelio.com/inventory/categories/item-categories/information/"+ strconv.Itoa(categoryId) +"/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer " + token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	fmt.Printf("Ini result %v",resp.Body)

	var resultMap map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&resultMap); err != nil {
		return "", err
	}

	fmt.Println(resultMap)
	return "", nil

}

func GetAllCategory(token string)(result string, err error){
	url:= "https://api2.jubelio.com/inventory/categories/item-categories/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("authorization", "Bearer " + token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	fmt.Printf("Ini result category all %v",resp.Body)

	var resultMap map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&resultMap); err != nil {
		return "", err
	}

	fmt.Println(resultMap)
	return "", nil
}
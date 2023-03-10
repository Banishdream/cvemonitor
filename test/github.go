package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.github.com/search/repositories?q=Sql%E6%B3%A8%E5%85%A5&sort=updated"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

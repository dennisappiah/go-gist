package gist

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(){
	client := http.Client{Timeout: time.Duration(1) * time.Second}

	request, err := http.NewRequest("GET","https://api.github.com/", nil)

	if err != nil{
		fmt.Printf("error %s", err)
        return
	}
	// adding request header
	request.Header.Add("Accept", `application/json`)

	response, err := client.Do(request)
    
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}

	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", response.Status)
}
package gist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type File struct {
    Content string `json:"content"`
}

type GistRequest struct{
	Files       map[string]File `json:"files"`
    Description string          `json:"description"`
    Public      bool            `json:"public"`
}

func CreateGist() {
	files := map[string]File{"main.go": {Content: "test"}}

	gist := GistRequest{
		Files: files,
		Description: "this is a test",
		Public: false,
	}

	// converting to json
	gistJson, err := json.Marshal(gist)

	client := http.Client{Timeout: time.Duration(1) * time.Second}

	req, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewBuffer(gistJson))

	if err != nil{
		fmt.Printf("%s", err)
		return
	}

	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("TOKEN")))


	resp, err := client.Do(req)

	if err != nil{
		fmt.Printf("%s", err)
		return
	}

	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)
}
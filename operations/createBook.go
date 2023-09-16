package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateBook(name string) {
	client := http.Client{}
	url := "http://127.0.0.1:8080/books"
	reqBody := Book{
		Name:       name,
		Author:     "Author name",
		Publisher:  "Publisher name",
		Origin:     "Book origin",
		TotalPages: 64,
		IsRead:     true,
	}
	marshalled, err := json.Marshal(reqBody)
	checkError(err)
	req, err := http.NewRequest("POST", url, bytes.NewReader(marshalled))
	req.Header.Add("Content-Type", "application/json")
	checkError(err)
	res, err := client.Do(req)
	checkError(err)
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	checkError(err)
	fmt.Println(string(body))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

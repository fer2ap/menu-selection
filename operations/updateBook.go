package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func UpdateBook(id int) {
	updatedBook := Book{
		Name:       "Updated name",
		Author:     "Updated Author name",
		Publisher:  "Updated Publisher name",
		Origin:     "Updated Book origin",
		TotalPages: 128,
		IsRead:     false,
	}

	marshalled, err := json.Marshal(updatedBook)
	checkError(err)

	base := "http://127.0.0.1:8080/books"
	url, err := url.Parse(base)
	checkError(err)

	url = url.JoinPath(fmt.Sprint(id))

	req, err := http.NewRequest("PUT", url.String(), bytes.NewReader(marshalled))
	checkError(err)

	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	checkError(err)

	fmt.Println(res.Status)
}

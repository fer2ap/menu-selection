package operations

import (
	"fmt"
	"net/http"
	"net/url"
)

func DeleteBookById(bookId int) {
	base := "http://127.0.0.1:8080/books"
	url, err := url.Parse(base)
	checkError(err)
	url = url.JoinPath(fmt.Sprint(bookId))
	req, err := http.NewRequest("DELETE", url.String(), nil)
	checkError(err)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	checkError(err)
	fmt.Println(res.Status)
}

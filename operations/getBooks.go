package operations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type BooksResponse struct {
	Books []Book `json:"books"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
}

func GetBookById(bookId int) {
	client := http.Client{}
	baseUrl := "http://127.0.0.1:8080/books/%d"
	formatedUrl := fmt.Sprintf(baseUrl, bookId)
	req, err := http.NewRequest("GET", formatedUrl, nil)
	checkError(err)
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()
	var book Book
	body, err := io.ReadAll(resp.Body)
	checkError(err)
	json.Unmarshal(body, &book)
	fmt.Println(book)
}

func GetBooks(page int, size int, name string) {
	// fmt.Printf("Looking for the term \"%s\", page %d, size %d \n", name, page, size)
	client := http.Client{}
	base, err := url.Parse("http://127.0.0.1:8080/books")
	checkError(err)
	params := url.Values{}
	params.Add("page", fmt.Sprint(page))
	params.Add("size", fmt.Sprint(size))
	params.Add("name", fmt.Sprint(name))
	// params.Add("author", "author")
	base.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", base.String(), nil)
	checkError(err)
	res, err := client.Do(req)
	checkError(err)
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	checkError(err)
	var booksResponse BooksResponse
	err = json.Unmarshal(body, &booksResponse)
	checkError(err)
	for _, v := range booksResponse.Books {
		fmt.Println(v)
	}

}

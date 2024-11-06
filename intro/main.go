wwpackage main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FactResponse struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

func main() {

	// 1. Buat Request
	req, err := http.NewRequest("GET", "https://cat-fact.herokuapp.com/facts/random", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 2. buat client
	client := http.Client{}

	// 3. pajnggil request dengna client
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// tutup response body

	defer res.Body.Close()

	// baca response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(resBody)

	// convert ke tipe data FastResponse
	var factResponse FactResponse
	err = json.Unmarshal(resBody, &factResponse)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("text", factResponse.Text)
	fmt.Println("type", factResponse.Type)

}

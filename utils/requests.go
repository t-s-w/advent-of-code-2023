package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Get(url string) []byte {
	Env()
	client := &http.Client{}
	cookie := os.Getenv("COOKIE")
	req, httpErr := http.NewRequest("GET", url, nil)
	if httpErr != nil {
		fmt.Printf("Error making request: %s", httpErr)
	}
	req.Header.Set("Cookie", cookie)
	res, resErr := client.Do(req)
	if resErr != nil {
		fmt.Printf("Error from response: %s", resErr)
	}
	defer res.Body.Close()
	body, bodyErr := io.ReadAll(res.Body)
	if bodyErr != nil {
		fmt.Print(bodyErr)
	}
	return body
}

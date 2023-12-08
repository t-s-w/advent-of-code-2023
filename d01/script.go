package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/joho/godotenv"
)

func main() {
	loadErr := godotenv.Load("../.env")
	if loadErr != nil {
		fmt.Printf("Error loading env: %s", loadErr)
		os.Exit(1)
	}
	cookie := os.Getenv("COOKIE")

	req, httpErr := http.NewRequest("GET", "https://adventofcode.com/2023/day/1/input", nil)
	if httpErr != nil {
		fmt.Printf("Error making request: %s", httpErr)
	}
	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	res, resErr := client.Do(req)
	if resErr != nil {
		fmt.Printf("Error from response: %s", resErr)
	}
	defer res.Body.Close()
	body, bodyErr := io.ReadAll(res.Body)
	if bodyErr != nil {
		fmt.Print(bodyErr)
	}
	strArr := strings.Split(string(body), "\n")
	output := 0
	for _, str := range strArr {
		lastDigit := -1
		for _, r := range []rune(str) {
			if unicode.IsDigit(r) {
				if lastDigit == -1 {
					output = output + int(r-'0')*10
				}
				lastDigit = int(r - '0')
			}
		}

		if lastDigit > -1 {
			output = output + lastDigit
		}
	}
	fmt.Print(output)

}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/devasherr/advent-of-code/day3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/3/input", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Cookie", "session="+os.Getenv("SESSION"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(day3.Solve(body))
}

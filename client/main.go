package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func calculate(d []byte) (float64, error) {
	resp, err := http.Post("http://localhost:8080/calculate", "application/json", bytes.NewBuffer(d))
	if err != nil {
		// If there is an error, return 0.0 and create a new error
		return 0.0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0.0, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
		return 0.0, err
	}

	type Operation struct {
		Operator string    `json:"operator"`
		Operands []float64 `json:"operands"`
		Result   float64   `json:"result"`
	}

	var op Operation

	err = json.Unmarshal(body, &op)

	return op.Result, nil
}

func main() {
	fmt.Println("Hello, Client!")

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", string(body))

	// Prepare data for the POST request
	data := []byte(`{"operator":"divide","operands":[2.0,0.0]}`)

	if result, err := calculate(data); err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Result:", result)
	}
}

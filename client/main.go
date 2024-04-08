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

var IP = "http://10.32.2.196:6969"

func calculate(d []byte) (float64, error) {
	resp, err := http.Post(IP+"/calculate", "application/json", bytes.NewBuffer(d))
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

	resp, err := http.Get(IP + "")
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

	type Operation struct {
		Operator string    `json:"operator"`
		Operands []float64 `json:"operands"`
		Result   float64   `json:"result"`
	}

	for {

		var op Operation

		fmt.Println("Enter the operator (add, subtract, multiply, divide), followed by the operands then exit character. Example: add 3 5 e")
		for {
			if n, err := fmt.Scan(&op.Operator); n == 0 || err != nil {
				fmt.Println("Error, retry:", err)
				continue
			}
			break
		}

		fmt.Println("Enter the operands:")
		var temp float64
		for {
			n, err := fmt.Scan(&temp)
			if n == 0 || err != nil {
				break
			}
			op.Operands = append(op.Operands, temp)

		}

		data, err := json.Marshal(op)

		// Convert the Operation struct to a JSON string
		if err != nil {
			fmt.Println("Error:", err)
			return
		} else {
			fmt.Println("Data:", string(data))
		}

		if result, err := calculate(data); err != nil {
			fmt.Println("Error:", err)
			return
		} else {
			fmt.Println("Result:", result)
		}
	}
}

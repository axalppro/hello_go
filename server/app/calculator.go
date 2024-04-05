package app

import "errors"

type Operation struct {
	Operator string    `json:"operator"`
	Operands []float64 `json:"operands"`
	Result   float64   `json:"result"`
}

type Calculator struct {
	pastOperations []Operation
}

func NewCalculator() *Calculator {
	return &Calculator{
		pastOperations: []Operation{},
	}
}

// Perform function
func (c *Calculator) Perform(operation Operation) (*Operation, error) {
	var result float64
	var err error

	switch operation.Operator {
	case "add":
		result, err = c.Add(operation.Operands)
	case "subtract":
		result, err = c.Subtract(operation.Operands)
	case "multiply":
		result, err = c.Multiply(operation.Operands)
	case "divide":
		result, err = c.Divide(operation.Operands)
	default:
		err = errors.New("invalid operator")
	}

	if err != nil {
		return nil, err
	}

	operation.Result = result
	c.pastOperations = append(c.pastOperations, operation)

	return &c.pastOperations[len(c.pastOperations)-1], nil

}

// Add function
func (c *Calculator) Add(v []float64) (float64, error) {
	result := v[0] + v[1]
	return result, nil
}

// Subtract function
func (c *Calculator) Subtract(v []float64) (float64, error) {
	result := v[0] - v[1]
	return result, nil
}

// Multiply function
func (c *Calculator) Multiply(v []float64) (float64, error) {
	result := v[0] * v[1]
	return result, nil

}

// Divide function
func (c *Calculator) Divide(v []float64) (float64, error) {
	if v[1] == 0 {
		return 0, errors.New("division by zero")
	}
	result := v[0] / v[1]
	return result, nil
}

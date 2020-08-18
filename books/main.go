package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func show() (*book, error) {
	bk, err := GetItem("978-0486298238")
	if err != nil {
		return nil, err
	}

	return bk, nil
}

func main() {
	lambda.Start(show)
}

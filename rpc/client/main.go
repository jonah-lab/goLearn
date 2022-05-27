package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	url := fmt.Sprintf("http://127.0.0.1:8000/add?a=%d&b=%b", a, b)
	res, _ := req.Get(url)
	body, _ := res.Body()
	fmt.Println(string(body))
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

func main() {
	print(Add(100, 1000))

}

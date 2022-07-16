package main

import (
	"encoding/base64"
	"fmt"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	var whatisit string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatisit = Reverse(string(sd))
	fmt.Println("The answer is", whatisit)
}

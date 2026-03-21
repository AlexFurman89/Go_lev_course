package main

import (
	"6_laba/task_1/server"
	"encoding/json"
	"fmt"
)

func toJSON(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil

}
func main() {
	sv := server.Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"127.0.0.1", "192.168.1.1"},
	}

	result, err := toJSON(sv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

package main

import "github.com/sobitada/go-jormungandr/wrapper"

func main() {
	api, err := wrapper.GetAPIFromHost("http://localhost:8080")
	if err == nil {
		var logs, err = api.GetLeadersLogs()
		if err == nil {
			print(">", logs)
		}
	}
}

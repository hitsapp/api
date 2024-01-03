package utils

import (
	"os"

	"github.com/imroc/req/v3"
)

func Backup() {
	client := req.C()
	items := GetKey("hits")

	if len(items) == 0 {
		return
	}

	println("Attempting to backup")
	resp, err := client.R().SetBody(map[string]interface{}{
		"content": items,
	}).SetHeader("Authorization", os.Getenv("IMPERIAL_API_TOKEN")).Post("https://api.impb.in/v1/documents")

	if err != nil {
		println("Error occurred whilst backing up", err)
	}

	s, err := resp.ToString()
	if resp.StatusCode != 200 {
		println("backup failed with ", resp.StatusCode, s)
	}

	println("Successfully backed up")
}

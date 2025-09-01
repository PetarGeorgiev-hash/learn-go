package maps

import "fmt"

func main() {
	websites := map[string]string{"Amazon Web Services": "https://aws.com",
		"Google": "https://gcp.com"}

	fmt.Println(websites["Google"])
}

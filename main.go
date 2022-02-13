package main

import "github.com/humbertovnavarro/kamaiitv-go/api"

func main() {
	go api.StartPublicApi(":3000")
	api.StartStreamApi(":3001")
}

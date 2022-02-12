package main

import "github.com/humbertovnavarro/kamaiitv-go/api"

func main() {
	go api.StartStreamApi(":3001")
}

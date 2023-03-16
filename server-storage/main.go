package main

import "server-test/server-storage/api"

func main() {
	api.GRPCSeverStorage("0.0.0.0:3003")
}

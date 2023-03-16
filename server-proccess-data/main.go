package main

import "server-test/server-proccess-data/api"

func main() {
	api.GRPCSeverStorage("0.0.0.0:3001")
}

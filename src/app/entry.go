package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	runMongoDB()
	runRedisDB()
}
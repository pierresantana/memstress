package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	size    = flag.Int("size", 1024, "size in MB")
	timeout = flag.Duration("timeout", time.Minute, "timeout to release the memory")
)

func main() {
	flag.Parse()
	if !flag.Parsed() {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sizeMB := (*size * 1000 * 1000) / 8

	log.Printf("Filling the memory with %d MB", *size)

	rand.Seed(time.Now().Unix())

	rand.Perm(sizeMB)

	log.Printf("Memory filled. Waiting %s", *timeout)

	time.Sleep(*timeout)

	log.Println("Releasing memory")
}

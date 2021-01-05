package main

import (
	"github.com/xxjwxc/gormt/data/cmd"
	"log"
)

// ./gormt -u root -p 1234 -d testDB -t cmdb_cdn_cache_fresh
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
}

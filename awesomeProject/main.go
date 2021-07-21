package main

import (
	"github.com/patrickmn/go-cache"
	"time"
)

func main(){
	var (
		defaultInterval = time.Minute * 1
		UserCache       = cache.New()
	)
}

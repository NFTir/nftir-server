package main

import (
	"NFTir/agent/initializers"
	"log"

	"github.com/jamespearly/loggly"
)

var (
	logglyClient *loggly.ClientType
)

func init()  {
	initializers.LoadEnvVars();
	logglyClient = loggly.New("NFTir")
}

func main()  {
	
	log.Println("NFTir")
}
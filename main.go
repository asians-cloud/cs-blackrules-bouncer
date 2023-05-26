package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/asians-cloud/cs-blackrules-bouncer/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

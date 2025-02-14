package main

import (
	"github.com/nurmuh-alhakim18/evermos-project/cmd"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
)

func main() {
	helpers.LoadConfig()

	helpers.LoadDatabase()

	helpers.LoadS3Session()

	cmd.ServeHTTP()
}

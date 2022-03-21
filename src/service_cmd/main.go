package main

import (
	"github.com/JavierR14/ratelimit/src/service_cmd/runner"
	"github.com/JavierR14/ratelimit/src/settings"
)

func main() {
	runner := runner.NewRunner(settings.NewSettings())
	runner.Run()
}

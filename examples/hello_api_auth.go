package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/api"
)

func main() {
	gbot := gobot.NewGobot()

	a := api.NewAPI(gbot)
	a.AddHandler(api.BasicAuth("gort", "klatuu"))
	a.Debug()

	a.AddHandler(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
	})
	a.Start()

	gbot.AddCommand("custom_gobot_command",
		func(params map[string]interface{}) interface{} {
			return "This command is attached to the mcp!"
		})

	hello := gbot.AddRobot(gobot.NewRobot("hello"))

	hello.AddCommand("hi_there", func(params map[string]interface{}) interface{} {
		return fmt.Sprintf("This command is attached to the robot %v", hello.Name)
	})

	gbot.Start()
}

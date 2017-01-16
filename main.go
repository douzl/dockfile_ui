package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/douzl/dockerfile_ui/src/router"
	"github.com/douzl/dockerfile_ui/src/router/middleware"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dockerfile_ui"
	app.Usage = "the UI to edit dockerfile"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "listen, l",
			Value: "5300",
			Usage: "Listen on",
		},
	}
	app.Action = func(c *cli.Context) error {
		server := &http.Server{
			Addr:           fmt.Sprintf(":%s", c.String("listen")),
			Handler:        router.Router(middleware.Authenticate),
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		fmt.Printf("services running on %s...", c.String("listen"))
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("http listen server error: %v", err)
		}
		return nil
	}

	app.Run(os.Args)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tapper"
	app.Version = Version
	app.Usage = ""
	app.Author = "handlename"
	app.Email = "nagata<at>handlena.me"
	app.Action = doMain
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "options",
			Value: "",
			Usage: "options for `go test` command",
		},
	}
	app.Run(os.Args)
}

func doMain(c *cli.Context) {
	var err error

	options := strings.Split(fmt.Sprintf("test -v %s", c.String("options")), " ")
	cmd := exec.Command("go", options...)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalf("error on create pipe: %s", err)
	}

	err = cmd.Start()

	if err != nil {
		log.Fatalf("error on run command: %s", err)
	}

	scanner := bufio.NewScanner(stdout)
	pattern, _ := regexp.Compile("^--- (PASS|FAIL): (.+)$")

	count := 0

	for scanner.Scan() {
		matches := pattern.FindStringSubmatch(scanner.Text())

		if 0 == len(matches) {
			continue
		}

		count++

		var result string
		if matches[1] == "PASS" {
			result = "ok"
		} else {
			result = "not ok"
		}

		fmt.Printf("%s %d - %s\n", result, count, matches[2])
	}

	fmt.Printf("%d..%d\n", 1, count)
}

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
)

func main() {
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
   {{if .Version}}
VERSION:
   {{.Version}}
   {{end}}{{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}
   {{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`
	usr, _ := user.Current()
	defaultDir := filepath.Join(usr.HomeDir, "Pictures/Wallpapers")
	var timeout int
	var directory string
	var random bool
	app := cli.NewApp()
	app.Name = "bgcarousel"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Gray Olson",
			Email: "gray@grayolson.com",
		},
	}
	app.Usage = "automatically rotate background image on a timer"
	app.Commands = nil
	app.ArgsUsage = ""
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "t, timeout",
			Value:       10,
			Usage:       "Specify timeout between image rotation in seconds",
			Destination: &timeout,
		},
		cli.StringFlag{
			Name:        "d, directory",
			Value:       defaultDir,
			Usage:       "Specify directory to search for images",
			Destination: &directory,
		},
		cli.BoolFlag{
			Name:        "random",
			Usage:       "selects a random image from the source directory instead of being cyclic",
			Destination: &random,
		},
	}
	app.Action = func(c *cli.Context) error {
		files, err := ioutil.ReadDir(directory)

		if err != nil {
			fmt.Println("Error opening directory: " + directory + ", ")
			return cli.NewExitError(err, 2)
		}

		var imgs []string
		for _, file := range files {
			if !file.IsDir() {
				imgs = append(imgs, file.Name())
			}
		}

		var i int
		if random {
			i = rand.Intn(len(imgs))
		} else {
			i = 0
		}

		for true {
			args := []string{"--bg-fill", filepath.Join(directory, imgs[i])}
			if err := exec.Command("feh", args...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return cli.NewExitError("There was an issue running feh", 1)
			}
			time.Sleep(time.Second * time.Duration(timeout))

			if random {
				i = rand.Intn(len(imgs))
			} else {
				i = (i + 1) % len(imgs)
			}
		}

		return nil
	}

	app.Run(os.Args)
}

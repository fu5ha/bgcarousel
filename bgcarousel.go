package main

import (
	"bytes"
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

func easeInOutQuad(t int, b int, c int, d int) int {
	t /= d / 2
	if t < 1 {
		return c/2*t*t + b
	}
	t--
	return -c/2*(t*(t-2)-1) + b
}

func bashExec(command string) (stdout string, stderr string, err error) {
	var stdo bytes.Buffer
	var stde bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdo
	cmd.Stderr = &stde
	err = cmd.Run()
	stdout = stdo.String()
	stderr = stde.String()
	return
}

func main() {
	usr, _ := user.Current()
	defaultDir := filepath.Join(usr.HomeDir, "Pictures/Wallpapers")
	var timeout int
	var directory string
	var random bool
	app := cli.NewApp()
	app.Name = "bgcarousel"
	app.Version = "0.1.0"
	app.Usage = "automatically rotate background image on a timer"
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

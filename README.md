# bgcarousel

`bgcarousel` is intended to be an easy-to-use, drop-in replacement for feh to be used for background switching. It will automatically
swap your background after a set amount of time, defaulting to 10 seconds.

# How to use

## Install from source

First, install Go and configure your environment. For example:

```sh
export GOPATH=$HOME/go
export PATH=$PATH:GOPATH/bin
```
Then,

```sh
$ go install -u github.com/termhn/bgcarousel
```
And finally in your i3 config, add:

```
exec_always --no-startup-id killall bgcarousel; bgcarousel --random
```

## Options
Use `bgcarousel -h` to print the help information, also for reference here:
```
$ bgcarousel -h
NAME:
   bgcarousel - automatically rotate background image on a timer

USAGE:
   bgcarousel [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -t value, --timeout value    Specify timeout between image rotation in seconds (default: 10)
   -d value, --directory value  Specify directory to search for images (default: "/home/amelie/Pictures/Wallpapers")
   --random                     selects a random image from the source directory instead of being cyclic
   --help, -h                   show help
   --version, -v                print the version
```


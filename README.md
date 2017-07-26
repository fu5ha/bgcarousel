# bgcarousel

`bgcarousel` is intended to be an easy-to-use, drop-in replacement for `feh` to be used for background switching. It will 
search for images in a directory (by default `~/Pictures/Wallpapers/`) and automatically swap your background after a set 
amount of time, defaulting to 10 seconds.

# How to use

## Install from source

First, install Go and configure your environment. For example:

```sh
$ sudo dnf install go
```
In bashrc or equivalent config file for other shells:
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
It also depends on feh to do the actual work behind the scenes, so install that as well.
```sh
$ sudo dnf install feh
```

With go installed, do

```sh
$ go get -u github.com/termhn/bgcarousel
```
And finally in your i3 config, add:

```
exec_always --no-startup-id killall bgcarousel; bgcarousel
```

## Options
Use `bgcarousel -h` to print the help information, also for reference here:
```
$ bgcarousel -h
NAME:
   bgcarousel - automatically rotate background image on a timer
   
VERSION:
   1.0.0
   
AUTHOR:
   Gray Olson <gray@grayolson.com>
   
USAGE:
   bgcarousel [global options]
   
GLOBAL OPTIONS:
   -t value, --timeout value    Specify timeout between image rotation in seconds (default: 10)
   -d value, --directory value  Specify directory to search for images (default: "/home/amelie/Pictures/Wallpapers")
   --random                     selects a random image from the source directory instead of being cyclic
   --help, -h                   show help
   --version, -v                print the version
```
For example, to make it choose a random image in `~/Pictures/Wallpapers/selects/` every 30 seconds, you could do
```
bgcarousel --random -t 30 -d ~/Pictures/Wallpapers/selects
```
# Todo
Handle errors that occur when removing images from source directory while running and support adding new images to source directory while running

# Possible future features
I want to work on a way to have a small fade transition between each image, but I'm not sure it will be possible 
without a rewrite using a different underlying program to set the background.


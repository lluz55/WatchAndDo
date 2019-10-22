package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/inancgumus/screen"
	"github.com/lluz55/scanfile"
)

var (
	file      string
	interval  int64
	waitExist bool
	command   string
)

func main() {
	screen.Clear()
	screen.MoveTopLeft()

	flag.StringVar(&file, "f", "", "Path to file that will be observed")
	flag.Int64Var(&interval, "i", 3, "Interval to watch file changes")
	flag.BoolVar(&waitExist, "w", false, "Wait until file be created when it doesn't exists")
	flag.StringVar(&command, "c", "", "Command that will be executed when file changes")

	flag.Parse()

	err := scanfile.WatchFile(file, &scanfile.WatchFileOpts{
		Interval:   interval,
		WaitExists: waitExist,
		Listener: func(msg string, changed bool) {
			log.Println(msg)
			if changed {
				cmds := strings.Split(command, " ")
				if len(cmds) < 1 {
					cmds = append(cmds, " ")
				}
				cmd := exec.Command(cmds[0], cmds[1:]...)
				err := cmd.Run()
				if err != nil {
					log.Println(err.Error())
					os.Exit(1)
				}
			}
		},
	})

	log.Println(err.Error())

}

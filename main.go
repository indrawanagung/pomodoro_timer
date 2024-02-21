package main

import (
	"flag"
	"fmt"
	"github.com/gen2brain/dlgs"
	"log"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/gosuri/uiprogress"
)

func main() {

	printHeader()
	workDuration := flag.Int("work", 30, "an int in minute")
	restDuration := flag.Int("rest", 30, "an int in minute")

	flag.Parse()
	workInterval := *workDuration * 60 // remind in minutes
	restInterval := *restDuration * 60 // remind in minutes

	uiprogress.Start()

	timeTicker := time.NewTicker(1 * time.Second)

	notifyChan := make(chan struct{})

	ticked := 0

	bar := uiprogress.AddBar(workInterval).AppendCompleted().PrependElapsed()

	timeType := "work"
	engine := timerEngine{
		NotifyChan:   &notifyChan,
		RestInterval: restInterval,
		WorkInterval: workInterval,
		ticker:       timeTicker,
		bar:          &bar,
		ticked:       ticked,
		timeType:     &timeType,
	}

	go engine.timeTicker()
	engine.notifying()

}

type timerEngine struct {
	NotifyChan   *chan struct{}
	WorkInterval int
	RestInterval int
	ticker       *time.Ticker
	bar          **uiprogress.Bar
	ticked       int
	timeType     *string
}

func (e timerEngine) notifying() {
	for {
		select {
		case <-*e.NotifyChan:
			if *e.timeType == "rest" {
				answer, err := dlgs.Question("Rest Time", "Do you want to rest ?", true)
				if err != nil {
					log.Fatal(err)
				}
				if !answer {
					log.Fatal("terminating")
				}
				*e.bar = uiprogress.AddBar(e.RestInterval).AppendCompleted().PrependElapsed()
				e.ticker.Reset(1 * time.Second)
			} else {
				answer, err := dlgs.Question("Work Time", "Do you want to work?", true)
				if err != nil {
					log.Fatal(err)
				}
				if !answer {
					log.Fatal("terminating")
				}
				*e.bar = uiprogress.AddBar(e.WorkInterval).AppendCompleted().PrependElapsed()
				e.ticker.Reset(1 * time.Second)
			}

		}
	}

}

func (e timerEngine) timeTicker() {
	for {
		select {
		case <-e.ticker.C:
			if *e.timeType == "work" {
				if e.ticked != e.WorkInterval {
					e.ticked++
					(*e.bar).Incr()
					continue
				}
				err := e.notify("It is time to rest")
				if err != nil {
					log.Fatal(err)
				}
				*e.timeType = "rest"
				*e.NotifyChan <- struct{}{}
				e.ticked = 0
				e.ticker.Stop()

			} else {
				if e.ticked != e.RestInterval {
					e.ticked++
					(*e.bar).Incr()
					continue
				}
				err := e.notify("It is time to work")
				if err != nil {
					log.Fatal(err)
				}
				*e.timeType = "work"
				*e.NotifyChan <- struct{}{}
				e.ticked = 0
				e.ticker.Stop()

			}
		}
	}
}

func (e timerEngine) notify(msg string) error {
	return beeep.Notify("Rest Time", msg, "")
}

func printHeader() {
	header := `
.-.----.                                                                                
\    /  \                      ____                                                     
|   :    \                   ..  . ..                ,---,
|   |  .\ :   .---.       .-+-,.' _ |   ,---.      ,---.'|   ,---.    __  ,-.   ,---.
.   :  |: |  '   ..\   .-+-. ;   . ||  '   ,'\     |   | :  '   ,'\ ,' ,'/ /|  '   ,'\
|   |   \ : /   /   | .--..|.   |  || /   /   |    |   | | /   /   |'  | |' | /   /   |
|   : .   /.   ; ,. :|   |  ,., |  |,.   ; ,. :  ,--.__| |.   ; ,. :|  |   ,'.   ; ,. :
;   | |.-. '   | |: :|   | /  | |--' '   | |: : /   ,'   |'   | |: :'  :  /  '   | |: : 
|   | ;    '   | .; :|   : |  | ,    '   | .; :.   '  /  |'   | .; :|  | '   '   | .; : 
:   ' |    |   :    ||   : |  |/     |   :    |'   ; |:  ||   :    |;  : |   |   :    | 
:   : :     \   \  / |   | |.-'       \   \  / |   | '/  ' \   \  / |  , ;    \   \  /  
|   | :      .----.  |   ;/            .----.  |   :    :|  .----.   ---'      .----.
.---'.|              '---'                      \   \  /                                
  .---.                                          .----.
🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅 🍅`
	fmt.Println(header)
}

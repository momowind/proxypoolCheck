package cron

import (
	"github.com/momowind/proxypoolCheck/config"
	"github.com/momowind/proxypoolCheck/internal/app"
	"github.com/jasonlvhit/gocron"
	"log"
	"runtime"
	"time"
)

func Cron() {
	_ = gocron.Every(15).Minutes().Do(appTask)
	<-gocron.Start()
}

func appTask() {
	err := config.Parse("")
	if err != nil{
		log.Printf("config parse error: %s\n", err.Error())
	}
	err = app.InitApp()
	if err != nil { // for wake up heroku
		log.Printf("init app err: %s\n Try in 2 minute\n", err.Error())
		time.Sleep(time.Minute*2)
		err = app.InitApp()
		if err != nil {
			log.Printf("crawl error: %s\n", err.Error())
		}
	}
	runtime.GC()
}

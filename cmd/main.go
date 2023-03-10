package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/router"
	"github.com/lixvyang/mixin-twitter/internal/utils"
)

func main() {
	signalch := make(chan os.Signal, 1)
	utils.InitSetting.Do(utils.Init)
	model.InitDB()

	router.InitRouter(signalch)

	//attach signal
	signal.Notify(signalch, os.Interrupt, syscall.SIGTERM)
	signalType := <-signalch
	signal.Stop(signalch)
	//cleanup before exit
	log.Printf("On Signal <%s>", signalType)
	log.Println("Exit command received. Exiting...")
}

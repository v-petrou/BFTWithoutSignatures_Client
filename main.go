package main

import (
	"BFTWithoutSignatures_Client/app"
	"BFTWithoutSignatures_Client/config"
	"BFTWithoutSignatures_Client/logger"
	"BFTWithoutSignatures_Client/messenger"
	"BFTWithoutSignatures_Client/variables"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// Initializer - Method that initializes all required processes
func initializer(id int, n int, scenario int, rem int) {
	variables.Initialize(id, n, rem)
	logger.InitializeLogger("./logs/client/", "./logs/client/")

	if variables.Remote {
		config.InitializeIP()
	} else {
		config.InitializeLocal()
	}
	config.InitializeScenario(scenario)

	logger.OutLogger.Print(
		"ID:", variables.ID, " | N:", variables.N, " | F:", variables.F,
		" | Scenario:", config.Scenario, " | Remote:", variables.Remote, "\n\n",
	)

	messenger.InitializeMessenger()
	messenger.Subscribe()
	messenger.TransmitRequests()

	app.Client()

	cleanup()
}

func cleanup() {
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for range terminate {
			for i := 0; i < variables.N; i++ {
				messenger.ServerSockets[i].Close()
				messenger.ResponseSockets[i].Close()
			}
			os.Exit(0)
		}
	}()
}

func main() {
	args := os.Args[1:]
	if len(args) == 4 {
		id, _ := strconv.Atoi(args[0])
		n, _ := strconv.Atoi(args[1])
		scenario, _ := strconv.Atoi(args[2])
		remote, _ := strconv.Atoi(args[3])

		initializer(id, n, scenario, remote)

		done := make(chan interface{}) // To keep the client running
		<-done

	} else {
		log.Fatal("Arguments should be '<id> <n> <scenario> <remote>'")
	}
}

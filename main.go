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
func initializer(id int, n int, clients int, scenario config.Scenario) {
	variables.Initialize(id, n, clients)

	if variables.Remote {
		config.InitializeIP()
	} else {
		config.InitializeLocal()
	}
	config.InitializeScenario(scenario)

	logger.InitializeLogger("./logs/client/", "./logs/client/")
	logger.OutLogger.Print(
		"ID:", variables.ID, " | N:", variables.N, " | Clients:", variables.Clients, "\n\n",
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
		clients, _ := strconv.Atoi(args[2])
		tmp, _ := strconv.Atoi(args[3])
		scenario := config.Scenario(tmp)

		initializer(id, n, clients, scenario)

		// To keep the client running
		done := make(chan interface{})
		<-done

	} else {
		log.Fatal("Arguments should be '<id> <n> <clients> <scenario>")
	}
}

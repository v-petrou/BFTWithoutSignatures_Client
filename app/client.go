package app

import (
	"BFTWithoutSignatures_Client/logger"
	"BFTWithoutSignatures_Client/messenger"
	"BFTWithoutSignatures_Client/types"
	"BFTWithoutSignatures_Client/variables"
	"log"
	"math/rand"
	"time"
)

func Client() {
	// START variables initialization
	runes := []rune("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")

	replies := make(map[int]map[int]bool) // id, from
	accepted := make(map[int]bool)        // if this id is accepted

	randS := rand.New(rand.NewSource(time.Now().UnixNano()))
	randR := rand.New(rand.NewSource(time.Now().UnixNano()))
	// END variables initialization

	// Request Sender
	go func() {
		for num := 1; num > 0; num++ {
			message := types.NewClientMessage(variables.ID, num, runes[randR.Intn(len(runes))])
			messenger.SendRequest(message, randS.Intn(variables.N))

			time.Sleep(5 * time.Second)
		}
	}()

	// Response Handler
	go func() {
		for message := range messenger.ResponseChannel {
			if _, in := replies[message.Value][message.From]; in {
				continue // Only one value can be received from each server
			}
			if replies[message.Value] == nil {
				replies[message.Value] = make(map[int]bool)
			}
			replies[message.Value][message.From] = true

			// Call countReplies and if more than f+1 with the same value, accept the array.
			if len(replies[message.Value]) >= (variables.F+1) && !accepted[message.Value] {
				accepted[message.Value] = true
				logger.OutLogger.Print("RECEIVED ACK for ", message.Value, "\n")
				log.Println(variables.ID, "|", "RECEIVED ACK for", message.Value)
			}
		}
	}()
}

// start := time.Now()
// log.Println(variables.ID, "|", "time-", time.Since(start))

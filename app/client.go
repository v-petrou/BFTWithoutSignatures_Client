package app

import (
	"BFTWithoutSignatures_Client/logger"
	"BFTWithoutSignatures_Client/messenger"
	"BFTWithoutSignatures_Client/types"
	"BFTWithoutSignatures_Client/variables"
	"math/rand"
	"time"
)

var (
	runes = []rune("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")

	replies  = make(map[int]map[int]bool) // id, from
	accepted = make(map[int]bool)         // if this id is accepted

	// Client metrics regarding the experiment evaluation
	sentTime  = make(map[int]time.Time)
	OpLatency = time.Duration(0)
	Rounds    = 0
)

func Client() {
	rand.Seed(int64((variables.ID + 3) * 9000)) // Pseudo-Random Generator

	time.Sleep(time.Duration(variables.ID) * time.Second) // Wait before sending 1st request
	go sendRune()

	go func() {
		for {
			timeout := time.NewTicker(45 * time.Second)
			var message types.Reply

			select {
			case message = <-messenger.ResponseChannel:
				if _, in := replies[message.Value][message.From]; in {
					continue // Only one value can be received from each server
				}
				if replies[message.Value] == nil {
					replies[message.Value] = make(map[int]bool)
				}
				replies[message.Value][message.From] = true

				// If more than f+1 with the same value, accept the array.
				if len(replies[message.Value]) >= (variables.F+1) && !accepted[message.Value] {
					accepted[message.Value] = true
					OpLatency += time.Since(sentTime[message.Value])
					logger.OutLogger.Print("RECEIVED ACK for ", message.Value, " [",
						time.Since(sentTime[message.Value]), "]\n")

					if Rounds < 2 {
						go sendRune()
					}
				}
			case <-timeout.C:
				if Rounds < 2 {
					logger.OutLogger.Println("ABORTING and resending", Rounds)
					Rounds--
					replies[message.Value] = nil
					go sendRune()
				}
			}
		}
	}()
}

func sendRune() {
	Rounds++
	sentTime[Rounds] = time.Now()

	message := types.NewClientMessage(variables.ID, Rounds, runes[rand.Intn(len(runes))])
	go messenger.SendRequest(message, rand.Intn(variables.N))
}

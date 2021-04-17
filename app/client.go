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

	sentTime     = make(map[int]time.Time)
	ReceivedTime = make(map[int]time.Duration)
	num          = 1
)

func Client() {
	rand.Seed(int64((variables.ID + 3) * 9000)) // Pseudo-Random Generator

	go sendRune()

	go func() {
		for message := range messenger.ResponseChannel {
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
				ReceivedTime[message.Value] = time.Since(sentTime[message.Value])
				logger.OutLogger.Print("RECEIVED ACK for ", message.Value, " [",
					ReceivedTime[message.Value], "]\n")

				if num <= 2 {
					go sendRune()
				}
			}
		}
	}()
}

func sendRune() {
	time.Sleep(time.Duration(variables.ID) * time.Second)

	sentTime[num] = time.Now()

	message := types.NewClientMessage(variables.ID, num, runes[rand.Intn(len(runes))])
	go messenger.SendRequest(message, rand.Intn(variables.N))

	num++
}

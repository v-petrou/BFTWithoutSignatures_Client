package app

import (
	"BFTWithoutSignatures_Client/messenger"
	"BFTWithoutSignatures_Client/variables"
	"bytes"
	"log"
	"time"
	"unicode/utf8"
)

func Client() {
	array := make([]rune, 0)

	replies := make(map[int]map[int][][]byte) // id, from
	ack := make(map[int]bool)

	go func() {
		for {
			if variables.ID == 0 {
				messenger.SendRequest('A', 4)
			} else if variables.ID == 1 {
				messenger.SendRequest('B', 8)
			}

			time.Sleep(30 * time.Second)
		}
	}()

	go func() {
		for message := range messenger.ResponseChannel {
			if _, in := replies[message.Id][message.From]; in {
				continue // Only one value can be received from each process
			}
			if replies[message.Id] == nil {
				replies[message.Id] = make(map[int][][]byte)
			}
			replies[message.Id][message.From] = message.Value

			// Count the replies with the same ID and if more than f+1 with the same value,
			// add this to the array.
			count, dict := countReplies(replies[message.Id])
			for k, v := range count {
				if v >= (variables.F+1) && !ack[message.Id] {
					ack[message.Id] = true
					for _, key := range dict[k] {
						val, _ := utf8.DecodeRune(key)
						array = append(array, val)
					}

					log.Println(variables.ID, "|", array)
				}
			}
		}
	}()
}

func countReplies(vector map[int][][]byte) (map[int]int, map[int][][]byte) {
	counter := make(map[int]int)
	dict := make(map[int][][]byte)
	for _, val := range vector {
		key := len(dict)
		for k, v := range dict {
			if len(v) != len(val) {
				continue
			}
			eq := true
			for i := range v {
				if !bytes.Equal(v[i], val[i]) {
					eq = false
					break
				}
			}
			if eq {
				key = k
				break
			}
		}
		dict[key] = val
		counter[key] = counter[key] + 1
	}

	return counter, dict
}

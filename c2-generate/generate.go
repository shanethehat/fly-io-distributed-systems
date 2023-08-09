package main

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/google/uuid"

	"flydistsys/common"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	messageCount := common.NewCounter()

	for scanner.Scan() {
		line := scanner.Bytes()

		var msg common.Message

		if err := json.Unmarshal(line, &msg); err != nil {
			panic(err)
		}

		var responseBody json.RawMessage

		switch msg.GetType() {
		case "init":
			var initBody common.InitMessageBody

			if err := json.Unmarshal(msg.Body, &initBody); err != nil {
				panic(err)
			}

			r, _ := json.Marshal(
				common.MessageBody{
					Type:      "init_ok",
					InReplyTo: initBody.MessageId})
			responseBody = r

		case "generate":
			var generateBody common.MessageBody

			if err := json.Unmarshal(msg.Body, &generateBody); err != nil {
				panic(err)
			}

			r, _ := json.Marshal(
				common.GenerateMessageBody{
					MessageBody: common.MessageBody{
						Type:      "generate_ok",
						MessageId: messageCount.IncrementAndRead(),
						InReplyTo: generateBody.MessageId},
					Id: uuid.NewString()})
			responseBody = r
		}

		response, _ := json.Marshal(
			common.Message{
				Src:  msg.Dest,
				Dest: msg.Src,
				Body: responseBody})

		os.Stdout.WriteString(string(response) + "\n")
	}
}

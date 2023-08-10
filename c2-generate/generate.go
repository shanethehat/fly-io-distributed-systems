package main

import (
	"encoding/json"

	"github.com/google/uuid"

	"flydistsys/common"
)

func main() {
	messageCount := common.NewCounter()
	node := common.NewNode()

	node.RegisterHandler("generate", func(n common.NodeInterface, message common.Message) error {
		var generateBody common.MessageBody

		if err := json.Unmarshal(message.Body, &generateBody); err != nil {
			return err
		}

		r, _ := json.Marshal(
			common.GenerateMessageBody{
				MessageBody: common.MessageBody{
					Type:      "generate_ok",
					MessageId: messageCount.IncrementAndRead(),
					InReplyTo: generateBody.MessageId},
				Id: uuid.NewString()})

		node.Send(message.Src, r)

		return nil
	})

	node.Start()
}

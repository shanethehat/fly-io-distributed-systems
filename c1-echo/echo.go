package main

import (
	"encoding/json"

	"flydistsys/common"
)

func main() {
	messageCount := common.NewCounter()
	node := common.NewNode()

	node.RegisterHandler("echo", func(n common.NodeInterface, message common.Message) error {
		var echoBody common.EchoMessageBody

		if err := json.Unmarshal(message.Body, &echoBody); err != nil {
			return err
		}

		r, _ := json.Marshal(
			common.EchoMessageBody{
				MessageBody: common.MessageBody{
					Type:      "echo_ok",
					MessageId: messageCount.IncrementAndRead(),
					InReplyTo: echoBody.MessageId},
				Echo: echoBody.Echo})

		n.Send(message.Src, r)

		return nil
	})

	node.Start()
}

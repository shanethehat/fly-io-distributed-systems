package main

import (
	"encoding/json"
	"slices"

	"flydistsys/common"
)

func main() {
	messageCount := common.NewCounter()
	node := common.NewNode()

	var messages = make([]int, 1)

	node.RegisterHandler("topology", func(n common.NodeInterface, message common.Message) error {
		var topologyBody common.TopologyMessageBody

		if err := json.Unmarshal(message.Body, &topologyBody); err != nil {
			return err
		}

		n.SetNeighbours(topologyBody.Topology)

		r, _ := json.Marshal(
			common.MessageBody{
				Type:      "topology_ok",
				MessageId: messageCount.IncrementAndRead(),
				InReplyTo: topologyBody.MessageId})

		node.Send(message.Src, r)

		return nil
	})

	node.RegisterHandler("broadcast", func(n common.NodeInterface, message common.Message) error {
		var broadcastBody common.BroadcastMessageBody

		if err := json.Unmarshal(message.Body, &broadcastBody); err != nil {
			return err
		}

		if !slices.Contains(messages, broadcastBody.Message) {
			messages = append(messages, broadcastBody.Message)

			m, _ := json.Marshal(
				common.BroadcastMessageBody{
					MessageBody: common.MessageBody{
						Type:      "broadcast",
						MessageId: messageCount.IncrementAndRead(),
						InReplyTo: 0,
					},
					Message: broadcastBody.Message,
				},
			)
			node.SendToNeighbours(m)
		}

		// only send a response if the origin is a client
		if message.Src[0] == 'c' {
			r, _ := json.Marshal(
				common.MessageBody{
					Type:      "broadcast_ok",
					MessageId: messageCount.IncrementAndRead(),
					InReplyTo: broadcastBody.MessageId})

			node.Send(message.Src, r)
		}

		return nil
	})

	node.RegisterHandler("read", func(n common.NodeInterface, message common.Message) error {
		var readBody common.MessageBody

		if err := json.Unmarshal(message.Body, &readBody); err != nil {
			return err
		}

		r, _ := json.Marshal(
			common.ReadMessageBody{
				MessageBody: common.MessageBody{
					Type:      "read_ok",
					MessageId: messageCount.IncrementAndRead(),
					InReplyTo: readBody.MessageId},
				Messages: messages})

		node.Send(message.Src, r)

		return nil
	})

	node.Start()
}

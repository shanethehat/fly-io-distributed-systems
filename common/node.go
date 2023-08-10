package common

import (
	"bufio"
	"encoding/json"
	"os"
)

type Handler func(node NodeInterface, message Message) error

type Node struct {
	nodeId       string
	clusterNodes []string

	handlers map[string]Handler
}

type NodeInterface interface {
	RegisterHandler(command string, handler Handler)
	Send(dest string, body json.RawMessage)
	Start()
}

func NewNode() NodeInterface {
	node := Node{}
	node.handlers = make(map[string]Handler)
	return &node
}

func (node *Node) RegisterHandler(command string, handler Handler) {
	node.handlers[command] = handler
}

func (node *Node) init(body InitMessageBody) {
	node.nodeId = body.NodeId
	node.clusterNodes = body.NodeIds
}

func (node *Node) Send(dest string, body json.RawMessage) {
	message, _ := json.Marshal(
		Message{
			Src:  node.nodeId,
			Dest: dest,
			Body: body})

	os.Stdout.WriteString(string(message) + "\n")
}

func (node *Node) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Bytes()

		var msg Message

		if err := json.Unmarshal(line, &msg); err != nil {
			panic(err)
		}

		messageType := msg.GetType()
		switch messageType {
		case "init":
			var initBody InitMessageBody

			if err := json.Unmarshal(msg.Body, &initBody); err != nil {
				panic(err)
			}

			node.init(initBody)

			r, _ := json.Marshal(
				MessageBody{
					Type:      "init_ok",
					InReplyTo: initBody.MessageId})

			node.Send(msg.Src, r)

		default:
			if handler := node.handlers[messageType]; handler != nil {
				handler(node, msg)
			}
		}
	}
}

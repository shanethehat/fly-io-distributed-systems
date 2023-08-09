package common

import (
	"encoding/json"
)

type Message struct {
	Src  string          `json:"src"`
	Dest string          `json:"dest"`
	Body json.RawMessage `json:"body"`
}

func (msg *Message) GetType() string {
	var body MessageBody
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		panic(err)
	}
	return body.Type
}

type MessageBody struct {
	Type      string `json:"type"`
	MessageId int    `json:"msg_id,omitempty"`
	InReplyTo int    `json:"in_reply_to,omitempty"`
}

type InitMessageBody struct {
	MessageBody
	NodeId  string   `json:"node_id"`
	NodeIds []string `json:"node_ids"`
}

type EchoMessageBody struct {
	MessageBody
	Echo string `json:"echo"`
}

package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageExtractsSource(t *testing.T) {
	input := `{
		"src": "s1",
		"dest": "d1",
		"body": ""
	}`

	var message Message
	json.Unmarshal([]byte(input), &message)

	assert.Equal(t, "s1", message.Src)
}

func TestMessageExtractsDestination(t *testing.T) {
	input := `{
		"src": "s1",
		"dest": "d1",
		"body": ""
	}`

	var message Message
	json.Unmarshal([]byte(input), &message)

	assert.Equal(t, "d1", message.Dest)
}

func TestMessageReturnsType(t *testing.T) {
	input := `{
		"src": "",
		"dest": "",
		"body": {
			"type": "t1",
			"msg_id": 1,
			"in_reply_to": 2
		}
	}`

	var message Message
	json.Unmarshal([]byte(input), &message)

	assert.Equal(t, "t1", message.GetType())
}

func TestMessageBodyExtractsType(t *testing.T) {
	input := `{
		"src": "",
		"dest": "",
		"body": {
			"type": "t1",
			"msg_id": 1,
			"in_reply_to": 2
		}
	}`

	var message Message
	json.Unmarshal([]byte(input), &message)

	var body MessageBody
	json.Unmarshal(message.Body, &body)

	assert.Equal(t, "t1", body.Type)
}

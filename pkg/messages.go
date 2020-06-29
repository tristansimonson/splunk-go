package splunk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// MessageSeverity contains the significance of a message in string form
type MessageSeverity string

// Message contains the name and content of the message in JSON
type Message struct {
	Name    string         `json:"name"`
	Content MessageContent `json:"content"`
}

// MessageContent contains the aggregation of MessageSeverity and Message
type MessageContent struct {
	Message  string          `json:"message"`
	Severity MessageSeverity `json:"severity"`
	Created  int64           `json:"timeCreated_epochSecs"`
}

// Content logs the time of the message
func (mc *MessageContent) Content() time.Time {
	return time.Unix(mc.Created, 0)
}

// Messages is the plural of the Message struct
type Messages struct {
	Origin   string    `json:"origin"`
	Messages []Message `json:"entry"`
}

const (
	// Info is a message of type informational
	Info MessageSeverity = "info"
	// Warn is a message of type warn
	Warn MessageSeverity = "warn"
	// Error is a message of type error
	Error MessageSeverity = "error"
)

// SendMessage sends an informational message to Splunk
func (conn Connection) SendMessage(message *Message) (string, error) {
	data := make(url.Values)
	data.Add("name", message.Name)
	data.Add("value", message.Content.Message)
	data.Add("severity", string(message.Content.Severity))
	response, err := conn.httpPost(fmt.Sprintf("%s/services/messages", conn.BaseURL), &data)
	return response, err
}

// GetMessage fetches and unmarshalls a message from Splunk
func (conn Connection) GetMessage(name string) ([]Message, error) {
	data := make(url.Values)
	data.Add("name", name)
	data.Add("output_mode", "json")
	response, err := conn.httpGet(fmt.Sprintf("%s/services/messages/%s", conn.BaseURL, name), &data)

	if err != nil {
		return []Message{}, err
	}

	bytes := []byte(response)
	var messages Messages
	unmarshallError := json.Unmarshal(bytes, &messages)
	return messages.Messages, unmarshallError
}

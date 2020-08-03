package splunk

import (
	"fmt"
	"net/url"
)

// RestartServer is used to send a process reboot request to the Splunk instance
func (conn *Connection) RestartServer() (string, error) {
	data := make(url.Values)
	data.Add("output_mode", "json")
	response, err := conn.httpPost(fmt.Sprintf("%s/services/server/control/restart", conn.BaseURL), &data)
	return response, err
}

// InspectControl is used to list actions that can be performed at the queried endpoint.
func (conn *Connection) InspectControl(endpoint string) (string, error) {
	data := make(url.Values)
	data.Add("output_mode", "json")
	response, err := conn.httpGet(fmt.Sprintf("%s/services/server/control/%s", conn.BaseURL, endpoint), &data)
	return response, err
}

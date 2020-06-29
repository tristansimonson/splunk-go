package splunk

import (
	"fmt"
)

// RestartServer performs a reboot operation Splunk
func (conn Connection) RestartServer() (string, error) {
	response, err := conn.httpPost(fmt.Sprintf("%s/services/server/control/restart", conn.BaseURL), nil)
	return response, err
}

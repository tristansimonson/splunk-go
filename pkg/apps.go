package splunk

import (
	"fmt"
	"net/url"
)

// InstallApp installs splunk apps
func (conn Connection) InstallApp(path string, update bool) (string, error) {
	data := make(url.Values)
	data.Add("name", path)

	updateApp := "false"
	if update == true {
		updateApp = "true"
	}

	data.Add("update", updateApp)
	response, err := conn.httpPost(fmt.Sprintf("%s/services/apps/appinstall/", conn.BaseURL), &data)
	return response, err
}

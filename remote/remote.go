package remote

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amonapp/amonagent/logging"
	"github.com/amonapp/amonagent/settings"
)

var remoteLogger = logging.GetLogger("amonagent.remote")

// DefaultTimeOut - 10 seconds
var DefaultTimeOut = 10 * time.Second

var tr = &http.Transport{
	ResponseHeaderTimeout: DefaultTimeOut,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, // for self-signed certificates
}

// SystemURL - XXX
func SystemURL() string {
	settings := settings.Settings()

	url := settings.AmonInstance + "/api/system/v2/?api_key=" + settings.APIKey

	return url
}

// SendData - XXX
func SendData(data interface{}) error {
	url := SystemURL()

	JSONBytes, err := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(JSONBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Can't connect to the Amon API on %s\n", err.Error())
	}
	remoteLogger.Info("Sending data to %s\n", url)

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 209 {
		return fmt.Errorf("received bad status code, %d\n", resp.StatusCode)
	}

	return nil

}

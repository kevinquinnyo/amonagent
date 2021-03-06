package collectors

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	pshost "github.com/shirou/gopsutil/host"
)

// conversion units
const (
	MINUTE = 60
	HOUR   = MINUTE * 60
	DAY    = HOUR * 24
)

// Uptime - returns uptime string
// uptime = "{0} days {1} hours {2} minutes".format(days, hours, minutes)
func Uptime() string {
	boot, _ := pshost.BootTime()
	secondsFromBoot := uint64(time.Now().Unix()) - boot

	days := secondsFromBoot / DAY
	hours := (secondsFromBoot % DAY) / HOUR
	minutes := (secondsFromBoot % HOUR) / MINUTE

	s := fmt.Sprintf("%v days %v hours %v minutes", days, hours, minutes)

	return s
}

// IPAddress - returns machine IP
func IPAddress() string {
	c1, _ := exec.Command("hostname", "-I").Output()
	ipOutput := string(c1)
	ipList := strings.Split(ipOutput, " ")
	if len(ipList) > 0 {
		return ipList[0]
	}
	return ""

}

func (p DistroStruct) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}

// DistroStruct - returns information about the currently instaled distro
type DistroStruct struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

// Distro - gets distro info
// {'version': '14.04', 'name': 'Ubuntu'}
func Distro() DistroStruct {
	host, _ := pshost.HostInfo()

	d := DistroStruct{
		Version: host.PlatformVersion,
		Name:    host.Platform,
	}

	return d
}

// GetMetadataURL - XXX
func GetMetadataURL(provider string, url string) string {
	transport := &http.Transport{DisableKeepAlives: true}
	timeout := 2 * time.Second

	req, RequestErr := http.NewRequest("GET", url, nil)
	if provider == "google" {
		req.Header.Set("Metadata-Flavor", "Google")
	}
	if RequestErr != nil {
		return ""
	}

	client := &http.Client{Transport: transport}

	timer := time.AfterFunc(timeout, func() {
		transport.CancelRequest(req)
		fmt.Println(url, "Metadata URL time out")
	})
	defer timer.Stop()

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 209 {
		return ""
	}

	data, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return ""
	}

	id := string(data)

	return id
}

// CloudID - XXX
func CloudID() string {
	MetadataURLs := map[string]string{
		"google":       "http://169.254.169.254/computeMetadata/v1/instance/id",
		"amazon":       "http://169.254.169.254/latest/meta-data/instance-id",
		"digitalocean": "http://169.254.169.254/metadata/v1/id",
	}
	var CloudID string
	wg := sync.WaitGroup{}
	for provider, url := range MetadataURLs {
		wg.Add(1)

		go func(provider string, url string) {
			defer wg.Done()
			response := GetMetadataURL(provider, url)
			if len(response) > 0 {
				CloudID = response
			}

		}(provider, url)

	}

	wg.Wait()

	return CloudID
}

//MachineID - XXX
func MachineID() string {
	var machineidPath = "/etc/opt/amonagent/machine-id" // Default machine id path, generated on first install
	var MachineID string
	if _, err := os.Stat(machineidPath); os.IsNotExist(err) {
		machineidPath = "/etc/machine-id"
		// Does not exists, probably an older distro or docker container. TODO - REMOVE THIS IN FUTURE RELEASES
		if _, err := os.Stat(machineidPath); os.IsNotExist(err) {

			// Try one last path
			var machineidPath = "/var/lib/dbus/machine-id"
			if _, err := os.Stat(machineidPath); os.IsNotExist(err) {
				machineidPath = ""
			}

		}
	}

	if len(machineidPath) > 0 {
		file, err := os.Open(machineidPath)
		if err != nil {
			fmt.Printf(err.Error())
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if len(lines) > 0 {
			MachineID = lines[0]
		}
	}

	// Can't detect, return an empty string and ask for a server key
	if len(MachineID) != 32 {
		MachineID = ""
	}

	return MachineID
}

// Host - XXX
func Host() string {
	host, _ := os.Hostname()

	return host
}

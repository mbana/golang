package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/moul/http2curl"
	"github.com/sirupsen/logrus"
)

const (
	printCurlCommand = false
)

// RequestToCurlCommand ...
func RequestToCurlCommand(req *http.Request, methodName string) {
	if !printCurlCommand {
		return
	}

	command, err := http2curl.GetCurlCommand(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":     err,
			"command": command,
		}).Fatal(fmt.Sprintf("%s:GetCurlCommand", methodName))
	}

	logrus.WithFields(logrus.Fields{
		"command": command,
	}).Info(fmt.Sprintf("%s:GetCurlCommand", methodName))
}

// GetOutboundIP preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"ok":        ok,
			"localAddr": localAddr,
		}).Debug("GetOutboundIP")
	}
	ip := localAddr.IP

	logrus.WithFields(logrus.Fields{
		"ip": ip,
	}).Debug("GetOutboundIP")

	return ip
}

// ReadJSONFromFile ...
func ReadJSONFromFile(filename string, data interface{}) interface{} {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"filename": filename,
			"err":      err,
			"length":   len(bytes),
		}).Fatal("utils:ReadJSONFromFile:ReadFile")
	}

	logrus.WithFields(logrus.Fields{
		"filename": filename,
		"length":   len(bytes),
	}).Info("utils:ReadJSONFromFile:ReadFile")

	if err := json.Unmarshal(bytes, data); err != nil {
		logrus.WithFields(logrus.Fields{
			"filename": filename,
			"err":      err,
			"length":   len(bytes),
		}).Fatal("utils:ReadJSONFromFile:Unmarshal")
	}

	return data
}

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// colours
var (
	Green  = color.New(color.FgGreen).SprintFunc()  // nolint:gochecknoglobals
	Red    = color.New(color.FgRed).SprintFunc()    // nolint:gochecknoglobals
	Yellow = color.New(color.FgYellow).SprintFunc() // nolint:gochecknoglobals
	// Blue   = color.New(color.FgBlue).SprintFunc()   // nolint:gochecknoglobals
)

// PrintInBox -
func PrintInBox(seperatorColor func(a ...interface{}) string, vals ...string) {
	separator := strings.Repeat("-", 50)

	fmt.Printf("%+v\n", seperatorColor(separator))
	for _, val := range vals {
		fmt.Printf("%+v\n", val)
	}
	fmt.Printf("%+v\n", seperatorColor(separator))
}

// GetCustomerIP -
func GetCustomerIP() string {
	// conn, err := net.Dial("udp", "8.8.8.8:80")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()

	// localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	// if !ok {
	// 	logrus.WithFields(logrus.Fields{
	// 		"ok":        ok,
	// 		"localAddr": localAddr,
	// 	}).Debug("GetOutboundIP")
	// }
	// ip := localAddr.IP

	// logrus.WithFields(logrus.Fields{
	// 	"ip": ip,
	// }).Debug("GetOutboundIP")

	// return ip
	return "104.25.212.99"
}

// ReadJSONFromFile -
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

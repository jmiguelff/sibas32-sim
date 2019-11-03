package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/tarm/serial"
	"gopkg.in/yaml.v2"
)

// SerialOptT struct used to keep the YAML data
type SerialOptT struct {
	SerialConf struct {
		Name     string `yaml:"name"`
		Device   string `yaml:"device"`
		Baud     int    `yaml:"baud"`
		Stopbits int    `yaml:"stopbits"`
		Parity   string `yaml:"parity"`
	} `yaml:"serialConf"`
}

func main() {
	// Open yaml file
	fd, err := ioutil.ReadFile("settings.yml")
	if err != nil {
		log.Fatalln(err)
	}

	var serialOpts SerialOptT

	err = yaml.Unmarshal(fd, &serialOpts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Serial port name: %s\n", serialOpts.SerialConf.Name)
	fmt.Printf("\tDevice: %s\n", serialOpts.SerialConf.Device)
	fmt.Printf("\tBaudrate: %d\n", serialOpts.SerialConf.Baud)
	fmt.Printf("\tStopbits: %d\n", serialOpts.SerialConf.Stopbits)
	fmt.Printf("\tParity: %s\n", serialOpts.SerialConf.Parity)

	// Start serial port
	cSerial := new(serial.Config)
	cSerial.Name = serialOpts.SerialConf.Device
	cSerial.Baud = serialOpts.SerialConf.Baud
	cSerial.StopBits = serial.StopBits(serialOpts.SerialConf.Stopbits)
	cSerial.Parity = serial.Parity(serialOpts.SerialConf.Parity[0])
	cSerial.ReadTimeout = time.Millisecond * 500

	fmt.Println("Opening serial port")

	sfd, err := serial.OpenPort(cSerial)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Closing serial port")

	err = sfd.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Wait for command

}

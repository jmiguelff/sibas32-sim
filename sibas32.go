package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	//"github.com/tarm/serial"
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

	// Wait for command

}

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xeipuuv/gojsonschema"
)

type Device struct {
	Name      string `json:"Name"`
	Type      string `json:"Type"`
	Info      string `json:"Info"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

type Data struct {
	Devices []Device `json:"Devices"`
}

type Output struct {
	ValueTotal int      `json:"ValueTotal"`
	UUIDS      []string `json:"UUIDS"`
}

func IsBefore(d *Device) bool {
	deviceTimeInt64, err := strconv.ParseInt(d.Timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return deviceTimeInt64 < time.Now().Unix()
}

func GetUuid(d *Device) string {
	start := strings.Index(d.Info, ":")
	end := strings.Index(d.Info, ",")
	uuid := d.Info[start+1 : end]
	return uuid
}

func DecodeValue(d *Device) int {
	valueBytes, err := base64.StdEncoding.DecodeString(d.Value)
	if err != nil {
		panic(err)
	}
	value := string(valueBytes)
	currentValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return currentValue
}

func main() {
	// parse data from data.json
	dataBytes, err := os.ReadFile("./data/data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var data Data
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		fmt.Println(err)
		return
	}

	// discard the devices where the timestamp value is before the current time
	var validDevices []Device
	for _, device := range data.Devices {
		if !IsBefore(&device) {
			validDevices = append(validDevices, device)
		}
	}

	// get the total of all value
	var valueTotal int
	var uuids []string
	for _, device := range validDevices {
		// decode base64
		valueTotal += DecodeValue(&device)

		// get uuid
		uuids = append(uuids, GetUuid(&device))
	}

	// create output
	output := Output{ValueTotal: valueTotal, UUIDS: uuids}

	// write output to JSON file
	outputBytes, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := os.WriteFile("output.json", outputBytes, 0644); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Output written to output.json")

	// validate output file format with provided JSON schema
	schemaLoader := gojsonschema.NewReferenceLoader("file://./output-schema/schema.json")
	outputLoader := gojsonschema.NewReferenceLoader("file://./output.json")

	result, err := gojsonschema.Validate(schemaLoader, outputLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The output file is valid\n")
	} else {
		fmt.Printf("The output file is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

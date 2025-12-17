package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Name  string
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

// low performance mode
func main() {
	startTime := time.Now()
	delimiter := ";"
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, delimiter)
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]
		temp, _ := strconv.ParseFloat(rawTemp, 64)
		measurement, ok := data[location]

		if !ok {
			measurement = Measurement{
				Name: location,
				Min:  temp,
				Max:  temp,
				Sum:  temp,
				Count: 1,
			} 
		}else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}
		data[location] = measurement
	}

	locations := make([]string, 0, len(data))
	for location := range data {
		locations = append(locations, location)
	}
	sort.Strings(locations)

	for _, location := range locations {
		measurement := data[location]
		fmt.Printf("%s=%.1f/%.1f/%.1f, ",
		 location, 
		 measurement.Min, 
		 measurement.Sum/float64(measurement.Count),
		 measurement.Max,
		 )
	}
	fmt.Printf("}\n")
	fmt.Printf("Processing time: %.1f\n", time.Since(startTime).Seconds())
}
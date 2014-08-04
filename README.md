# Scratchgo
_The Remote Sensors Protocol can be used to connect golang to Scratch._

# Install Go lang
Refer to [golang.org](http://golang.org/doc/install)

# Enable Remote Sensor Connections
In your Scratch project, right click on one of the sensor tiles and click "enable remote sensor connections".

# Sample client program
``` go
package main

import (
	"fmt"
	"scratchgo"
)

func main() {
	conn, err := scratchgo.NewDefaultConnect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	// A global variable is changed
	err := conn.SensorUpdate("intX", "1")
	err = conn.SensorUpdate("intY", "1")

	// A broadcast is sent
	err = conn.BroadCast("update_pos")

	// receive the updated value
	// sensor-update - map[key:value]
	// broadcast - map[key:nil]
	msg, err := conn.Recv()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*msg) // {sensor-update map[scratch_value:1]}
}
```
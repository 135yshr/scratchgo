package main

import (
	"bufio"
	"fmt"
	_ "io/ioutil"
	"os"
	"scratchgo"
	"strings"
)

func main() {
	conn, err := scratchgo.NewDefaultConnect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		for {
			msg, err := conn.Recv()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(*msg)
		}
	}()

	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		exitOnErr("read buffer", err)

		line = strings.Replace(line, "\n", "", -1)
		buff := strings.Split(line, " ")

		switch strings.ToLower(buff[0]) {
		case "send":
			err = conn.SensorUpdate(buff[1], buff[2:])
			exitOnErr("update sensor", err)
		case "broadcast":
			err = conn.BroadCast(buff[1:])
			exitOnErr("broadcast message", err)
		case "exit", "quit":
			os.Exit(0)
		default:
			fmt.Println("unknown command.\n  send or broadcast")
		}
	}
	// for {
	// 	data := make([]byte, 255)
	// 	_, err := conn.Read(data)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(data)
	// }
	// data, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)
	os.Exit(0)
}

func exitOnErr(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

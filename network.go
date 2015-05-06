package main

import (
	"bufio"
	"fmt"
	"log" //DEBUG
	"net"
)

// TODO: external settings file
// TODO: const timeout time.Duration = time.ParseDuration("1000ms")

func retrieve(dest string, magicString string) (bodyLines []string, err error) {
	dest = dest + ":70"
	conn, err := net.Dial("tcp", dest)
	if err != nil {
		return
	}
	defer func() {
		e := conn.Close()
		if e != nil {
			log.Panicln(e) //DEBUG
		}
	}()
	fmt.Fprintf(conn, magicString)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		bodyLines = append(bodyLines, scanner.Text())
	}
	return
}

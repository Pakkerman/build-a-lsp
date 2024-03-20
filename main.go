package main

import (
	"bufio"
	"build-a-lsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/pakk/temp/build-lsp/log.txt")
	logger.Println("hello, there")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)

}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you need to give me a file, dude")
	}

	return log.New(logfile, "[build-a-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

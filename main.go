package main

import (
	"bufio"
	"build-a-lsp/lsp"
	"build-a-lsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/pakk/temp/build-lsp/log.txt")
	logger.Println("hello, there")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method any, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Cant do anything with this: %s", err)
		}
		logger.Printf("Connect to :%s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		// reply
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("sent the reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Cant do anything with this: %s", err)
		}
		logger.Printf("Opened: %s %s",
			request.Params.TextDocument.URI,
			request.Params.TextDocument.Text)
	}

}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you need to give me a file, dude")
	}

	return log.New(logfile, "[build-a-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

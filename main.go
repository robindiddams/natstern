package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/nats-io/nats.go"
)

type natsContext struct {
	URL   string `json:"url"`
	Creds string `json:"creds"`
}

func getContext() (string, []nats.Option, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil, err
	}
	contextFn := path.Join(homeDir, ".config/nats/context.txt")
	contextFile, err := os.ReadFile(contextFn)
	if err != nil {
		return "", nil, err
	}
	context := string(contextFile)

	contextPayloadFn := path.Join(homeDir, ".config", "nats", "context", context+".json")
	contextPayload, err := os.ReadFile(contextPayloadFn)
	if err != nil {
		return "", nil, err
	}
	natsCtx := natsContext{}
	if err := json.Unmarshal(contextPayload, &natsCtx); err != nil {
		return "", nil, err
	}
	if natsCtx.URL == "" {
		natsCtx.URL = nats.DefaultURL
	}
	var opts []nats.Option
	if natsCtx.Creds != "" {
		opts = append(opts, nats.UserCredentials(natsCtx.Creds))
		fmt.Println("using credential", natsCtx.Creds)
	}
	fmt.Println("using url", natsCtx.URL)
	return natsCtx.URL, opts, nil
}

func main() {
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	url, opts, err := getContext()
	if err != nil {
		panic(err)
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		panic(err)
	}
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}
	streamName := "changefeed"
	for data := range js.Consumers(streamName) {
		if command == "nz" && data.NumPending == 0 {
			continue
		}
		fmt.Printf("%s pending: %d, ackPending: %d\n", data.Name, data.NumPending, data.NumAckPending)
	}
}

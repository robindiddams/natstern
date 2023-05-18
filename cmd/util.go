package cmd

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

	fmt.Println("using context", context)
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
	}
	return natsCtx.URL, opts, nil
}

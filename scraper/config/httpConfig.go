package config

import (
	"time"

	"github.com/sethgrid/pester"
)

func SetHttpReqConfig(client *pester.Client) {
	client.Backoff = pester.ExponentialBackoff
	client.MaxRetries = 3
	client.Timeout = 30 * time.Second
}

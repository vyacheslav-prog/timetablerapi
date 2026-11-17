package services

import (
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewClient() (*kgo.Client, error) {
	client, err := kgo.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka client: %w", err)
	}
	return client, nil
}

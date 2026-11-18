package services

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewClient(ctx context.Context) (*kgo.Client, error) {
	client, err := kgo.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka client: %w", err)
	}
	err = client.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect kafka client: %w", err)
	}
	return client, nil
}

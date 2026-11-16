package services

import (
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewClient() (*kgo.Client, error) {
	return nil, fmt.Errorf("kafka connection is not implemented")
}

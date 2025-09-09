package main

import "fmt"

type registrarRepo struct {
}

func (rr *registrarRepo) SaveAndIdentifyPerformer(name string) (string, error) {
	return "", fmt.Errorf("no implemented")
}

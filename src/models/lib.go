package models

import (
	"strings"

	"github.com/YushchenkoAndrew/llama-cli/src/types"
)

func Active() []*types.Model {
	models := Load()
	active := []*types.Model{}

	for _, m := range models.Data {
		if m.Status.Value != "unloaded" {
			active = append(active, &m)
		}
	}

	return active
}

func Cached(model string) *types.Model {
	models := Load()

	for _, m := range models.Data {
		if m.Source == "cache" && strings.Contains(strings.ToUpper(m.ID), strings.ToUpper(model)) {
			return &m
		}
	}

	return nil
}

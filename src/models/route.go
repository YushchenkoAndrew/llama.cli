package models

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/YushchenkoAndrew/llama-cli/src/config"
	"github.com/YushchenkoAndrew/llama-cli/src/types"
)

var _models *types.Models

func Load() *types.Models {
	if _models != nil {
		return _models
	}

	res, err := http.Get(config.LLAMA_HOST + "/models")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	var models types.Models

	if err = json.NewDecoder(res.Body).Decode(&models); err != nil {
		panic(err)
	}

	_models = &models
	return _models
}

func Props(model string) types.Props {
	params := url.Values{}
	params.Add("model", model)

	res, err := http.Get(config.LLAMA_HOST + "/props?" + params.Encode())
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	var models types.Props

	if err = json.NewDecoder(res.Body).Decode(&models); err != nil {
		panic(err)
	}

	return models
}

func Stop(model string) {
	data := map[string]string{"model": model}
	jsonBytes, _ := json.Marshal(data)

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest(http.MethodPost, config.LLAMA_HOST+"/models/unload", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

func Delete(model string) {
	data := map[string]string{"model": model}
	jsonBytes, _ := json.Marshal(data)

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest(http.MethodDelete, config.LLAMA_HOST+"/models", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

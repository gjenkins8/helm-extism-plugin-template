package main

import (
	"fmt"

	pdk "github.com/extism/go-pdk"
)

type Input struct {
	Foo string `json:"foo"`
}

type Output struct {
	Bar string `json:"bar"`
}

func ReplaceMeImplementationGoesHere(input Input) (Output, error) {
	return Output{
		Bar: input.Foo,
	}, nil
}

func RunPlugin() error {
	var input Input
	if err := pdk.InputJSON(&input); err != nil {
		return fmt.Errorf("failed to parse input json: %w", err)
	}

	output, err := ReplaceMeImplementationGoesHere(input)
	if err != nil {
		pdk.Log(pdk.LogError, fmt.Sprintf("failed: %s", err.Error()))
		return err
	}

	if err := pdk.OutputJSON(output); err != nil {
		return fmt.Errorf("failed to write output json: %w", err)
	}

	return nil
}

//go:wasmexport helm_plugin
func HelmPlugin() uint64 {

	pdk.Log(pdk.LogDebug, "running plugin")

	if err := RunPlugin(); err != nil {
		pdk.Log(pdk.LogError, err.Error())
		pdk.SetError(err)
		return 1
	}

	return 0
}

func main() {}

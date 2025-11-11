package main

/**
 * Helm Wasm/Extism plugin template is a simple skeleton for creating a Helm Wasm plugin based on Extism's Go PDK.
 * The implementation handles Extism's PDK input/output and error handling.
 *
 * To implement your plugin logic, modify the `replaceMeImplementationGoesHere` function below.
 *
 *
 * For more details on using the Extism Go PDK, see: https://pkg.go.dev/github.com/extism/go-pdk
 */
import (
	"fmt"

	pdk "github.com/extism/go-pdk"
)

// Input defines the (JSON serialzable) plugin input
type Input struct {
	Foo string `json:"foo"`
}

// Input defines the (JSON serialzable) plugin output
type Output struct {
	Bar string `json:"bar"`
}

// ReplaceMeImplementationGoesHere is a placeholder for the actual plugin logic
func replaceMeImplementationGoesHere(input Input) (Output, error) {

	// ##############################
	// !! Implementation goes here !!
	// ##############################

	return Output{
		Bar: input.Foo,
	}, nil
}

// runPlugin wraps the plugin implementation with PDK input/output amnd error handling
func runPlugin() error {
	var input Input
	if err := pdk.InputJSON(&input); err != nil {
		return fmt.Errorf("failed to parse input json: %w", err)
	}

	output, err := replaceMeImplementationGoesHere(input)
	if err != nil {
		pdk.Log(pdk.LogError, fmt.Sprintf("failed: %s", err.Error()))
		return err
	}

	if err := pdk.OutputJSON(output); err != nil {
		return fmt.Errorf("failed to write output json: %w", err)
	}

	return nil
}

// Plugin entry point as invoked by Extism
//
//go:wasmexport helm_plugin_main
func HelmPluginMain() uint64 {

	pdk.Log(pdk.LogDebug, "running plugin")

	if err := runPlugin(); err != nil {
		pdk.Log(pdk.LogError, err.Error())
		pdk.SetError(err)
		return 1
	}

	return 0
}

func main() {} // Required for golang build (ignore)

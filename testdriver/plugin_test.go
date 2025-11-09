package plugin_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	extism "github.com/extism/go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/wazero"
)

type PluginInput struct {
	Foo string `json:"foo"`
}

type PluginOutput struct {
	Bar string `json:"bar"`
}

func init() {
	extism.SetLogLevel(extism.LogLevelDebug)
}

// TODO: replace this logic with Helm's Extism plugin runtime

func loadFilePlugin(ctx context.Context) (*extism.Plugin, error) {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: "../plugin.wasm",
				Name: "exmaple-plugin",
			},
			//extism.WasmData{
			//	Data: pluginBytes,
			//	Name: "gotemplate-renderer",
			//},
		},
		Memory: &extism.ManifestMemory{
			MaxPages: 65535,
			//MaxHttpResponseBytes: 1024 * 1024 * 10,
			//MaxVarBytes:          1024 * 1024 * 10,
		},
		Config: map[string]string{},
		//AllowedHosts: []string{"ghcr.io"},
		AllowedPaths: map[string]string{},
		Timeout:      0,
	}

	config := extism.PluginConfig{
		ModuleConfig:  wazero.NewModuleConfig().WithSysWalltime(),
		RuntimeConfig: wazero.NewRuntimeConfig().WithCloseOnContextDone(false),
		EnableWasi:    true,
		//EnableHttpResponseHeaders: true,
		//ObserveAdapter: ,
		//ObserveOptions: &observe.Options{},
	}

	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize plugin: %w", err)
	}

	plugin.SetLogger(func(logLevel extism.LogLevel, s string) {
		fmt.Printf("%s %s: %s\n", time.Now().Format(time.RFC3339), logLevel.String(), s)
	})

	return plugin, nil
}

func invokePugin(plugin *extism.Plugin, input *PluginInput) (*PluginOutput, error) {

	inputData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	exitCode, outputData, err := plugin.Call("helm_plugin", inputData)
	if err != nil {
		return nil, err
	}

	if exitCode != 0 {
		return nil, fmt.Errorf("plugin failed: exit code = %d", exitCode)
	}

	output := PluginOutput{}
	if err := json.Unmarshal(outputData, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

func TestExample(t *testing.T) {

	ctx := t.Context()

	plugin, err := loadFilePlugin(ctx)
	require.NoError(t, err)

	input := PluginInput{
		Foo: "example",
	}
	output, err := invokePugin(plugin, &input)
	require.NoError(t, err)

	assert.Equal(t, "example", output.Bar)
}

func BenchmarkExample(b *testing.B) {

	ctx := b.Context()

	plugin, err := loadFilePlugin(ctx)
	if err != nil {
		b.Fail()
	}

	input := PluginInput{
		Foo: "example",
	}

	for b.Loop() {
		_, err := invokePugin(plugin, &input)
		if err != nil {
			b.Fail()
		}
	}
}

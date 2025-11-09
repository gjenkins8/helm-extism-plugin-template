# helm-extism-plugin-template

Helm Wasm/Extism plugin template is a simple skeleton for creating a Helm Wasm plugin based on [Extism's Go PDK](https://pkg.go.dev/github.com/extism/go-pdk).

## Usage

```Makefile
make build      # Build the Wasm plugin
make test       # Run tests
make vet        # Static analysis
```

You can also pass custom flags to the test target:

```sh
make test TEST_FLAGS="-bench ."
```

## Testing

The `testdriver/` directory shows a simple driver for loading the plugin and executing tests against it.

# go_wasip1_wasm_exec

This project provides a Go-based runtime selector for WebAssembly (WASI) runtimes.

## Installation

To install the `go_wasip1_wasm_exec` package and its wrapper, use the following commands:

```sh
go install github.com/Zxilly/go_wasip1_wasm_exec/go_wasip1_wasm_exec@latest
go install github.com/Zxilly/go_wasip1_wasm_exec/go_wasip1_wasm32_exec@latest
```

## Usage

After installation, wrapper will automatically available in the `go` command.

### Environment Variables

- `GOWASIRUNTIME`: Specifies the runtime to use (default is `wasmtime`).
- `GOWASIRUNTIMEARGS`: Additional arguments to pass to the runtime.

### Example

```sh
export GOWASIRUNTIME=wasmer
```

## Supported Runtimes

- `wasmedge`
- `wasmer`
- `wazero`
- `wasmtime`

## License

This project is licensed under the MIT License.

# go-firewall

`go-firewall` is a Go wrapper library designed to provide programmatic control over the `iptables` firewall on Linux systems. This library aims to offer a simple and efficient way to interact with `iptables` by abstracting away the complexity of direct command-line interactions.

## Features

- **Simple API**: Intuitive Go functions to manage `iptables` rules.
- **Batch Operations**: Support for batch operations to apply multiple `iptables` changes efficiently.
- **Error Handling**: Provides detailed error information, making debugging easier.
- **Logging Support**: Integration with Go's native logging to track `iptables` modifications.
- **Extensive Coverage**: Covers most `iptables` features, including filtering, NAT, and mangle tables.

## Getting Started

### Prerequisites

- A Linux operating system with `iptables` installed.
- Go version 1.13 or later.

### Installation

To start using `go-firewall`, install Go and run `go get`:

```bash
go get -u github.com/bgrewell/go-firewall
```

This will retrieve the library and install it in your GOPATH.

## Usage

Here's a simple example of how to use go-firewall to append a rule to the output chain of the filter table:

```go
# Add example
```

## Contributing

We welcome contributions! Please submit a pull request or create an issue for any features or bug fixes.

## License

`go-firewall` is dual-licensed under the terms of the GNU General Public License v3.0 for open-source projects and a separate commercial license for commercial use.

### Open Source License

If you are creating an open-source project using `go-firewall`, you may use it under the terms of the GPL-3.0 License. Please see the [LICENSE.md](LICENSE.md) file for more details.

### Commercial Use

If you wish to use `go-firewall` in a commercial setting, a commercial license is required. Please see the [COMMERCIAL_LICENSE.md](COMMERCIAL_LICENSE.md) file for details on obtaining a commercial license.


## Disclaimer

This library is not affiliated with the official iptables project. It is a wrapper written by me to simplify its use 
in Go applications.
## Characters GO!

A sample backend application for storing characters and their item inventories for a multiplayer game.

## Getting Started

### Prerequisites
A working [Go environment](https://go.dev/doc/install). Goa tools are also required if making changes to the design, see [Goa prerequisites](https://goa.design/learn/getting-started/#prerequisites).

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/austinhuck/characters-go.git
   ```
2. Install the server and optionally the test client
   ```shell
   go install github.com/austinhuck/characters-go/cmd/gateway
   ```
   ```shell
   go install github.com/austinhuck/characters-go/cmd/gateway-cli
   ```

### Usage
Run the gateway server and make a test call with the test client
```shell
gateway
```
```shell
gateway-cli gateway get-characters
```

### Smoke Test
The `samples.sh` script can be run against a fresh server state and should finish cleanly. The output can be examined to verify correct operation.

## Development
### Project Structure
* Goa design - design
* Generated Goa code - gen
* Service implementations - service
* Gateway server - cmd/gateway

### Building
The following will build the gateway server and test client.
```shell
go build -o ./bin github.com/austinhuck/characters-go/cmd/gateway &&
go build -o ./bin github.com/austinhuck/characters-go/cmd/gateway-cli
```

### Running
Run in-dev gateway server and make a test call with the test client
```shell
./bin/gateway
```
```shell
./bin/gateway-cli gateway get-characters
```

### Tests
```shell
go test github.com/austinhuck/characters-go/service/character &&
go test github.com/austinhuck/characters-go/service/item &&
go test github.com/austinhuck/characters-go/service/inventory
```

### Design Changes
Modifications to the design require regenerating with Goa. To regenerate service interfaces, endpoints, transport code and OpenAPI spec:
```shell
goa gen github.com/austinhuck/characters-go/design
```

The implementations within the service module must then be adjusted accordingly. The gateway server and test tool implementation (cmd/gateway and cmd/gateway-cli) may also need to be adjusted depending on the change.

## License

Distributed under the MIT License. See `LICENSE` for more information.

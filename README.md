# Undeck

## Compiling and running

If you have `make` installed, you can simply run `$ make` to build the binary.

Execute `$ ./build/undeck serve` to start the server.

The application server listens on port `1337` and thus requires it to be free.

## Testing

#### Automated Testing
Execute `$ go test ./...` to run automated tests provided.

#### Manual Testing

The file [requests.http](requests.http) contains some sample http requests which can be run using the appropriate software, e.g. [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) for [Visual Studio Code](https://code.visualstudio.com/) or the built-in utility in [JetBrains](https://jetbrains.com) IDEs.

> __Note__: The ids in the urls need to be changed because they are auto-generated uuids.
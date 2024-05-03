# Rulebricks Go SDK

This SDK allows you to easily integrate Rulebricks' powerful rule engine into your Go applications.

## Installation

To use the Rulebricks Go SDK in your project, you can install it using `go get`:

```sh
go get github.com/rulebricks/go-sdk
```

## Configuration

Before you can start using the SDK, you need to configure it with your Rulebricks API key:

```go
client := rulebricks.NewClient("YOUR_API_KEY")
```

## Basic Usage

Here's a simple example of how to create a rule:

```go
rule, err := client.Rules.Create(rulebricks.Rule{
    Name:       "My Rule",
    Condition:  "if data.temperature > 100 { return true }",
    Action:     "fmt.Println('Alert! High temperature detected.')",
})
if err != nil {
    // Handle error
}
```

## Asynchronous Usage

The SDK supports asynchronous operations using Go routines:

```go
go func() {
    rule, err := client.Rules.Create(rulebricks.Rule{...})
    if err != nil {
        // Handle error
    }
    // Use the created rule
}()
```

## Error Handling

The SDK returns errors that can be handled in a typical Go error handling pattern:

```go
rule, err := client.Rules.Create(...)
if err != nil {
    // Handle the error
}
```

## Feedback and Contributions

We welcome feedback and contributions to the SDK. Please report any issues or suggestions through the [GitHub issue tracker](https://github.com/rulebricks/go-sdk/issues).

For contributions, please submit a pull request with a clear description of the changes and any relevant tests.

## License

The Rulebricks Go SDK is released under the MIT License. See the LICENSE file for more details.

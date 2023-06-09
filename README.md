# Go Service Patterns
This repo contains examples of patterns that can be useful in services.

## Package-Oriented Design
Each package should handle one concern, and have minimal surface area. For example, you may have a `handlers` package that is solely concerned with processing http/grpc requests, transforming that request into an internal request type, dispatching down to a `logic` package that implements pure application logic, and transforming the result into an http/grpc response.

## Pseudo-Enum
Enums define a type that can only be one of several named values. They can give you guarantees about the validity of data as it flows through your application. Go does not support true enums with compile-time guarantees, but we can emulate enum semantics with just a little caution in use.

To use these enums, you can refer to the symbolic names in your code, or parse them from strings at runtime. The one caveat is that you must not convert arbitrary numbers to the enum type.

See `internal/models/device/status` for an example.

```go
// good:
okStatus := status.Reachable

maybeStatus, err := status.Parse("MAINTENANCE")
if err != nil {
    // ...
}

// bad:
invalidStatus := status.Enum(99)
```

## Adapters to external services
Adapters help you integrate with external services by allowing you to define an
internal API that validates data coming from the wire, and transforms it into rich
domain models. This pattern helps you wall off your internal codebase from any changes
made to any client code you import, and can simplify mocks for tests.

See `internal/adapters/thingd` for an adapter to a faux-grpc client defined in `z_external/thingd`.

## TODO:
* handlers and logic layering
* mocking with gomock
* integration test suite
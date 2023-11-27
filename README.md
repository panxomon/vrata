# Vrata

Vrata is a Rust framework for building applications following the CQRS (Command Query Responsibility Segregation) and Event Sourcing patterns. It provides abstractions and utilities to help you implement these patterns efficiently.

## Features

- **CQRS:** Separate the responsibility of handling commands and queries.
- **Event Sourcing:** Store and retrieve the state of your application using a series of events.
- **HTTP Interface:** Easily expose your application through HTTP using the provided HTTP handlers.
- **gRPC Interface:** Implement gRPC services for seamless communication between microservices.

## Example: Legacy

Vrata includes a comprehensive example named "legacy," demonstrating the usage of a command (e.g., `create_user`), an event (e.g., `user_created`), and a query. This example illustrates a common scenario in CQRS and Event Sourcing applications.

```rust

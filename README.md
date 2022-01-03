# Go Patterns

A curated collection of idiomatic design & application patterns for Go language.

## Creational Patterns

Creational patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code.

### [Abstract Factory](./creational/abstractFactory/abstract_factory.go)

Provides an interface for creating families of related objects.

### [Factory Method](./creational/factoryMethod/factory_method.go)

Defers instantiation of an object to a specialized function for creating instances.

### [Builder](./creational/builder/builder.go)

Builds a complex object using simple objects.

### [Singleton](./creational/singleton/singleton.go)

Restricts instantiation of a type to one object.

### [Object Pool](./creational/objectPool/object_pool.go)

Instantiates and maintains a group of objects instances of the same type.

### [Prototype](./creational/prototype/prototype.go)

When the type of objects to create is determined by a prototypical instance, which is cloned to produce new objects.

## Structural Patterns

Structural Patterns are about organizing different classes and objects to from larger structures and provide new functionality.

### [Adapter](./structural/adapter/adapter.go)

Allows objects with incompatible interfaces to collaborate.

### [Bridge](./structural/bridge/bridge.go)

Decouples an interface from its implementation so that the two can vary independently.

### [Composite](./structural/composite/composite.go)

Encapsulates and provides access to a number of different objects.

### [Decorator](./structural/decorator/decorator.go)

Adds behavior to an object, statically or dynamically.

### [Facade](./structural/facade/facade.go)

Uses one type as an API to a number of others.

### [Flyweight](./structural/flyweight/flyweight.go)

Reuses existing instances of objects with similar/identical state to minimize resource usage.

### [Proxy](./structural/proxy/proxy.go)

Provides a surrogate for an object to control it's actions.

## Behavioral Patterns

Behavioral Patterns are about identifying common communication patterns between objects and realize these patterns.

### [Chain of Responsibility](./behavioral/chainOfResponsibility/chain_of_responsibility.go)

Avoids coupling a sender to receiver by giving more than object a chance to handle the request.

### [Command](./behavioral/command/command.go)

Bundles a command and arguments to call later.

### [Mediator](./behavioral/mediator/mediator.go)

Connects objects and acts as a proxy.

### [Memento](./behavioral/memento/memento.go)

Generate an opaque token that can be used to go back to a previous state.

### [Observer](./behavioral/observer/observer.go)

Provide a callback for notification of event/changes to data.

### [State](./behavioral/state/state.go)

Encapsulates varying behavior for the same object based on its internal state.

### [Strategy](./behavioral/strategy/strategy.go)

Enables an algorithm's behavior to be selected at runtime.

### [Template](./behavioral/template/template.go)

Defines a skeleton class which defers some methods to subclasses.

### [Visitor](./behavioral/visitor/visitor.go)

Separates an algorithm from an object on which it operates.

### [Iterator](./behavioral/iterator/iterator.go)

Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

## Synchronization Pattern

## Concurrency-Patterns

## Messageing-Patterns

## Stability-Patterns

## Profiling-Patterns

## Idioms

## Anti-Patterns

## End

Recommended Reading:

- [Software design pattern](https://en.wikipedia.org/wiki/Software_design_pattern)
- [Design patterns](https://refactoring.guru/design-patterns)
- [Golang by example](https://golangbyexample.com/)

## License

[MIT](./LICENSE)
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

Pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

### [Command](./behavioral/command/command.go)

Turns a request into a stand-alone object that contains all information about the request. This transformation pass requests as a method arguments, delay or queue a request's execution, and support undoable operations.

### [Iterator](./behavioral/iterator/iterator.go)

Traverse elements of a collection without exposing its underlying representation(list, stack, tree, etc).

### [Mediator](./behavioral/mediator/mediator.go)

Reduces coupling between components of a program by making them communicate indirectly, through a speacial mediator object.

### [Memento](./behavioral/memento/memento.go)

Save and restore the previous state of an object without revealing the details of its implementation.

### [Observer](./behavioral/observer/observer.go)

Define a subscription mechanism to notify multiple objects about any events that happen to the object they're observing.

### [State](./behavioral/state/state.go)

Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.

### [Strategy](./behavioral/strategy/strategy.go)

Define a family of algorithms, put each of time into a separate class, and make their objects interchangeable.

### [Template Method](./behavioral/templateMethod/template_method.go)

Defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.

#### [Visitor](./behavioral/visitor/visitor.go)

Separate algorithms from the objects on which they operate.

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
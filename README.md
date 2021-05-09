# Go Patterns

list of Go design patterns

### Creational Patterns

Creational patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code.

#### [Factory Method](./creational/factoryMethod/factory_method.go)

Provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. 

#### [Abstract Factory](./creational/abstractFactory/abstract_factory.go)

Provides an interface for creating families of related or dependent objects without specifiying their concrete classes.

#### [Builder](./creational/builder/builder.go)

The pattern produce different types and representations of an object using the same construction code.

#### [Prototype](./creational/prototype/prototype.go)

Copy existing objects without making code dependent on classes.

#### [Singleton](./creational/singleton/singleton.go)

Ensure a class has only one instance, and provide a global point of access to it.

### Structural Patterns

Structural Patterns are about organizing different classes and objects to from larger structures and provide new functionality.

#### [Adapter](./structural/adapter/adapter.go)

 Allows objects with incompatible interfaces to collaborate.

#### [Bridge](./structural/bridge/bridge.go)

Split a large class or a set of closely related classes into two separate hierarchies-abstraction and implementation, which can be developed independently of each other.

### Behavioral Patterns

Behavioral Patterns are about identifying common communication patterns between objects and realize these patterns.

## End

Recommended Reading:

- [Software design pattern](https://en.wikipedia.org/wiki/Software_design_pattern)
- [Design patterns](https://refactoring.guru/design-patterns)

## MIT

[Mit License](./LICENSE)
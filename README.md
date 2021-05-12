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

#### [Composite](./structural/composite/composite.go)

Compose objects into tree structures and then work with these structures as if they were individual objects.

#### [Decorator](./structural/decorator/decorator.go)

Attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.

#### [Facade](./structural/facade/facade.go)

Provides a simplified interface to a library, a framework, or any other complex set of classes.

#### [Flyweight](./structural/flyweight/flyweight.go)

Fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.

#### [Proxy](./structural/proxy/proxy.go)

Provide a surrogate/substitute or placeholder for another object to control access to it.

### Behavioral Patterns

Behavioral Patterns are about identifying common communication patterns between objects and realize these patterns.

#### [Chain of Responsibility](./behavioural/chainOfResponsibility/chain_of_responsibility.go)

Pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

#### [Command](./behavioural/command/command.go)

Turns a request into a stand-alone object that contains all information about the request. This transformation pass requests as a method arguments, delay or queue a request's execution, and support undoable operations.

#### [Iterator](./behavioural/iterator/iterator.go)

Traverse elements of a collection without exposing its underlying representation(list, stack, tree, etc).



## End

Recommended Reading:

- [Software design pattern](https://en.wikipedia.org/wiki/Software_design_pattern)
- [Design patterns](https://refactoring.guru/design-patterns)
- [Golang by example](https://golangbyexample.com/)

## MIT

[Mit License](./LICENSE)
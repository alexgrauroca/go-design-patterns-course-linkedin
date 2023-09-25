# go-design-patterns-course-linkedin

Go Design Patterns course, by LinkedIn Learning.

- [go-design-patterns-course-linkedin](#go-design-patterns-course-linkedin)
  - [Deisgn Pattern Categories](#deisgn-pattern-categories)
  - [Creational Pattern](#creational-pattern)
    - [Builder](#builder)
    - [Factory](#factory)
    - [Singleton](#singleton)
  - [Structural Pattern](#structural-pattern)
    - [Adapter](#adapter)
    - [Facade](#facade)

## Deisgn Pattern Categories

- Creational. How objects are created during the life time of a program.
  - Abstract Factory.
  - **Builder**.
  - **Factory**.
  - Lazy Initialization.
  - Multiton.
  - Object Pool.
  - Prototype.
  - **Singleton**.
- Structural. Organise various parts of the program into larger more complex structures, and to provide new functionality based on that organization.
  - **Adapter**.
  - Bridge.
  - Composite.
  - Decorator.
  - **Facade**.
  - Flyweight.
  - Proxy.
- Behavioral. Identify common patterns of communication and interaction between objects within a program, and are intended to increase flexibility an efficiency in carrying out these communications and interactions.
  - Chain of Responsability.
  - Command.
  - Interpreter.
  - **Iterator**.
  - Mediator.
  - Memento.
  - **Observer**.
  - Strategy.
  - Visitor.

## Creational Pattern

### Builder

Simplifies the creation of complex objects that have many possible representations.

- **Director**.
- **Builder**.
- **Complex Object**.

### Factory

Makes it easier to create objects without having to know the exact class or type of the objects that will be created.

- **Creator Interface**.
- **Concrete Creator**.
- **Product**.
- **Concrete Product**.

### Singleton

Ensures that only one instance of a given class or type definition can be instatiated at any one time.

- **Singleton**.

## Structural Pattern

### Adapter

Allows the interface of an existing subsystem or API to be used as another interface without modifying the code of the existing API.

- **Client**.
- **ExpectedInterface**.
- **OriginalClass**.
- **Adapter**.

### Facade

Provide a simple, front-facing interface to a more complex system, library, or API.

- **Subsystem**:
  - **Client A**.
  - **Client B**.
  - **Client C**.
- **Client 1**.
- **Client 2**.
- **Facade**:
  - **simpleAPI1**.
  - **simpleAPI2**.

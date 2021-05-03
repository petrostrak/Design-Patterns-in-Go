### Design Patterns in Go

This repo provides a comprehensive overview of Design Patterns in Go from a practical perspective. In particular, it covers patterns with the use of:

- The latest versions of the Go programming language
- Use of modern programming libraries and frameworks
- Discussions of pattern variations and alternative approaches

This repo provides an overview of all the Gang of Four (GoF) design patterns as outlined in their seminal book, together with modern-day variations, adjustments, discussions of intrinsic use of patterns in the language.

#### What are Design Patterns?

Design Patterns are reusable solutions to common programming problems. They were popularized with the 1994 book Design Patterns: Elements of Reusable Object-Oriented Software by Erich Gamma, John Vlissides, Ralph Johnson and Richard Helm (who are commonly known as a Gang of Four, hence the GoF acronym).

The original book GoF book used C++ and Smalltalk for its examples, but, since then, design patterns have been adapted to every programming language imaginable: C#, Java, Swift, Python, JavaScript and now — Go!

The appeal of design patterns is immortal: we see them in libraries, some of them are intrinsic in programming languages, and you probably use them on a daily basis even if you don't realize they are there.

#### What Patterns Does This Repo Cover?

It covers all the GoF design patterns. In fact, here's the full list of what is covered:

- SOLID Design Principles: Single Responsibility Principle, Open-Closed Principle, Liskov Substitution Principle, Interface Segregation Principle and Dependency Inversion Principle
- Creational Design Patterns: Builder, Factories (Factory Method and Abstract Factory), Prototype and Singleton
- Structrural Design Patterns: Adapter, Bridge, Composite, Decorator, Façade, Flyweight and Proxy
- Behavioral Design Patterns: Chain of Responsibility, Command, Interpreter, Iterator, Mediator, Memento, Observer, State, Strategy, Template Method and Visitor

#### SOLID is a mnemonic acronym for five design principles intended to make software designs more understandable, flexible, and maintainable. The SOLID concepts are:

- The Single-responsibility principle: "There should never be more than one reason for a class to change." In other words,   every class should have only one responsibility.
- The Open–closed principle: "Software entities ... should be open for extension, but closed for modification."
- The Liskov substitution principle: "Functions that use pointers or references to base classes must be able to use objects of   derived classes without knowing it".
- The Interface segregation principle: "Many client-specific interfaces are better than one general-purpose interface."
- The Dependency inversion principle: "Depend upon abstractions, [not] concretions."

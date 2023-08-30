# Day 41 of #66DaysOfGo

_Last update:  Aug 29, 2023_.

---

Today, I've started a new series related to the Clean Architecture.

---

## What is the Clean Architecture?

<img src="https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg" alt="The Clean Architecture" width="450"/>

The Clean Architecture is a software architectural pattern introduced by [Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), also known as Uncle Bob. It aims to create a separation of concerns in software systems, making them more maintainable, flexible, and independent of external factors. The Clean Architecture promotes the idea of designing software in layers, with each layer having a specific responsibility and level of abstraction.

It's related to the Hexagonal Architecture (introduced by [Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/), also known as the Ports and Adapters Architecture). It's another architectural pattern that aims to create highly decoupled and maintainable software systems. This architecture focuses on isolating the core business logic from external concerns, such as databases, frameworks, and user interfaces. Like the Clean Architecture, the Hexagonal Architecture also emphasizes the separation of concerns, but it provides a slightly different approach in terms of how these concerns are organized.

When implementing the Clean Architecture, the following components and principles are typically involved:

1. **Entities**: Entities are the core business objects or concepts that encapsulate the most general and high-level rules of the application. In Go, you can create entity structs that represent these business concepts. These entities should not be influenced by external concerns, such as databases or frameworks.

2. **Use Cases**: Use cases represent the application's specific business rules and use cases. These are the application's core functionalities that manipulate the entities and make decisions based on business logic. Use cases are typically implemented as methods or functions within the application's domain layer.

3. **Interfaces (Ports)**: The use cases define interfaces or ports that represent the interactions between the application's core and the external world. In Go, you can define interfaces to represent these ports. This allows you to keep the core business logic independent of external frameworks and technologies.

4. **Interactors (Use Case Implementations)**: Interactors are implementations of the use case interfaces defined in the previous step. These implementations orchestrate the flow of data and actions between the entities, and they apply the business rules defined by the use cases. Interactors reside in the domain layer of the application.

5. **Framework and Driver Adapters**: These are the outer layers of the application and are responsible for interacting with external components such as databases, web frameworks, and UIs. Adapters convert data from the format expected by the external component to the format used by the core application. In Go, these adapters might be implemented as HTTP handlers, database connectors, etc.

6. **Dependency Rule**: The key principle of the Clean Architecture is the dependency rule, which states that dependencies should always point inward toward the core of the application. In other words, inner layers (such as entities, use cases, and interactors) should not depend on outer layers (such as frameworks and databases). This helps maintain modularity and flexibility in the codebase.

In the context of the Go programming language, here are some considerations:

- **Structs**: Use structs to represent entities, which are the core business concepts.
- **Interfaces**: Define interfaces for the use cases (ports) and any external dependencies to ensure loose coupling.
- **Package Structure**: Organize your code into packages that reflect the Clean Architecture's layers, such as "domain", "usecase", "adapter" and "framework."
- **Dependency Injection**: Use dependency injection to provide implementations of use cases and external dependencies to the core of the application.
- **Testing**: The separation of concerns in Clean Architecture makes unit testing easier. You can test the core business logic (entities, use cases, interactors) without needing to worry about external components.

---

## References

- [(2012) The Clean Architecture by Robert C. Martin (uncle Bob)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [(2005) Hexagonal Architecture by Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/)
- [Clean Architecture : Part 2 – The Clean Architecture
](https://crosp.net/blog/software-architecture/clean-architecture-part-2-the-clean-architecture/)
- [https://manakuro.medium.com/clean-architecture-with-go-bce409427d31](https://manakuro.medium.com/clean-architecture-with-go-bce409427d31)
- [https://medium.com/full-stack-tips/dependency-injection-in-go-99b09e2cc480](https://medium.com/full-stack-tips/dependency-injection-in-go-99b09e2cc480)

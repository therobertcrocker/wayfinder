![Wayfinder - Logo](<assets/banner.png>)

---

**Wayfinder** is a suite of tools designed for tabletop role-playing game (TTRPG) enthusiasts, with the main goal of assisting the Game Master (GM) in generating and managing game content.


>The project is currently in the early stages of development, with the primary focus being on the **Machinations** module. The goal of this module is to provide a tool for managing the various factions and organizations in the world, utilizing the "Faction Turn" mechanics from Stars Without Number.


# Project Structure

The project is divided into several modules, each of which is designed to be as independent as possible. The planned modules are as follows:

### MVP
- **Core**: The main service handler that provides access to the various modules through an event bus.
- **Codex**: Manages the state of the world and various world entities, such as regions, locations, NPCs, timelines, and historical events. Contains references to the objects managed by the other modules, such as factions and assets. The Codex is the primary source of truth for the world state.
- **Machinations**: A tool for managing the various factions and organizations in the world. Will eventually handle the "Faction Turn" mechanics from Stars Without Number. 


## Design Philosophy
- **Modular**: Each module is designed to be as independent as possible, allowing for easy integration into existing projects.
- **Maintainable**: The code is written with a focus on readability and maintainability, with a consistent style and clear documentation.
- **Extensible**: The code is designed to be easily extended and modified, allowing for the addition of new features and functionalities.

## Development Principles
- **Event Driven**: The various modules communicate with each other through an event bus, allowing for a high degree of flexibility and customization.
- **Test Driven Development**: Each module is developed using a test-driven approach, with a focus on code coverage and documentation.


### TTD  Workflow
0. **New Feature**: The first step is to identify a new feature or functionality to be added to the project.
1. **Write a test**: The first step is to write a test for the desired functionality.
2. **Run the test**: The test should fail, since the functionality has not been implemented yet.
3. **Write the code**: The next step is to write the code to make the test pass.
4. **Refactor**: Once the test passes, the code should be refactored to ensure readability and maintainability.
5. **Repeat**: The process is repeated for each new feature or functionality.


## Technologies Used
- **Go**: The primary language for the project.
- **bbolt**: A fast, embedded, pure Go key/value database.
- **Cobra**: A Go library for creating powerful modern CLI applications.




# Development

## Minimum Viable Product (MVP)

The primary focus of this initial phase is on the **Machinations** module, providing functionalities such as:

- CRUD operations for factions and their assets.
- Storage of factions and assets in a bbolt database.
- CLI for interacting with the module.
- Event bus for communicating with other modules.

---

## Installation & Usage

*Details will be added once the project reaches a certain level of functionality.*

---

## License

This project is licensed under the [MIT License](LICENSE).

---

Stay tuned for more updates as Wayfinder continues to grow and evolve!

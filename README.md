![Wayfinder - Logo](<assets/banner.png>)

---

**Wayfinder** is a suite of tools designed for tabletop role-playing game (TTRPG) enthusiasts, with the main goal of assisting the Game Master (GM) in generating and managing game content.

```
The project is currently in the early stages of development, with the primary focus being on the **Factions** module. The goal of this module is to provide a tool for managing the various factions and organizations in the world, utilizing the "Faction Turn" mechanics from Stars Without Number.
```

# Project Structure

The project is divided into several modules, each of which is designed to be as independent as possible. The planned modules are as follows:

### MVP
- **Core**: An event bus and messaging system with shared utilities.
- **Codex**: Manages the state of the world and various world entities, such as regions, locations, NPCs, timelines, and historical events.
- **Machinations**: A tool for managing the various factions and organizations in the world. Will eventually handle the "Faction Turn" mechanics from Stars Without Number. 

### Future
- **Chronicle**: A tool for creating and managing in-world quests.
- **Generation**: Various random content generation tools (names, loot, etc.).


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

The primary focus of this initial phase is on the **Factions** module, providing functionalities such as:

- CRUD operations for factions and their assets.
- Association of factions with locations.
- Storage of factions and assets in a bbolt database.


### MVP Tasks

**1. Database Setup with bbolt:**
- Set up the foundational database layer with bbolt.
- Establish schemas and structures to ensure data consistency and easy retrieval.
- Implement basic CRUD operations for persistence and retrieval of data.

**2. Codex Initialization:**

- Design the subsystem for Locations within the Codex.
- Once the design is finalized, implement CRUD operations for Location entities.
  
**3. Faction Management MVP:**
- Design and implement the ECS framework for Factions and Assets.
- Implement CRUD operations for both entities.
- establish relations between Factions, Assets, and the Codex (Locations, specifically).

**4. CLI Setup using Cobra:**
- Setup commands for CRUD operations of Factions and Assets.
- Implement other desired commands.



## Future Plans

Once the MVP for the Factions module is complete, development will expand to include the following modules:

- **Core**: An event bus and messaging system with shared utilities.
- **Codex**: Manages the state of the world and various world entities, such as regions, locations, NPCs, timelines, and historical events.
- **Quests**: A tool for creating and managing in-world quests.
- **Generation**: Various random content generation tools (names, loot, etc.).

---

## Installation & Usage

*Details will be added once the project reaches a certain level of functionality.*

---

## License

This project is licensed under the [MIT License](LICENSE).

---

Stay tuned for more updates as Wayfinder continues to grow and evolve!

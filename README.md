# Goblin Tower

Goblin Tower is a fun text-based game where you play as a hero trying to conquer a tower infested with goblins. Your goal is to progress through the tower, level up, defeat goblins, and stay alive as long as possible. Are you up for the challenge?

## Features

- Randomly generated hero with unique stats.
- Randomly generated goblins with varying difficulties.
- Turn-based combat with goblins.
- Potion shop to buy health potions.
- Level up after every 10 steps.
- Interactive gameplay with user choices.
- Game over when the hero's health reaches zero.

## Details
### Mechanics
Character Stats: Each new hero is generated with random max health points, attack power, defense power, and a set of potions. The goblins also have randomized stats.
Turn-Based Combat: When encountering a goblin, the hero and the goblin engage in turn-based combat until one of them's health reaches zero.
Leveling Up: The hero advances one level after every 10 steps.
Potion Shop: You can visit the potion shop to purchase health potions with your gold.
Game Over: The game ends when the hero's health reaches zero.
### MVC Architecture
#### Model:
The model package contains data structures for the hero and goblins, as well as game logic.
#### View:
The view package handles the user interface and displaying game information.
#### Controller:
The controller package controls the game's flow, user interactions, and decision-making.
#### Patterns Used
MVC (Model-View-Controller): The game follows the MVC architectural pattern, separating data (Model), presentation (View), and control (Controller) aspects.
### Data Structures used
Hero: The hero character is represented as a struct with attributes for health, attack, defense, gold, potions, and more.
Goblin: Goblins are also represented as structs with randomized attributes.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/)

### Installation

```sh
# Clone the repository
git clone https://github.com/yourusername/goblin-tower.git

# Navigate to the project directory
cd goblin-tower

# Run the game
go run main.go
How to Play
Use 's' to take a step.
Use 'p' to drink a potion.
Use 'q' to quit the game.
Customization
You can customize the game by modifying the settings in the config.go file. Adjust the initial hero stats, goblin stats, and other game parameters.
```

License
This project is licensed under the MIT License - see the LICENSE file for details.
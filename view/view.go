package view

import (
    "fmt"
	"goblin_tower/model"
)

func DisplayGameInfo(hero *model.Hero) {
    fmt.Printf("\nYou are at Level %d with %d HP, %d Attack, %d Defense, and %d Gold.\n", hero.Level, hero.HP, hero.Attack, hero.Defense, hero.Gold)
    fmt.Printf("You have %d potions and %d steps taken.\n", len(hero.Potions), hero.Level*10)
}

func AskForAction() string {
    var action string
    for action != "s" && action != "p" && action != "q" {
        fmt.Print("Choose an action (s for step, p for potion, q for quit): ")
        fmt.Scanln(&action)
    }
    return action
}

func AskForPotionPurchase() bool {
    var input string
    fmt.Print("Do you want to purchase a potion? (y/n): ")
    fmt.Scanln(&input)
    return input == "y" || input == "Y"
}

func DisplayPotionBought() {
    fmt.Println("You bought a potion.")
}

func DisplayPotionShop(hero *model.Hero) {
    fmt.Printf("You reached a Potion Shop!\nYou have %d Gold. Each potion costs 4 Gold.\n", hero.Gold)
}

func DisplayPotionLimitReached() {
    fmt.Println("You can't carry more than 5 potions!")
}

func DisplayLeavePotionShop() {
    fmt.Println("You leave the Potion Shop.")
}

func DisplayPotionDrank(healthRestored, newHP int) {
    fmt.Printf("You drink a potion and restore %d HP. You now have %d HP.\n", healthRestored, newHP)
}

func DisplayNoPotionsLeft() {
    fmt.Println("You have no potions left.")
}

func DisplayEncounterGoblin(goblin *model.Goblin) {
    fmt.Printf("You encounter a Goblin with %d HP, %d Attack, and %d Defense!\n", goblin.HP, goblin.Attack, goblin.Defense)
}

func DisplayGoblinDefeated() {
    fmt.Println("You defeated the Goblin and gained 2 Gold!")
}

func DisplayDefeatedByGoblin() {
    fmt.Println("You were defeated by the Goblin.")
}

func DisplayGameSummary(hero *model.Hero) {
    fmt.Printf("You reached Level %d and defeated %d goblins. Thanks for playing!\n", hero.Level, hero.Gold/2)
}

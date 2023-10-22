package controller

import (
    "goblin_tower/model"
    "goblin_tower/view"
    "math/rand"
    "github.com/eiannone/keyboard"
)

func RunGame() {
    hero := model.NewHero()

    err := keyboard.Open()
    if err != nil {
        panic(err)
    }
    defer keyboard.Close()

    for {
        view.DisplayGameInfo(hero)
        char, key, err := keyboard.GetKey()
        if err != nil {
            panic(err)
        }

        if key == keyboard.KeyEsc {
            return
        }

        if char == 's' {
            hero.Level++
            hero.HP--
            if hero.HP < 0 {
                hero.HP = 0
            }
            if hero.HP == 0 {
                view.DisplayDefeatedByGoblin()
                view.DisplayGameSummary(hero)
                return
            }
            if hero.Level%10 == 0 {
                visitPotionShop(hero)
            } else {
                encounterGoblin(hero)
            }
        } else if char == 'p' {
            drinkPotion(hero)
        }
    }
}
func encounterGoblin(hero *model.Hero) {
    if rand.Intn(2) == 0 {
        goblin := model.NewGoblin()
        view.DisplayEncounterGoblin(goblin)
        fight(hero, goblin)
    }
}

func fight(hero *model.Hero, goblin *model.Goblin) {
    for hero.HP > 0 && goblin.HP > 0 {
        damageToGoblin := hero.Attack - goblin.Defense
        if damageToGoblin < 1 {
            damageToGoblin = 1
        }
        goblin.HP -= damageToGoblin

        damageToHero := goblin.Attack - hero.Defense
        if damageToHero < 1 {
            damageToHero = 1
        }
        hero.HP -= damageToHero

        if hero.HP < 0 {
            hero.HP = 0
        }
    }

    if hero.HP > 0 {
        hero.Gold += 2
        view.DisplayGoblinDefeated()
    } else {
        view.DisplayDefeatedByGoblin()
    }
}

func visitPotionShop(hero *model.Hero) {
    view.DisplayPotionShop(hero)
    for hero.Gold >= 4 {
        choice := view.AskForPotionPurchase()
        if choice {
            if len(hero.Potions) < 5 {
                hero.Potions = append(hero.Potions, 2)
                hero.Gold -= 4
                view.DisplayPotionBought()
            } else {
                view.DisplayPotionLimitReached()
                break
            }
        } else {
            break
        }
    }
    view.DisplayLeavePotionShop()
}

func drinkPotion(hero *model.Hero) {
    if len(hero.Potions) == 0 {
        view.DisplayNoPotionsLeft()
        return
    }

    potion := hero.Potions[len(hero.Potions)-1]
    hero.Potions = hero.Potions[:len(hero.Potions)-1]

    newHP := hero.HP + potion
    if newHP > hero.MaxHP {
        newHP = hero.MaxHP
    }
    healthRestored := newHP - hero.HP
    hero.HP = newHP

    view.DisplayPotionDrank(healthRestored, newHP)
}

func levelUp(hero *model.Hero) {
    if hero.HP == 0 {
        view.DisplayDefeatedByGoblin()
        view.DisplayGameSummary(hero)
        return
    }

    if hero.Level%10 == 0 {
        visitPotionShop(hero)
    }
}

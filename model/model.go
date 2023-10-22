package model

import (
    "math/rand"
    "time"
)

type Hero struct {
    MaxHP   int
    HP      int
    Attack  int
    Defense int
    Gold    int
    Potions []int
    Level   int
}

type Goblin struct {
    HP     int
    Attack int
    Defense int
}

func NewHero() *Hero {
    rand.Seed(time.Now().UnixNano())
    return &Hero{
        MaxHP:   rand.Intn(11) + 20,
        HP:      rand.Intn(11) + 20,
        Attack:  rand.Intn(3) + 1,
        Defense: rand.Intn(5) + 1,
        Gold:    0,
        Potions: make([]int, 5),
        Level:   1,
    }
}

func (h *Hero) RestoreHP() {
    h.HP = h.MaxHP
}

func NewGoblin() *Goblin {
    rand.Seed(time.Now().UnixNano())
    return &Goblin{
        HP:     rand.Intn(6) + 5,
        Attack: rand.Intn(2) + 2,
        Defense: rand.Intn(2) + 1,
    }
}

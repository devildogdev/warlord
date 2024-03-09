package store

import (
    "fmt"
)

const maxInventory int = 100

type Weapon struct {
    Model string
    Price int
}

type Store struct {
    Region string
    Event string
    Inventory map[string][]Weapon
}

var models = map[string]int{
    "G19": 600,
    "1911": 1000,
    "M4": 1400,
    "AK-47": 2200,
    "RPG": 15000,
    "Mk19": 41000,
    "M24": 18000,
    "M107": 27000,
    "M2": 62000,
    "GAU-17": 250000,
}

func New(region string) Store {
    s := Store{
        Region: region,
        Inventory : make(map[string][]Weapon),
    }
    for m, p := range models {
        s.stockUp(m, p)
    }
    return s
}

func (s *Store) stockUp(model string, price int) {
    stock := make([]Weapon, maxInventory)
    for i := range stock {
        stock[i] = Weapon{Model: model, Price: price}
    }
    s.Inventory[model] = stock
}

func (s Store) ShowInventory() {
    for m, wl := range s.Inventory {
        fmt.Printf("%s: %d\n", m, len(wl))
    }
}


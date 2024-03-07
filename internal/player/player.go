package player

import (
    "errors"

    "github.com/j-tew/warlord/internal/store"
)

type Player struct {
    name, Location string
    health int8
    cash, bank int
    inventory map[string][]store.Weapon
}

func New(name string) Player {
    return Player{
        name: name,
        health: 100,
        cash: 15000,
        bank: 0,
        inventory: make(map[string][]store.Weapon),
        Location: "North America", 
    }
}

func (p *Player) BuyWeapon(s store.Store, model string, qty int) error {
    if qty <= 0 {
        return errors.New("Quantity must be greater than Zero!")
    }
    cart := s.Inventory[model][:qty]
    var cost int
    for _, w := range cart {
        cost += w.Price
    }
    if p.cash >= cost {
        p.cash -= cost
        s.Inventory[model] = s.Inventory[model][qty:]
        p.inventory[model] = append(p.inventory[model], cart...)
    }
    return nil
}

func (p *Player) SellWeapon(s store.Store, model string, qty int) error {
    if len(p.inventory) < 1 {
        return errors.New("You don't have any weapons to sell")
    } else if qty > len(p.inventory) {
        return errors.New("You cannot sell more than you have")
    }
    var profit int
    sold := p.inventory[model][:qty]
    for _, w := range sold {
        profit += w.Price
    }
    p.inventory[model] = p.inventory[model][qty:]
    s.Inventory[model] = append(s.Inventory[model], sold...)
    return nil
}

func (p *Player) Damage(value int8) {
    if value != p.health {
        p.health -= value
    } else {
        p.health = 0
    }
}


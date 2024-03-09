package player

import (
    "errors"
    "fmt"

    "github.com/j-tew/warlord/internal/store"
)

var Stores =  map[string]store.Store{
    "North America": store.New("North America"),
    "South America": store.New("South America"),
    "Europe": store.New("Europe"),
    "North Africa": store.New("North Africa"),
    "South East Asia": store.New("South East Asia"),
    "Middle East": store.New("Middle East"),
}


type Player struct {
    Name, Region string
    Health int8
    Cash, Bank int
    Inventory map[string][]store.Weapon
}

func New(name string) Player {
    return Player{
        Name: name,
        Health: 100,
        Cash: 15000,
        Bank: 0,
        Inventory: make(map[string][]store.Weapon),
        Region: "North America", 
    }
}

func (p Player) ShowInventory() {
    for m, wl := range p.Inventory {
        fmt.Printf("%s: %d\n", m, len(wl))
    }
}

func (p *Player) Move(region string) error {
    _, exists := Stores[region]
    if exists {
        p.Region = region
    } else {
        return errors.New("Invalid region")
    }
    return nil
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
    if p.Cash >= cost {
        p.Cash -= cost
        s.Inventory[model] = s.Inventory[model][qty:]
        p.Inventory[model] = append(p.Inventory[model], cart...)
    }
    return nil
}

func (p *Player) SellWeapon(s store.Store, model string, qty int) error {
    if len(p.Inventory) < 1 {
        return errors.New("You don't have any weapons to sell")
    } else if qty > len(p.Inventory) {
        return errors.New("You cannot sell more than you have")
    }
    var profit int
    sold := p.Inventory[model][:qty]
    for _, w := range sold {
        profit += w.Price
    }
    p.Inventory[model] = p.Inventory[model][qty:]
    s.Inventory[model] = append(s.Inventory[model], sold...)
    return nil
}

func (p *Player) Damage(value int8) {
    if value != p.Health {
        p.Health -= value
    } else {
        p.Health = 0
    }
}


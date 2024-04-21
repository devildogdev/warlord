package player

import (
    "errors"
    "strconv"
    "slices"

    "github.com/j-tew/warlord/internal/store"
    "github.com/j-tew/warlord/internal/weapon"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/lipgloss/table"
)

type Player struct {
    Name, Region string
    Health int8
    Cash, Bank int
    Inventory map[string][]weapon.Weapon
    InventoryTable *table.Table
}

func New(name string) *Player {
    p := &Player{
        Name: name,
        Health: 100,
        Cash: 15000,
        Bank: 0,
        Inventory: make(map[string][]weapon.Weapon),
        Region: "North America", 
    }

    for m := range weapon.Models {
        p.Inventory[m] = make([]weapon.Weapon, 0)
    }
    var rows [][]string
    for wm, wl := range p.Inventory {
        rows = append(rows, []string{wm, strconv.Itoa(len(wl))}) 
    }
    p.InventoryTable = table.New().
	StyleFunc(func(row, col int) lipgloss.Style {
	    if row == 0 {
		return lipgloss.NewStyle().
		    Align(lipgloss.Center).
		    Bold(true)
	    } else {
	        return lipgloss.NewStyle().
                PaddingLeft(1)
	    }
	}).
        Border(lipgloss.NormalBorder()).
        BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
        Width(30).
        Headers("Model", "Qty").
        Rows(rows...)

    return p
}

func (p *Player) Move(region string) error {
    if slices.Contains(store.Regions, region) {
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


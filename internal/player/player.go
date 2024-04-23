package player

import (
    "errors"
    "strconv"
    "slices"

    "github.com/j-tew/warlord/internal/store"
    "github.com/j-tew/warlord/internal/weapon"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/table"
)

type Player struct {
    Name, Region string
    Health int8
    Cash, Bank int
    Inventory map[string][]weapon.Weapon
    InventoryTable table.Model
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
    var rows []table.Row
    for wm, wl := range p.Inventory {
        rows = append(rows, table.Row{wm, strconv.Itoa(len(wl))}) 
    }

    columns := []table.Column{
        {Title: "Model", Width: 10},
        {Title: "Qty", Width: 10},
    }

    p.InventoryTable = table.New(
            table.WithColumns(columns),
            table.WithRows(rows),
            table.WithHeight(15),
        )
    s := table.DefaultStyles()
    s.Header = s.Header.
            BorderStyle(lipgloss.NormalBorder()).
            BorderForeground(lipgloss.Color("240")).
            BorderBottom(true).
            Bold(false)
    s.Selected = s.Selected.
            Foreground(lipgloss.Color("229")).
            Background(lipgloss.Color("57")).
            Bold(false)
    p.InventoryTable.SetStyles(s)

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

func (p *Player) BuyWeapon(s *store.Store, model string, qty int) error {
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

func (p *Player) SellWeapon(s *store.Store, model string, qty int) error {
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


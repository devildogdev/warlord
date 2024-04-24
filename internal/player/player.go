package player

import (
    "errors"
    "strconv"
    "slices"

    "github.com/j-tew/warlord/internal/store"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/table"
)

type Player struct {
    Name, Region string
    Health int8
    Cash, Bank int
    Inventory map[string]int
    Table table.Model
}

func New(name string) *Player {
    p := &Player{
        Name: name,
        Health: 100,
        Cash: 15000,
        Bank: 0,
        Region: "North America", 
        Inventory: make(map[string]int),
    }

    var rows []table.Row

    for model, price := range store.Models {
        p.Inventory[model] = 0
        w := store.Weapon{Price: price, Qty: 0}
        rows = append(rows, table.Row{model, strconv.Itoa(w.Qty)}) 
    }

    columns := []table.Column{
        {Title: "Model", Width: 10},
        {Title: "Qty", Width: 10},
    }

    p.Table = table.New(
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
    p.Table.SetStyles(s)

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

func (p *Player) BuyWeapon(model *store.Weapon, qty int) error {
    if qty <= 0 {
        return errors.New("Quantity must be greater than Zero!")
    }
    cost := qty * model.Price
    if p.Cash <= cost {
        return errors.New("Not enough cash!")
    } else {
        p.Cash -= cost
        p.Inventory[model.Name] += qty
        model.Qty -= qty
    }
    return nil
}

func (p *Player) SellWeapon(s *store.Weapon, qty int) error {
    if len(p.Inventory) < 1 {
        return errors.New("You don't have any weapons to sell")
    } else if qty > len(p.Inventory) {
        return errors.New("You cannot sell more than you have")
    }
    p.Cash += s.Price
    p.Inventory[s.Name] -= qty
    s.Qty -= qty
    return nil
}

func (p *Player) Damage(value int8) {
    if value != p.Health {
        p.Health -= value
    } else {
        p.Health = 0
    }
}


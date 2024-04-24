package player

import (
    "errors"
    "strconv"
    "slices"

    "github.com/j-tew/warlord/internal/store"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/table"
    //"github.com/wk8/go-ordered-map/v2"
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

    for model := range store.Models {
        p.Inventory[model] = 0
    }

    p.UpdateTable()

    return p
}

func (p *Player) UpdateTable() {
    var rows []table.Row

    for model, qty := range p.Inventory {
        rows = append(rows, table.Row{model, strconv.Itoa(qty)}) 
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
}

func (p *Player) Move(region string) error {
    if slices.Contains(store.Regions, region) {
        p.Region = region
    } else {
        return errors.New("Invalid region")
    }
    return nil
}

func (p *Player) BuyWeapon(s *store.Store, w *store.Weapon, qty int) error {
    if qty <= 0 {
        return errors.New("Quantity must be greater than Zero!")
    }
    cost := qty * w.Price
    if p.Cash <= cost {
        return errors.New("Not enough cash!")
    } else {
        p.Cash -= cost
        p.Inventory[w.Name] += qty
        w.Qty -= qty
    }
    s.UpdateTable()
    p.UpdateTable()
    return nil
}

func (p *Player) SellWeapon(s *store.Store, w *store.Weapon, qty int) error {
    if p.Inventory[w.Name] < 1 {
        return errors.New("You don't have any weapons to sell")
    } else if qty > p.Inventory[w.Name] {
        return errors.New("You cannot sell more than you have")
    }
    p.Cash += w.Price
    p.Inventory[w.Name] -= qty
    w.Qty -= qty
    s.UpdateTable()
    return nil
}

func (p *Player) Damage(value int8) {
    if value != p.Health {
        p.Health -= value
    } else {
        p.Health = 0
    }
}


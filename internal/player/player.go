package player

import (
	"errors"
	"slices"
	"strconv"

	"github.com/j-tew/warlord/internal/store"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Player struct {
    Name, Region string
    Health int8
    Cash, Bank int
    Inventory map[string]int
    Table *table.Table
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

    for _, model := range store.Models {
        p.Inventory[model] = 0
    }

    p.UpdateTable()

    return p
}

func (p *Player) UpdateTable() {
    var rows [][]string

    for _, m := range store.Models {
        rows = append(rows, []string{m, strconv.Itoa(p.Inventory[m])}) 
    }

    p.Table = table.New().
	StyleFunc(func(row, col int) lipgloss.Style {
	    if row == 0 {
		return lipgloss.NewStyle().
		    Align(lipgloss.Center).
		    Bold(true)
	    } else if col == 1 {
	        return lipgloss.NewStyle().
                    Align(lipgloss.Center)
	    } else {
	        return lipgloss.NewStyle().
                    PaddingLeft(1)
            }
	}).
        Border(lipgloss.NormalBorder()).
        BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
        Width(20).
        Headers("Model", "Qty").
        Rows(rows...)
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


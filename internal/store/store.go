package store

import (
    "strconv"

    "github.com/j-tew/warlord/internal/weapon"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/lipgloss/table"
)

const maxInventory int = 100

var Regions = []string{
    "North America",
    "South America",
    "South East Asia",
    "North Africa",
    "Middle East",
    "Europe",
}

type Store struct {
    Region string
    Event string
    Inventory map[string][]weapon.Weapon
}

func New(region string) *Store {
    s := &Store{
        Region: region,
        Inventory : make(map[string][]weapon.Weapon),
    }
    for m, p := range weapon.Models {
        s.stockUp(m, p)
    }
    return s
}

func (s *Store) stockUp(model string, price int) {
    stock := make([]weapon.Weapon, maxInventory)
    for i := range stock {
        stock[i] = weapon.Weapon{Model: model, Price: price}
    }
    s.Inventory[model] = stock
}

func (s *Store) GetInventory() *table.Table {
    var rows [][]string
    for wm, wl := range s.Inventory {
        rows = append(rows, []string{wm, strconv.Itoa(len(wl)), strconv.Itoa(wl[0].Price)}) 
    }
    return table.New().
        Border(lipgloss.NormalBorder()).
        BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
        Headers("Model", "Qty", "Price").
        Rows(rows...)
}


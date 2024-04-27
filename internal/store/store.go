package store

import (
    "strconv"

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

var Models = map[string]int {
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

type Weapon struct {
    Name string
    Price, Qty int
}

type Store struct {
    Region string
    Event string
    Inventory map[string]*Weapon
    Table *table.Table
}

func New(region string) *Store {
    s := &Store{
        Region: region,
        Inventory: make(map[string]*Weapon),
    }

    for model, price  := range Models {
        w := &Weapon{Name: model, Price: price, Qty: maxInventory}
        s.Inventory[model] = w
    }

    s.UpdateTable()

    return s
}

func (s *Store) UpdateTable() {
    var rows [][]string

    for _, w := range s.Inventory {
        rows = append(rows, []string{w.Name, strconv.Itoa(w.Qty), strconv.Itoa(w.Price)})
    }

    s.Table = table.New().
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
        Headers("Model", "Qty", "Price").
        Rows(rows...)
}

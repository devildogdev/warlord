package store

import (
    "strconv"

    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/table"
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
    Table table.Model
}

func New(region string) *Store {
    s := &Store{
        Region: region,
        Inventory: make(map[string]*Weapon),
    }

    var rows []table.Row

    for model, price  := range Models {
        w := &Weapon{Name: model, Price: price, Qty: maxInventory}
        s.Inventory[model] = w
        rows = append(rows, table.Row{w.Name, strconv.Itoa(w.Price), strconv.Itoa(w.Qty)})
    }

    columns := []table.Column{
        {Title: "Model", Width: 10},
        {Title: "Qty", Width: 5},
        {Title: "Price", Width: 7},
    }

    s.Table = table.New(
        table.WithColumns(columns),
        table.WithRows(rows),
        table.WithHeight(15),
        )

    style := table.DefaultStyles()
    style.Header = style.Header.
            BorderStyle(lipgloss.NormalBorder()).
            BorderForeground(lipgloss.Color("240")).
            BorderBottom(true).
            Bold(false)
    style.Selected = style.Selected.
            Foreground(lipgloss.Color("229")).
            Background(lipgloss.Color("57")).
            Bold(false)
    s.Table.SetStyles(style)

    return s
}


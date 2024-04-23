package store

import (
    "strconv"

    "github.com/j-tew/warlord/internal/weapon"

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

type Store struct {
    Region string
    Event string
    Inventory map[string][]weapon.Weapon
    InventoryTable table.Model
}

func New(region string) *Store {
    s := &Store{
        Region: region,
        Inventory: make(map[string][]weapon.Weapon),
    }
    for m, p := range weapon.Models {
        s.stockUp(m, p)
    }
    var rows []table.Row
    for wm, wl := range s.Inventory {
        rows = append(rows, table.Row{wm, strconv.Itoa(len(wl)), strconv.Itoa(wl[0].Price)}) 
    }

    columns := []table.Column{
        {Title: "Model", Width: 10},
        {Title: "Qty", Width: 5},
        {Title: "Price", Width: 7},
    }
    s.InventoryTable = table.New(
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
    s.InventoryTable.SetStyles(style)

    return s
}

func (s *Store) stockUp(model string, price int) {
    stock := make([]weapon.Weapon, maxInventory)
    for i := range stock {
        stock[i] = weapon.Weapon{Model: model, Price: price}
    }
    s.Inventory[model] = stock
}


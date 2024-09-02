package store

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
    maxInventory int = 100
)

var (
    Regions = []string{
        "North America",
        "South America",
        "South East Asia",
        "North Africa",
        "Middle East",
        "Europe",
    }

    Models = []string {
        "G19",
        "1911",
        "M4",
        "AK-47",
        "RPG",
        "Mk19",
        "M24",
        "M107",
        "M2",
        "GAU-17",
    }

    Prices = map[string]int {
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
)

type Weapon struct {
    Name string
    Price, Qty int
}

type Store struct {
    Region string
    Event string
    Inventory map[string]*Weapon
    Prices map[string]int
    Table *table.Table
}

func New(region string) *Store {
    s := &Store{
        Region: region,
        Inventory: make(map[string]*Weapon),
    }

    for _, model := range Models {
        base := Prices[model]
        src := rand.NewPCG(uint64(time.Now().Unix()), uint64(base))
        r := rand.New(src)
        upper, lower := int(float64(base) * float64(1.5)), int(float64(base) * float64(0.5))
        price := r.IntN(upper - lower) + lower
        w := &Weapon{Name: model, Price: price, Qty: maxInventory}
        s.Inventory[model] = w
    }

    s.UpdateTable()

    return s
}

func (s *Store) UpdateTable() {
    var rows [][]string

    for _, m := range Models {
        w := s.Inventory[m]
        rows = append(
            rows, []string{
                w.Name,
                strconv.Itoa(w.Qty),
                fmt.Sprintf("$%d", w.Price),
            },
        )
    }

    s.Table = table.New().
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
        BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("2"))).
        Width(30).
        Headers("Model", "Qty", "Price").
        Rows(rows...)
}

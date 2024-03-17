package main

import (
	"fmt"
	"os"

	"github.com/j-tew/warlord/internal/player"
	"github.com/j-tew/warlord/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
    intro string =  `
Warlord

You are a small time arms dealer, trying to make
a name for yourself. To get you started, you get
a little capital from an "investor". They aren't
exactly a Credit Union, so this is going to cost
you. Keep an eye on your debt. The interest tacks
on each week. You have one year (52 weeks) build
your fortune.

Watch out for law enforcement!

`
    weeks int = 52
    maxInvetory int = 100
)

type storage interface {
    InventoryTable() *table.Table
}

func GetInventory(s storage) string {
    return s.InventoryTable().Render()
}

type model struct {
    player *player.Player
    store *store.Store
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, cmd
}

func (m model) View() string {
    labelStyle := lipgloss.NewStyle().
        Align(lipgloss.Center).
        Bold(true)

    s := lipgloss.JoinVertical(
        lipgloss.Center,
        labelStyle.Render(m.player.Region),
        GetInventory(m.store),
    )

    p := lipgloss.JoinVertical(
        lipgloss.Center,
        labelStyle.Render(m.player.Name),
        GetInventory(m.player),
    )
    return lipgloss.JoinVertical(lipgloss.Left, s, p)
}

func main() {
    p := player.New("Outlaw")
    s := store.New(p.Region)

    m := model{
        player: p,
        store: s,
    }
    if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}

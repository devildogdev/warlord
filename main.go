package main

import (
	"fmt"
	"os"

	"github.com/j-tew/warlord/internal/player"
	"github.com/j-tew/warlord/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
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

var width, height int

type storage interface {
    InventoryTable() *table.Table
}

func getInventory(s storage, label string) string {
    tableStyle := lipgloss.NewStyle().Margin(5)
    labelStyle := lipgloss.NewStyle().
        MarginBottom(1).
        Bold(true)

    return lipgloss.JoinVertical(
        lipgloss.Center,
        labelStyle.Render(label),
        tableStyle.Render(s.InventoryTable().Render()),
    )
}

type model struct {
    player *player.Player
    store  *store.Store
    list   list.Model
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
    case tea.WindowSizeMsg:
        width = msg.Width
        height = msg.Height
    }
    return m, cmd
}

func (m model) View() string {
    s := getInventory(m.store, m.store.Region)

    p := getInventory(m.player, m.player.Name)

    screenStyle := lipgloss.NewStyle().
        Width(width).
        Height(height).
        Align(lipgloss.Center, lipgloss.Center)
    return screenStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, s, p))
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

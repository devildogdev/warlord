package main

import (
    "fmt"
    "os"

    "github.com/j-tew/warlord/internal/player"
    "github.com/j-tew/warlord/internal/store"

    tea "github.com/charmbracelet/bubbletea"
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
    ShowInventory()
}

type model struct {
    playerTable *table.Table
    storeTable *table.Table
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
    return m.storeTable.Render() + "\n" + m.playerTable.Render()
}

func main() {
    p := player.New("Outlaw")
    // Not sure I like having stores in player package
    s := store.New(p.Region)

    m := model{
        playerTable: p.GetInventory(),
        storeTable: s.GetInventory(),
    }
    if _, err := tea.NewProgram(m).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}

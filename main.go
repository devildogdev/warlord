package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/j-tew/warlord/internal/player"

    "github.com/charmbracelet/bubbles/table"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
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

// Need to look at lipgloss docs
var baseStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("240"))

type model struct {
    playerTable table.Model
    storeTable table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "esc":
            if m.playerTable.Focused() {
                // Blur not working. May be terminal.
                m.storeTable.Focus()
            } else if m.storeTable.Focused() {
                m.playerTable.Focus()
            } else {
                m.playerTable.Focus()
            }
        case "q", "ctrl+c":
            return m, tea.Quit
        case "enter":
            if m.playerTable.Focused() {
                return m, tea.Batch(
                    tea.Printf("You bought a %s!", m.playerTable.SelectedRow()[0]),
                )
            }
        }
    }
    m.playerTable, cmd = m.playerTable.Update(msg)
    m.storeTable, cmd = m.storeTable.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return baseStyle.Render(m.playerTable.View()) + "\n" + baseStyle.Render(m.storeTable.View()) + "\n"
}

func main() {
    p := player.New("Outlaw")
    // Not sure I like having stores in player package
    st := player.Stores[p.Region]

    columns := []table.Column{
        {Title: "Model", Width: 6},
        {Title: "Qty", Width: 3},
        {Title: "Price", Width: 6},
    }

    var rows []table.Row
    for wm, wl := range st.Inventory {
        // Row only accepts strings
        rows = append(rows, table.Row{wm, strconv.Itoa(len(wl)), strconv.Itoa(wl[0].Price)}) 
    }

    // Figure out how to render more than one table
    storeT := table.New(
        table.WithColumns(columns),
        table.WithRows(rows),
        table.WithFocused(true),
        table.WithHeight(7),
    )

    playerT := table.New(
        table.WithColumns(columns),
        table.WithRows(rows),
        table.WithFocused(true),
        table.WithHeight(7),
    )

    s := table.DefaultStyles()
    s.Header = s.Header.
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("240")).
        BorderBottom(true).
        Bold(true)
    s.Selected = s.Selected.
        Foreground(lipgloss.Color("229")).
        Background(lipgloss.Color("57")).
        Bold(true)
    storeT.SetStyles(s)
    playerT.SetStyles(s)

    m := model{playerT, storeT}
    if _, err := tea.NewProgram(m).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}

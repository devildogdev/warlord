package main

import (
    "fmt"
    "io"
    "os"
    "strings"

    "github.com/j-tew/warlord/internal/player"
    "github.com/j-tew/warlord/internal/store"

    "github.com/charmbracelet/bubbles/list"
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

var width, height int

var (
    itemStyle = lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)
    selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170")).AlignHorizontal(lipgloss.Center)
    tableStyle = lipgloss.NewStyle().Margin(5)
    labelStyle = lipgloss.NewStyle().
        MarginBottom(1).
        Bold(true)
    mainMenu = []list.Item{
        item("Buy"),
        item("Sell"),
        item("Travel"),
    }
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int { return 1 }
func (d itemDelegate) Spacing() int { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
    i, ok := listItem.(item)
    if !ok {
            return
    }

    str := string(i)

    fn := itemStyle.Render
    if index == m.Index() {
            fn = func(s ...string) string {
                    return selectedItemStyle.Render(strings.Join(s, " "))
            }
    }

    fmt.Fprint(w, fn(str))
}

type model struct {
    player *player.Player
    store  *store.Store
    list   list.Model
    state string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "enter", "l":
            i, ok := m.list.SelectedItem().(item)
            if ok {
                m.state = string(i)
            }
        case "backspace", "h":
            if m.state != "" {
                m.state = ""
            }
        }
    case tea.WindowSizeMsg:
        width = msg.Width
        height = msg.Height
    }

    menu := []list.Item{}
    switch m.state {
    case "Buy":
        for m := range m.store.Inventory {
            menu = append(menu, item(m))
        }
        m.list = list.New(menu, itemDelegate{}, 10, 15)
    case "Sell":
        m.list = list.New(menu, itemDelegate{}, 10, 15)
        for m, q := range m.player.Inventory {
            if q > 0 {
                menu = append(menu, item(m))
            }
        }
    default:
        m.list = list.New(mainMenu, itemDelegate{}, 10, 15)
    }

    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m model) View() string {
    s := m.store
    p := m.player

    m.list.SetShowHelp(false)
    m.list.SetShowTitle(false)
    m.list.SetShowStatusBar(false)
    m.list.SetFilteringEnabled(false)

    choices := m.list.View()
    choicesStyle := lipgloss.NewStyle().
        AlignHorizontal(lipgloss.Left).
        MarginTop(5)

    stats := fmt.Sprintf("Health: %d\nBank: $%d\nCash: $%d", p.Health, p.Bank, p.Cash)
    statsStyle := lipgloss.NewStyle().
        AlignHorizontal(lipgloss.Right).
        MarginTop(5)

    screenStyle := lipgloss.NewStyle().
        Width(width).
        Height(height).
        Align(lipgloss.Center, lipgloss.Center)

    playerStyle := lipgloss.NewStyle().MarginLeft(5)

    playerTable := playerStyle.Render(
        lipgloss.JoinVertical(
        lipgloss.Center,
        labelStyle.Render(p.Name),
        p.Table.Render(),
        ),
    )

    storeStyle := lipgloss.NewStyle().MarginRight(5)

    storeTable := storeStyle.Render(
        lipgloss.JoinVertical(
        lipgloss.Center,
        labelStyle.Render(s.Region),
        s.Table.Render(),
        ),
    )

    tables := lipgloss.JoinHorizontal(
        lipgloss.Top,
        lipgloss.JoinVertical(lipgloss.Left, storeTable, choicesStyle.Render(choices)),
        lipgloss.JoinVertical(lipgloss.Right, playerTable, statsStyle.Render(stats)),
    )

    return screenStyle.Render(tables)
}

func main() {
    p := player.New("Outlaw")
    s := store.New(p.Region)

    m := model{
        player: p,
        store: s,
        state: "",
        list: list.New(mainMenu, itemDelegate{}, 15, 10),
    }

    if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}

package main

import (
    "fmt"
    "os"

    "github.com/j-tew/warlord/internal/ui"
    "github.com/j-tew/warlord/internal/player"
    "github.com/j-tew/warlord/internal/store"

    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type state int

const (
    nav state = iota
    intro
    buy 
    sell
    travel
    event
)

var (
    storeStyle = lipgloss.NewStyle().
                    MarginLeft(5).
                    MarginBottom(1)
    playerStyle = lipgloss.NewStyle().
                    MarginRight(5).
                    MarginBottom(1)
    labelStyle = lipgloss.NewStyle().
                    MarginBottom(1).
                    Bold(true)
    choicesStyle = lipgloss.NewStyle().
                    AlignHorizontal(lipgloss.Left)
    statsStyle = lipgloss.NewStyle().
                    AlignHorizontal(lipgloss.Right)
)

var width, height int

type Model struct {
    player  *player.Player
    store  *store.Store
    state   state
    list    list.Model
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        width = msg.Width
        height = msg.Height
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "enter":
            i, ok := m.list.SelectedItem().(ui.Item)
            s := string(i)
            if ok {
                switch m.state {
                case intro:
                    m.state = nav
                    m.list = ui.MainMenu()
                case nav:
                    switch s {
                    case "Buy":
                        m.state = buy
                        m.list = ui.BuyMenu()
                    case "Sell":
                        m.state = sell 
                        m.list = ui.SellMenu(m.player.Inventory)
                    case "Travel":
                        m.state = travel 
                        m.list = ui.TravelMenu()
                    }
                case buy:
                    m.player.BuyWeapon(m.store, m.store.Inventory[string(i)], 1)
                case sell:
                    m.player.SellWeapon(m.store, m.store.Inventory[string(i)], 1)
                case travel:
                    r := string(i)
                    m.player.Move(r)
                    m.store = store.New(r)
                    m.state = nav
                    m.list = ui.MainMenu()
                }
            }
        }
    }

    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m Model) View() string {
    var layout string

    if m.state == intro {
        layout = ui.Intro()
    } else {

        s := m.store
        p := m.player

        m.list.SetShowHelp(false)
        m.list.SetShowTitle(false)
        m.list.SetShowStatusBar(false)
        m.list.SetFilteringEnabled(false)

        choices := choicesStyle.Render(m.list.View())

        stats := statsStyle.Render(
            fmt.Sprintf(
                "Week: %d\nCash: $%d\nHealth: %d",
                p.Week,
                p.Cash,
                p.Health,
            ),
        )

        playerTable := playerStyle.Render(
            lipgloss.JoinVertical(
            lipgloss.Center,
            labelStyle.Render(p.Name),
            p.Table.Render(),
            ),
        )

        storeTable := storeStyle.Render(
            lipgloss.JoinVertical(
            lipgloss.Center,
            labelStyle.Render(p.Region),
            s.Table.Render(),
            ),
        )

        layout = lipgloss.JoinHorizontal(
            lipgloss.Top,
            lipgloss.JoinVertical(lipgloss.Left, playerTable, choices),
            lipgloss.JoinVertical(lipgloss.Right, storeTable, stats),
        )
    }

    return lipgloss.NewStyle().
                    Width(width).
                    Height(height).
                    Align(lipgloss.Center, lipgloss.Center).
                    Render(layout)
}

func main() {
    p := player.New("Outlaw")
    s := store.New(p.Region)

    m := Model{
        player: p,
        store: s,
        state: intro,
        list: ui.MainMenu(),
    }

    if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}

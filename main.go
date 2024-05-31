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

var (
    storeStyle = lipgloss.NewStyle().
                    MarginRight(5).
                    MarginBottom(1)
    playerStyle = lipgloss.NewStyle().
                    MarginLeft(5).
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

type mainModel struct {
    player  *player.Player
    store   *store.Store
    menu    string
    list    list.Model
}

func (m mainModel) Init() tea.Cmd { return nil }

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        width = msg.Width
        height = msg.Height
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
        switch m.menu {
        case "Intro":
            switch msg.String() {
            case "enter":
                m.menu = "Main"
                m.list = ui.MainMenu()
            }
        case "Main":
            switch msg.String() {
            case "q", "ctrl+c":
                return m, tea.Quit
            case "enter":
                i, ok := m.list.SelectedItem().(ui.Item)
                if ok {
                    switch string(i) {
                    case "Buy":
                        m.menu = "Buy" 
                        m.list = ui.BuyMenu()
                    case "Sell":
                        m.menu = "Sell" 
                        m.list = ui.SellMenu(m.player.Inventory)
                    case "Travel":
                        m.menu = "Travel" 
                        m.list = ui.TravelMenu()
                    }
                }
            }
        case "Buy":
            switch msg.String() {
            case "q", "ctrl+c":
                return m, tea.Quit
            case "enter":
                i, ok := m.list.SelectedItem().(ui.Item)
                if ok {
                    m.player.BuyWeapon(m.store, m.store.Inventory[string(i)], 1)
                }
            case "backspace":
                if m.menu != "Main" {
                    m.menu = "Main"
                    m.list = ui.MainMenu()
                }
            }
        case "Sell":
            switch msg.String() {
            case "q", "ctrl+c":
                return m, tea.Quit
            case "enter":
                i, ok := m.list.SelectedItem().(ui.Item)
                if ok {
                    m.player.SellWeapon(m.store, m.store.Inventory[string(i)], 1)
                }
            case "backspace":
                if m.menu != "Main" {
                    m.menu = "Main"
                    m.list = ui.MainMenu()
                }
            }
        case "Travel":
            switch msg.String() {
            case "q", "ctrl+c":
                return m, tea.Quit
            case "enter":
                i, ok := m.list.SelectedItem().(ui.Item)
                if ok {
                    r := string(i)
                    m.player.Move(r)
                    m.store = store.New(r)
                }
            case "backspace":
                if m.menu != "Main" {
                    m.menu = "Main"
                    m.list = ui.MainMenu()
                }
            }
        }
    }

    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m mainModel) View() string {
    var layout string

    if m.menu == "Intro" {
        layout = ui.Intro
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
                "Week: %d\nHealth: %d\nCash: $%d",
                p.Week,
                p.Health,
                p.Cash,
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
            lipgloss.JoinVertical(lipgloss.Left, storeTable, choices),
            lipgloss.JoinVertical(lipgloss.Right, playerTable, stats),
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

    m := mainModel{
        player: p,
        store: s,
        menu: "Intro",
        list: ui.MainMenu(),
    }

    if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}

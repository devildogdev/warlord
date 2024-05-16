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

const (
    mainView sessionState = iota
)

var (
    storeStyle = lipgloss.NewStyle().MarginRight(5)
    playerStyle = lipgloss.NewStyle().MarginLeft(5)
    labelStyle = lipgloss.NewStyle().
                    MarginBottom(1).
                    Bold(true)
    choicesStyle = lipgloss.NewStyle().
                    AlignHorizontal(lipgloss.Left).
                    MarginTop(5)
    statsStyle = lipgloss.NewStyle().
                    AlignHorizontal(lipgloss.Right).
                    MarginTop(5)
    screenStyle = lipgloss.NewStyle().
                    Width(width).
                    Height(height).
                    Align(lipgloss.Center, lipgloss.Center)
)

var width, height int

type sessionState int

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
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
        switch m.menu {
        case "Main":
            switch msg.String() {
            case "q", "ctrl+c":
                return m, tea.Quit
            case "enter", "l":
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
                        m.list = ui.TravelMenu(m.player.Inventory)
                    }
                }
            }
        case "Buy":
            switch msg.String() {
            case "q", "ctrl+c":
                return m, tea.Quit
            case "enter", "l":
                i, ok := m.list.SelectedItem().(ui.Item)
                if ok {
                    m.player.BuyWeapon(m.store, m.store.Inventory[string(i)], 1)
                }
            case "backspace", "h":
                if m.menu != "Main" {
                    m.menu = "Main"
                    m.list = ui.MainMenu()
                }
            }
        case "Sell":
            switch msg.String() {
            case "q", "ctrl+c":
                return m, tea.Quit
            case "enter", "l":
                i, ok := m.list.SelectedItem().(ui.Item)
                if ok {
                    m.player.SellWeapon(m.store, m.store.Inventory[string(i)], 1)
                }
            case "backspace", "h":
                if m.menu != "Main" {
                    m.menu = "Main"
                    m.list = ui.MainMenu()
                }
            }
        case "Travel":
            switch msg.String() {
            case "q", "ctrl+c":
                return m, tea.Quit
            case "enter", "l":
                i, ok := m.list.SelectedItem().(ui.Item)
                if ok {
                    m.player.Move(string(i))
                }
            case "backspace", "h":
                if m.menu != "Main" {
                    m.menu = "Main"
                    m.list = ui.MainMenu()
                }
            }
        }
    case tea.WindowSizeMsg:
        width = msg.Width
        height = msg.Height
    }

    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m mainModel) View() string {
    s := m.store
    p := m.player

    m.list.SetShowHelp(false)
    m.list.SetShowTitle(false)
    m.list.SetShowStatusBar(false)
    m.list.SetFilteringEnabled(false)

    choices := m.list.View()

    stats := fmt.Sprintf("Health: %d\nBank: $%d\nCash: $%d", p.Health, p.Bank, p.Cash)

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

    m := mainModel{
        player: p,
        store: s,
        menu: "Main",
        list: ui.MainMenu(),
    }

    if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}

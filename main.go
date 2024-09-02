package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/devildogdev/warlord/internal/player"
	"github.com/devildogdev/warlord/internal/store"
	"github.com/devildogdev/warlord/internal/ui"

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
	law
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
	player *player.Player
	store  *store.Store
	state  state
	list   list.Model
	event  bool
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width = msg.Width
		height = msg.Height
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
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
					m.player.Move(s)
					m.store = store.New(s)
					src := rand.NewPCG(uint64(time.Now().Unix()), uint64(10))
					r := rand.New(src)
					// Randomizing event trigger
					if r.IntN(100)%3 == 0 {
						m.event = true
						m.state = law
						m.list = ui.LawMenu()
					} else {
						m.state = nav
						m.list = ui.MainMenu()
					}
				case law:
					switch s {
					case "Run":
						p := m.player
						if p.Escape() {
							m.state = nav
							m.list = ui.MainMenu()
							m.event = false
						} else {
							p.Damage(5)
						}
					case "Bribe":
						p := m.player
						if p.Bribe(4000) {
							m.state = nav
							m.list = ui.MainMenu()
							m.event = false
						}
						// case "Attack":
						//     m.state = travel
						//     m.list = ui.TravelMenu()
					}
				}
			}
		case tea.KeyBackspace:
			if !m.event {
				m.state = nav
				m.list = ui.MainMenu()
			}
		}
	}

	m.list.SetShowHelp(false)
	m.list.SetShowTitle(false)
	m.list.SetShowStatusBar(false)
	m.list.SetFilteringEnabled(false)

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var layout string

	if m.state == intro {
		layout = ui.Intro()
	} else if m.state == law {
		layout = lipgloss.JoinVertical(lipgloss.Center, ui.LawWarning(m.player), m.list.View())
	} else {

		s := m.store
		p := m.player

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
		store:  s,
		state:  intro,
		list:   ui.MainMenu(),
		event:  false,
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

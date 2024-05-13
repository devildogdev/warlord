package ui

import (
    "io"
    "fmt"
    "strings"

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

var (
    itemStyle = lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)
    selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170")).AlignHorizontal(lipgloss.Center)
    tableStyle = lipgloss.NewStyle().Margin(5)
    MainMenu = list.New([]list.Item{
            Item("Buy"),
            Item("Sell"),
            Item("Travel"),
        },
        ItemDelegate{},
        10,
        15,
    )
)

type Item string

func (i Item) FilterValue() string { return "" }

type ItemDelegate struct{}

func (d ItemDelegate) Height() int { return 1 }
func (d ItemDelegate) Spacing() int { return 0 }
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
    i, ok := listItem.(Item)
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

// Update
//    for m := range m.store.Inventory {
//        storeGuns = append(storeGuns, item(m))
//    }
//
//    m.list.SetShowHelp(false)
//    m.list.SetShowTitle(false)
//    m.list.SetShowStatusBar(false)
//    m.list.SetFilteringEnabled(false)


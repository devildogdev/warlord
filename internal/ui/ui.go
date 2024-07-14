package ui

import (
	"fmt"
	"io"
	"strings"

    "github.com/j-tew/warlord/internal/store"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
    title string = `
██╗    ██╗ █████╗ ██████╗ ██╗      ██████╗ ██████╗ ██████╗ 
██║    ██║██╔══██╗██╔══██╗██║     ██╔═══██╗██╔══██╗██╔══██╗
██║ █╗ ██║███████║██████╔╝██║     ██║   ██║██████╔╝██║  ██║
██║███╗██║██╔══██║██╔══██╗██║     ██║   ██║██╔══██╗██║  ██║
╚███╔███╔╝██║  ██║██║  ██║███████╗╚██████╔╝██║  ██║██████╔╝
 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═════╝ 
                                                           
`
    story string =  `

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
    selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).AlignHorizontal(lipgloss.Center)
    tableStyle = lipgloss.NewStyle().Margin(5)
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

func Intro() string {
    titleStyle := lipgloss.NewStyle().
                    Bold(true).
                    Foreground(lipgloss.Color("2"))
    return lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(title), story)
}

func MainMenu() list.Model {
    return list.New([]list.Item{
            Item("Buy"),
            Item("Sell"),
            Item("Travel"),
        },
        ItemDelegate{},
        10,
        15,
    )
}

func BuyMenu() list.Model {
    opts := []list.Item{}
    for _, m := range store.Models {
        opts = append(opts, Item(m))
    }
    return list.New(opts, ItemDelegate{}, 10, 15)
}

func SellMenu(inventory map[string]int) list.Model {
    opts := []list.Item{}
    for _, m := range store.Models {
        if inventory[m] > 0 {
            opts = append(opts, Item(m))
        }
    }
    return list.New(opts, ItemDelegate{}, 10, 15)
}

func TravelMenu() list.Model {
    opts := []list.Item{}
    for _, r := range store.Regions {
        opts = append(opts, Item(r))
    }
    return list.New(opts, ItemDelegate{}, 10, 15)
}


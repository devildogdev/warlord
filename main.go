package main

import (
	"fmt"
	// "math/rand"

	"github.com/j-tew/warlord/internal/player"
	"github.com/j-tew/warlord/internal/store"
)

const intro string =  `
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

const weeks int = 52
const maxInvetory int = 100

var locations =  map[string]store.Store{
    "North America": store.New("North America"),
    "South America": store.New("South America"),
    "Europe": store.New("Europe"),
    "North Africa": store.New("North Africa"),
    "South East Asia": store.New("South East Asia"),
    "Middle East": store.New("Middle East"),
}

func main() {
    var name string

    // fmt.Print(intro)
    // fmt.Println("What is your name?")
    // fmt.Scanln(&name)

    p := player.New(name)
    s := locations[p.Location]

    fmt.Printf("%s: %v\n", s.Location, s.Inventory)
    fmt.Printf("Player: %v\n", p)
    p.BuyWeapon(s, "M4", 2)
    fmt.Println("********************")
    fmt.Printf("%s: %v\n", s.Location, s.Inventory)
    fmt.Printf("Player: %v\n", p)
    p.Damage(20)
    p.SellWeapon(s, "M4", 1)
    fmt.Printf("Player: %v\n", p)
}

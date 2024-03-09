package main

import (
    "fmt"
    // "math/rand"

    "github.com/j-tew/warlord/internal/player"
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

type storage interface {
    ShowInventory()
}

func main() {
    var name string

    fmt.Print(intro)
    fmt.Println("What is your name?")
    fmt.Scanln(&name)

    p := player.New(name)
    s := player.Stores[p.Region]

    fmt.Println("North America Store")
    fmt.Println("********************")
    s.ShowInventory()

    p.BuyWeapon(s, "M4", 2)
    fmt.Printf("%s bought 2 M4s\n\n", p.Name)

    fmt.Println("Store's inventory after purchase")
    s.ShowInventory()
    fmt.Println()

    fmt.Printf("\n%s's Inventory\n", p.Name) 
    fmt.Println("********************")
    p.ShowInventory()

    fmt.Println()
    p.SellWeapon(s, "M4", 1)
    fmt.Printf("%s sold an M4\n\n", p.Name)
    p.ShowInventory()
}

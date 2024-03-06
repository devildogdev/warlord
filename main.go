package main

import (
    "fmt"
    // "math/rand"
    
    "github.com/j-tew/warlord/player"
    "github.com/j-tew/warlord/store"
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

var locations =  [6]string{
    "North America",
    "South America",
    "Europe",
    "North Africa",
    "South East Asia",
    "Middle East",
}

func main() {
    var name string

    // fmt.Print(intro)
    // fmt.Println("What is your name?")
    // fmt.Scanln(&name)

    p := player.New(name)
    s := store.New(locations[0])

    fmt.Println(s)
    fmt.Println(p)
    p.BuyWeapon(s, 2)
    fmt.Println(s)
    fmt.Println(p)
    p.Damage(20)
    p.SellWeapon(s, 1)
    fmt.Println(p)
}

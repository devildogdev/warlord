package store

type Weapon struct {
    Model string
    Price int
}

type Store struct {
    Location string
    Inventory map[string][]Weapon
}

var models = map[string]int{
    "G19": 600,
    "1911": 1000,
    "M4": 1400,
    "AK-47": 2200,
    "RPG": 15000,
    "Mk19": 41000,
    "M24": 18000,
    "M107": 27000,
    "M2": 62000,
    "GAU-17": 250000,
}

func New(location string) Store {
    s := Store{
        Location: location,
        Inventory: make(map[string][]Weapon),
    }
    for m, p := range models {
        s.stockUp(m, p, 10)
    }
    return s
}

func (s *Store) stockUp(model string, price int, qty int) {
    for i := 0; i < qty; i++ {
        s.Inventory[model] = append(s.Inventory[model], Weapon{Model: model, Price: price})
    }
}

package store

type Weapon struct {
    Model string
    Price int
}

type Store struct {
    Location string
    Inventory []Weapon
}

func New(location string) *Store {
    return &Store{
        Location: location,
        Inventory: []Weapon{
            {
                Model: "G19",
                Price: 600,
            },
            {
                Model: "1911",
                Price: 1000,
            },
            {
                Model: "M4",
                Price: 1400,
            },
            {
                Model: "AK-47",
                Price: 2200,
            },
            {
                Model: "RPG",
                Price: 15000,
            },
            {
                Model: "Mk19",
                Price: 41000,
            },
            {
                Model: "M24",
                Price: 18000,
            },
            {
                Model: "M107",
                Price: 27000,
            },
            {
                Model: "M2",
                Price: 62000,
            },
            {
                Model: "GAU-17",
                Price: 250000,
            },
        },
    }
}


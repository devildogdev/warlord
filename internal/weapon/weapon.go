package weapon

type Weapon struct {
    Model string
    Price int
}

var Models = map[string]int{
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


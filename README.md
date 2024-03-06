# Warlord

A spin on the classic "Drug Wars" where you play as an arms dealer instead of a drug dealer.

## *work in progress*

### Mechanics

- Loan shark
- Inventory and Marketplace
- Travel between locations each day
- Random events
- Final score

### Back Story

You are a small time arms dealer, trying to make a name for yourself. To get you started, you get a little capital from
an *"investor"*. They aren't exactly a Credit Union, so this is going to cost you. Keep an eye on your debt. The interest tacks on each week.
You have one year (52 weeks) build your fortune. Watch out for law enforcement.

### Stating Funds

*$15,000* with *45%* interest (or something)

### Weapons

- Handguns
    - G19
    - 1911

- Infantry Rifles
    - AK-47
    - M4

- Sniper Rifles
    - M24 Sniper Rifle
    - M107 .50 Cal Rifle

- Explosives
    - RPG
    - MK19 Grenade Launcher

- Machine Guns
    - M2 .50 Cal Machine Gun
    - GAU-17 Gatling Gun

### Locations

- North America
- South America
- South East Asia
- Middle East
- Europe
- North Africa

### Events

- Market
    - Shortage of ammo
    - Firearm bans
    - Civil conflict
    - Looting
    - Political unrest

- Law Enforcement
    - Interpol
    - ATF
    - FBI

- Evasion:
    - Run
    - Bribe
    - Attack

### Duration

*52* weeks

## To Do & Notes

- [X] Setup guns and player
- [X] Implement store and buying and selling functionality
- [ ] Factor multipliers for prices
- [ ] Random Events
- [ ] 

Need a way to create the stores, based on location.
Create one every time location changes?
Create all stores on init, then alter prices?
One store map, location as keys, inventory as map with model as keys and slice of weapons.

## Store
{
    location: "NA",
    inventory: {
        w1{, w2, w3}
}

stores := map\[location\]store{
    "North America": \[\]weapon{
        {
            model: "G19",
            price: 600,
        }
    }
}

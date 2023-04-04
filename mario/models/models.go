
package models

type Player struct {
    Lives int
    Stage int
    Inventory []string
}

// recv function to add a life
func (p *models.Player) Greenmushroom() {
    p.Lives++
}

// recv function to add an inventory item
func (p *Player) Pickup(powerup string) {
    p.Inventory = append(p.Inventory, powerup)
}

// rec function to check on the current stasge
func (p Player) CanWhistle() bool {
    return p.Stage >= 5
    }

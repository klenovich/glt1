package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type Player struct {
	Health  int
	Room    *Room
	HasPotion, HasKey, HasSword bool
}

type Room struct {
	Description string
	Item        *Item
	Enemy       *Enemy
	Exits       map[string]*Room
	IsLocked    bool
}

type Item struct {
	Name string
	HealthBoost int
}

type Enemy struct {
	Health int
	Attack int
	Name string
}

type Game struct {
	Player *Player
	Rooms  []*Room
}

func (p *Player) Move(direction string) {
	if newRoom, ok := p.Room.Exits[direction]; ok {
		if newRoom.IsLocked {
			if p.HasKey {
				newRoom.IsLocked = false
			} else {
				fmt.Println("The door is locked.")
				return
			}
		}

		p.Room = newRoom
		fmt.Println(p.Room.Description)
		
		if p.Room.Item != nil {
			itemColor := color.New(color.FgGreen).PrintfFunc()
			itemColor("You found a %s!\n", p.Room.Item.Name)

			switch p.Room.Item.Name {
			case "Healing Potion":
				p.Health += p.Room.Item.HealthBoost
				fmt.Printf("Your health has increased by %d points.\n", p.Room.Item.HealthBoost)
				p.HasPotion = true
			case "Key":
				p.HasKey = true
			case "Sword":
				p.HasSword = true
			}

			p.Room.Item = nil
		}

		if p.Room.Enemy != nil {
			enemyColor := color.New(color.FgRed).PrintfFunc()
			enemyColor("You've encountered a %s! ", p.Room.Enemy.Name)

			if p.HasSword {
				fmt.Println("You draw your sword and prepare to fight.")
				p.Room.Enemy.Health -= p.Room.Enemy.Attack
			} else {
				fmt.Printf("It attacks you for %d damage.\n", p.Room.Enemy.Attack)
				p.Health -= p.Room.Enemy.Attack
			}
			
			if p.Room.Enemy.Health <= 0 {
				fmt.Println("You defeated the enemy!")
				p.Room.Enemy = nil
			}
		}
	} else {
		fmt.Println("You can't go that way.")
	}
}

func main() {
	room1 := &Room{Description: "You're in a dark room. There's an exit to the north.", Exits: make(map[string]*Room)}
	room2 := &Room{Description: "You've entered a brightly lit room. Exits are to the south and east.",
		Item: &Item{Name: "Healing Potion", HealthBoost: 50}, Exits: make(map[string]*Room)}
	room3 := &Room{Description: "This room looks ominous. There's an exit to the west.",
		Enemy: &Enemy{Name: "Goblin", Health: 50, Attack: 20}, Exits: make(map[string]*Room)}
	room4 := &Room{Description: "It looks like a storage room. There's an exit to the west.",
		Item: &Item{Name: "Key", HealthBoost: 0}, IsLocked: true, Exits: make(map[string]*Room)}
	room5 := &Room{Description: "This is a mighty messy room but you see a sword in the corner. Exits are to the east and south.",
		Item: &Item{Name: "Sword", HealthBoost: 0}, Exits: make(map[string]*Room)}
	
	room1.Exits["north"] = room2
	room2.Exits["south"] = room1
	room2.Exits["east"] = room3
	room3.Exits["west"] = room2
	room3.Exits["east"] = room4
	room4.Exits["west"] = room3
	room4.Exits["north"] = room5
	room5.Exits["south"] = room4
	room5.Exits["east"] = room2

	game := &Game{
		Player: &Player{Health: 100, Room: room1},
		Rooms:  []*Room{room1, room2, room3, room4, room5},
	}

	for game.Player.Health > 0 {
		var cmd string
		fmt.Println("Enter a direction (north, south, east, west):")
		fmt.Scanf("%s\n", &cmd)
		cmd = strings.ToLower(cmd)

		switch cmd {
		case "north", "south", "east", "west":
			game.Player.Move(cmd)
		default:
			fmt.Println("That is not a valid direction.")
		}

		playerHealthColor := color.New(color.FgYellow).PrintfFunc()
		playerHealthColor("Your current health is: %d\n", game.Player.Health)

		if game.Player.Health <= 0 {
			fmt.Println("You have died... ")
			return
		}

		if game.Player.HasKey && game.Player.HasPotion && game.Player.HasSword {
			gameWinColor := color.New(color.FgBlue).PrintfFunc()
			gameWinColor("Congrats, you've collected all items and won the game!\n")
			break
		}
	
	}
}
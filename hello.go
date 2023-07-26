package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type Player struct {
	Health  int
	Room    *Room
	HasPotion bool
}

type Room struct {
	Description string
	Item        *Item
	Enemy       *Enemy
	Exits       map[string]*Room
}

type Game struct {
	Player *Player
	Rooms  []*Room
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

func (p *Player) Move(direction string) {
	if newRoom, ok := p.Room.Exits[direction]; ok {
		p.Room = newRoom
		fmt.Println(p.Room.Description)
		
		if p.Room.Item != nil {
			itemColor := color.New(color.FgGreen).PrintfFunc()
			itemColor("You found a %s! Your health has increased by %d points.", p.Room.Item.Name, p.Room.Item.HealthBoost)
			p.HasPotion = true
			p.Health += p.Room.Item.HealthBoost
			p.Room.Item = nil
		}
		
		if p.Room.Enemy != nil {
			enemyColor := color.New(color.FgRed).PrintfFunc()
			enemyColor("You've encountered a %s! It attacks you for %d damage.", p.Room.Enemy.Name, p.Room.Enemy.Attack)
			p.Room.Enemy.Health -= p.Room.Enemy.Attack
			
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

	room1.Exits["north"] = room2
	room2.Exits["south"] = room1
	room2.Exits["east"] = room3
	room3.Exits["west"] = room2

	game := &Game{
		Player: &Player{Health: 100, Room: room1},
		Rooms:  []*Room{room1, room2, room3},
	}

	for game.Player.Health > 0 {
		var cmd string
		fmt.Println("Enter a direction (north, south, east, west):")
		fmt.Scanf("%s", &cmd)
		cmd = strings.ToLower(cmd)

		switch cmd {
		case "north", "south", "east", "west":
			game.Player.Move(cmd)
		default:
			fmt.Println("That is not a valid direction.")
		}

		if game.Player.HasPotion {
			gameWinColor := color.New(color.FgBlue).PrintfFunc()
			gameWinColor("You have the magic potion and made it out of the room! Congrats, you've won!")
			break
		}

		playerHealthColor := color.New(color.FgYellow).PrintfFunc()
		playerHealthColor("Your current health is: %d", game.Player.Health)
	}
}
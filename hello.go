package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Health  int
	Room    *Room
	HasItem bool
}

type Room struct {
	Description string
	Item        *Item
	Exits       map[string]*Room
}

type Game struct {
	Player *Player
	Rooms  []*Room
}

type Item struct {
	Name        string
	HealthBoost int
}

func (p *Player) Move(direction string) {
	if newRoom, ok := p.Room.Exits[direction]; ok {
		p.Room = newRoom
		fmt.Println(p.Room.Description)
		if p.Room.Item != nil {
			fmt.Printf("\nYou found a %s! Your health has increased by %d points.", p.Room.Item.Name, p.Room.Item.HealthBoost)
			p.HasItem = true
			p.Health += p.Room.Item.HealthBoost
			p.Room.Item = nil
		}
	} else {
		fmt.Println("You can't go that way.\n")
	}
}

func main() {
	room1 := &Room{Description: "\nYou're in a dark room. There's an exit to the north.", Exits: make(map[string]*Room)}
	room2 := &Room{Description: "\nYou've entered a brightly lit room. Exits are to the south and east.", Exits: make(map[string]*Room)}
	room3 := &Room{Description: "\nThis room looks ominous. There's an exit to the west.", Item: &Item{Name: "Healing Potion", HealthBoost: 50}, Exits: make(map[string]*Room)}

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
			fmt.Println("That is not a valid direction.\n")
		}

		if game.Player.HasItem {
			fmt.Println("You have the item and made it out of the room! Congrats, you've won!\n")
			break
		}

		fmt.Printf("Your current health is: %d", game.Player.Health, "\n")
	}
}
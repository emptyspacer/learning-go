package main

import (
	"fmt"
)

type creature struct {
	name      string `required maxChars="50"` // you can create tags in types using these `` back tick speech marks. These are purely visual and xan be accessed using the "reflect" module
	attack    int
	defense   int
	health    int
	abilities []string
}

func main() {
	whiteDragon := creature{
		name:    "White dragon of the north",
		attack:  1200,
		defense: 800,
		health:  9000,
		abilities: []string{
			"when below 4500 health, the white dragon starts breathing ice that slows enemies",
			"the dragon can summon ice spikes that target flying enemies",
		},
	}
	fmt.Println(whiteDragon.health, whiteDragon.attack, whiteDragon.defense)
}

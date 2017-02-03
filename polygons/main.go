package main

import "fmt"

// Polygon has names, sides, angles, and coolness factors
type Polygon struct {
	name     string
	sides    int
	angles   int
	coolness bool
}

func main() {
	polygons := createPolygons()
	for _, polygon := range polygons {
		var cool string
		if polygon.coolness == true {
			cool = "COOL!"
		} else {
			cool = "not cool..."
		}
		fmt.Println("A", polygon.name, "has", polygon.sides, "sides with", polygon.sides, polygon.angles, "degree angles.", polygon.name+"s", "are", cool)
	}
}

func createPolygons() []Polygon {
	shapes := []Polygon{}
	names := []string{"Triangle", "Square", "Pentagon", "Hexagon", "Hectagon", "Octagon"}
	for i, shape := range names {
		polygon := Polygon{name: shape, sides: i + 3, angles: 360 / (i + 3), coolness: false}
		setCoolness(&polygon)
		shapes = append(shapes, polygon)
	}
	return shapes
}

func setCoolness(p *Polygon) {
	p.coolness = p.sides%2 != 0
}

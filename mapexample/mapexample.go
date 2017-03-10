package main

import "fmt"

type Location struct {
	lat, long float64
}

func InitMap() []map[string]Location {

	var locations = make([]map[string]Location, 3)

	var l1 = make(map[string]Location)
	l1["Sunnyvale"] = Location{
		37.368, 122.0363,
	}
	l1["San Francisco"] = Location{
		37.7749, 122.4194,
	}
	locations[0] = l1

	var l2 = map[string]Location{
			"San Jose" : Location{
				37.338, 121.8863,
			},
			"Los Gatos" : Location{
				37.141, 121.5742,
			},
		}
	locations[1] = l2

	var l3 = map[string]Location {
			"Santa Clara" : {37.211, 121.589},
			"Cupertino" : {37.193, 122.231},
		}
	locations[2] = l3
	return locations
}

func main() {
	var locations = InitMap()
	for _, places := range locations {
		for place := range places {
			fmt.Printf("%s:\n\tLat:%f\n\tLong:%f\n", place, places[place].lat, places[place].long)
		}
	}
}

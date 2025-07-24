package main

import "fmt"

type st_data struct {
	roll        int
	father_name string
	mother_name string
}

func main() {
	// creating a map[string]struct

	mp := make(map[string]st_data)

	mp["yash"] = st_data{
		roll:        8,
		father_name: "rajeshGoyal",
		mother_name: "rekha goyal",
	}
	mp["Disha"] = st_data{
		roll:        8,
		father_name: "rajeshGoyal",
		mother_name: "rekha goyal",
	}

	fmt.Println(mp["yash"].father_name)
	fmt.Println(mp["Disha"].father_name)

	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

}

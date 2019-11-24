package main

import (
	"encoding/json"
	"fmt"
)

type Tree struct {
	Name  string  `json:"name"`
	Nodes []*Node `json:"nodes,omitempty"`
}

type Node struct {
	Name  string  `json:"name"`
	Nodes []*Node `json:"nodes,omitempty"`
}

func main() {
	t := Tree{
		Name: "top",
		Nodes: []*Node{
			&Node{
				Name: "a",
				Nodes: []*Node{
					&Node{
						Name: "b",
					},
				},
			},
			&Node{
				Name: "c",
				Nodes: []*Node{
					&Node{
						Name: "d",
					},
				},
			},
		},
	}

	bs, _ := json.Marshal(t)
	mbs, _ := json.MarshalIndent(t, "", "  ")
	fmt.Printf("%s\n", string(mbs))

	tt := Tree{}
	if err := json.Unmarshal(bs, &tt); err != nil {
		panic(err)
	}

	mi, _ := json.MarshalIndent(tt, "", "  ")
	fmt.Printf("%s\n", string(mi))
}

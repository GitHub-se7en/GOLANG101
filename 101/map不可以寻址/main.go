package main

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhoujielun"}}
	m2 := map[string]*Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
	m2["people"].name = "wuyanzu"
}

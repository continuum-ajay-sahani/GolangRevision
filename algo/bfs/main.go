package main

var graph = make(map[string][]string)

func main() {
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	search("you")
}

func search(name string) bool {
	var searchQueue []string
	searchQueue = append(searchQueue, graph[name]...)
	var searched []string
	for len(searchQueue) != 0 {
		person := searchQueue[0]
		searchQueue = searchQueue[1:]
		if !isPersonVisited(person, searched) {
			if isPersonMangoSeller(person) {
				println(person + " is a mango seller")
				return true
			}
			searched = append(searched, person)
			searchQueue = append(searchQueue, graph[person]...)
		}
	}
	return false
}

func isPersonVisited(person string, searched []string) bool {
	for _, v := range searched {
		if v == person {
			return true
		}
	}
	return false
}

func isPersonMangoSeller(person string) bool {
	return person[len(person)-1] == 'm'
}

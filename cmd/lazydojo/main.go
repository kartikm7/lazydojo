package main

type model struct {
	cursor   int
	items    []string
	selected map[int]struct{}
}

func initModel() model {
	return model{
		items:    []string{"dummy data", "just for", "mocking sakes"},
		selected: make(map[int]struct{}),
	}
}

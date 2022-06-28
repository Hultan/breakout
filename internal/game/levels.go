package game

type level struct {
	name   string
	bricks []string
}

var levels = []level{
	{
		"Level 1", []string{
			"111 111 111 111 111 111",
			"222 222 222 222 222 222",
			"333 333 333 333 333 333",
		},
	},
	{
		"Level 2", []string{
			"33 33 33 33 33 33 33 33",
			"222 222 222 222 222 222",
			"333 333 333 111 333 333",
		},
	},
}

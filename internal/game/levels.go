package game

type level struct {
	name   string
	bricks []string
}

var levels = []level{
	{
		"Level 1", []string{
			"000 333 333 333 333 333 333 333 000",
			"000 222 222 222 222 222 222 222 000",
			"000 222 222 222 222 222 222 222 000",
			"000 000 111 111 111 111 111 000 000",
			"000 000 000 111 111 111 000 000 000",
		},
	},
	{
		"Level 2", []string{
			"00 33 33 33 33 33 33 33 33 33 33 00",
			"000 222 222 111 111 111 222 222 000",
			"000 222 222 111 111 111 222 222 000",
			"000 000 111 111 111 111 111 000 000",
			"000 000 000 111 111 111 000 000 000",
		},
	},
}

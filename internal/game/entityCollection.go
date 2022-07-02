package game

type entityCollection []gameObject

func newEntityCollection() entityCollection {
	return make(entityCollection, 0, 10)
}

func (e *entityCollection) add(o gameObject) {
	*e = append(*e, o)
}

func (e *entityCollection) remove(o gameObject) {
	// Find pause entity
	pauseIndex := -1
	for i, ent := range *e {
		if ent == o {
			pauseIndex = i
			break
		}
	}

	// Remove pause entity
	if pauseIndex > 0 {
		(*e)[pauseIndex] = (*e)[len(*e)-1]
		*e = (*e)[:len(*e)-1]
	}
}

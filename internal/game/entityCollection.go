package game

type entityCollection []gameObject

func newEntityCollection() entityCollection {
	return make(entityCollection, 0, newEntityCollectionSize)
}

func (e *entityCollection) add(o gameObject) {
	*e = append(*e, o)
}

func (e *entityCollection) remove(o gameObject) {
	// Find index of entity o
	index := -1
	for i, ent := range *e {
		if ent == o {
			index = i
			break
		}
	}

	// Remove item at index
	e.removeIndex(index)
}

func (e *entityCollection) removeIndex(index int) {
	// Remove item at index
	(*e)[index] = (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
}

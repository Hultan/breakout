package game

type entityCollection []gameObject

func newEntityCollection() entityCollection {
	return make(entityCollection, 0, newEntityCollectionSize)
}

func (e *entityCollection) add(o gameObject) {
	*e = append(*e, o)
}

func (e *entityCollection) remove(o gameObject) {
	slice := *e

	// Find pause entity
	pauseIndex := -1
	for i, ent := range slice {
		if ent == o {
			pauseIndex = i
			break
		}
	}

	// Remove pause entity
	if pauseIndex > 0 {
		slice[pauseIndex] = slice[len(slice)-1]
		slice = slice[:len(slice)-1]
	}
}

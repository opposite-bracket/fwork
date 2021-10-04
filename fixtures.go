package fwork

// FixtureGenerator can be used as a function
// to allow data fixture customization
type FixtureGenerator func(int) interface{}

// ExceedsPageCap Checks if the page requested contains more items
// than the api is meant to handle
func ExceedsPageCap(pageNum int, pageSize int, cap int) bool {
	return (pageNum * pageSize) >= cap
}

// IndexStartingPoint evaluates the starting point of the index
// value when generating dummy lists
func IndexStartingPoint(pageNum int, pageSize int) int {
	return pageNum * pageSize
}

// IndexEndingPoint evaluates the ending point of the index
// value when generating dummy lists
func IndexEndingPoint(pageNum int, pageSize int) int {
	return pageNum*pageSize + pageSize
}

// GenerateFixtureList is a shortcut to not having to duplicate
// how to create dummy data
func GenerateFixtureList(pageNum int, pageSize int, cap int, generator FixtureGenerator) []interface{} {
	indexStartingPoint := IndexStartingPoint(pageNum, pageSize)
	indexEndingPoint := IndexEndingPoint(pageNum, pageSize)
	var items []interface{}
	for i := indexStartingPoint; i < indexEndingPoint; i++ {
		if i >= cap && items == nil {
			return make([]interface{}, 0)
		} else if i >= cap {
			return items
		}
		items = append(items, generator(i))
	}

	return items
}

package main

type disjointSetElement struct {
	size   int
	parent int
	rank   int
}

type DisjointSet struct {
	size         int
	maxTableSize int
	sets         []disjointSetElement
}

func (ds *DisjointSet) Merge(destination, source int) int {

	realDestination := ds.getParent(destination)
	realSource := ds.getParent(source)
	if realDestination != realSource {

		var merged int
		// merge two components
		// use union by rank heuristic
		if ds.sets[realDestination].rank > ds.sets[realSource].rank {
			// merge source into destination
			ds.sets[realSource].parent = realDestination
			ds.sets[realDestination].size += ds.sets[realSource].size
			ds.sets[realSource].size = 0
			merged = realDestination

		} else {
			// merge destination into source
			ds.sets[realDestination].parent = realSource
			ds.sets[realSource].size += ds.sets[realDestination].size
			ds.sets[realDestination].size = 0
			if ds.sets[realDestination].rank == ds.sets[realSource].rank {
				ds.sets[realSource].rank++
			}
			merged = realSource
		}
		// update max_table_size
		if ds.sets[merged].size > ds.maxTableSize {
			ds.maxTableSize = ds.sets[merged].size
		}
	}
	return ds.maxTableSize
}

func (ds *DisjointSet) getParent(child int) int {
	if ds.sets[child].parent != child {
		ds.sets[child].parent = ds.getParent(ds.sets[child].parent)
	}

	return ds.sets[child].parent
}

func NewDisjointSet(tables []int) *DisjointSet {

	sets := make([]disjointSetElement, len(tables))
	maxTableSize := 0
	for i, t := range tables {
		sets[i] = disjointSetElement{size: t, parent: i, rank: 0}
		if t > maxTableSize {
			maxTableSize = t
		}
	}

	return &DisjointSet{size: len(tables), maxTableSize: maxTableSize, sets: sets}
}

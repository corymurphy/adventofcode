package main

type Path struct {
	VisitedMap Map
	HeightMap  Map
	Visited    map[Vector]bool
	Ended      bool
	Success    bool
	Start      Vector
	Finish     Vector
	Current    *Vector
}

func NewPath(start Vector, height Map, visitedMap Map) *Path {
	return &Path{
		VisitedMap: visitedMap,
		Visited:    map[Vector]bool{},
		Ended:      false,
		Success:    false,
		Start:      start,
		Current:    NewCoordinates(start.X, start.Y),
	}
}

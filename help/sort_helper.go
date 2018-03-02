package help

var SortByAsc = 0
var SortByDesc = 1

type CompareContainer interface {
	Len() int
	Swap(i int, j int)
	Less(i int, j int) bool
	More(i int, j int) bool
}


func ValidSort(compareContainer CompareContainer, sortEnum int) bool {
	for i := 0; i < compareContainer.Len() - 1; i++ {
		ascResult := sortEnum == SortByAsc && compareContainer.More(i, i+1)
		descResult := sortEnum == SortByDesc && compareContainer.Less(i, i+1)
		if ascResult || descResult {
			return false
		}
	}
	return true
}

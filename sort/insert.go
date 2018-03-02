package sort

import "algorithm-go/help"

/**
插入排序
 */
func Insert(container help.CompareContainer, sortEnum int)  {
	for i := 1; i < container.Len(); i++ {
		for j := i; j > 0; j-- {
			ascResult := sortEnum == help.SortByAsc && container.Less(j, j - 1)
			descResult := sortEnum == help.SortByDesc && container.More(j, j -1)
			if ascResult || descResult {
				container.Swap(j, j -1)
			} else {
				break
			}
		}
	}
}

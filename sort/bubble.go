package sort

import "algorithm-go/help"

/**
冒泡排序
 */
func Bubble(container help.CompareContainer, sortEnum int) {
	for i := 0; i < container.Len(); i++ {
		for j := 0; j < container.Len() - i - 1; j++{
			ascResult := sortEnum == help.SortByAsc && container.More(j, j+ 1)
			descResult := sortEnum == help.SortByDesc && container.Less(j, j+ 1)
			if ascResult || descResult {
				container.Swap(j, j + 1)
			}
		}
	}
}

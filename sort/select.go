package sort

import "algorithm-go/help"

/**
选择排序
 */
func Select(container help.CompareContainer, sortEnum int) {
	for i := 0; i < container.Len(); i++ {
		selectIndex := i
		for j := i; j < container.Len(); j++ {
			ascResult := sortEnum == help.SortByAsc && container.More(selectIndex, j)
			descResult := sortEnum == help.SortByDesc && container.Less(selectIndex, j)
			if ascResult || descResult {
				selectIndex = j
			}
		}
		if selectIndex != i {
			container.Swap(selectIndex, i)
		}
	}
}

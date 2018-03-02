package sort

import "algorithm-go/help"

/**
希尔排序
 */
func Hill(container help.CompareContainer, sortEnum int) {
	h := 1
	for ; h < container.Len()/3; {
		h = h*3 + 1
	}
	for ; h >= 1; {
		for i := h; i < container.Len(); i++ {
			for j := i;
				j >= h && (sortEnum == help.SortByAsc && container.Less(j, j-h) || sortEnum == help.SortByDesc && container.More(j, j-h)); j -= h{
				container.Swap(j, j - h)
			}
		}
		h /= 3
	}
}

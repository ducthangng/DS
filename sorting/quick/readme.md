## Quick Sort

Time Complexity: O(n logn)

Quick sort is a divivde and conquer algorithm for sorting, implemented using recursive technique. Here is how it works in a nutshell:
1. Select the pivot (or where to start): it is believed the right choice from pivot the optimized pretty much the time of the sorting. However, there has not been a official way to find the pivot. Some people choose the first, last or median value of the array to be the pivot. Every entry should done the job. In this implementation I choose the last element to be the pivot.
   
2. Rules of the pivot:

	a. The left partition of the pivot is always smaller than its value.
	
	b. The right partition of the pivot is always bigger than its value.

3. Implementation: Start from the pivot, every left element which smaller than the pivot is move to the front. There should be a variables to handle the position, therefore we have the maker (M, denotes as i in the quicksort). When the iterator reach the r-1 point, we swap the b[marker + 1] <-> b[r]. In this way, the pivot is in the right position, smaller on the left, bigger on the right. We continue to quicksort those 2 partition of the pivots.   
```
quicksort(left, marker -1);
quicksort(marker+1, right);

```

When the condition (left < right) is false, we receive the full sorted array!

![img](./QuickSort2.png)



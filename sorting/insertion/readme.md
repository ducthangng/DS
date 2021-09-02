## Insertion Sort

Complexity: 
1. Best Case: O(n)
2. Worst Case: O(n^2)

Simple. For a[i], we find j satisfy a[j] >= a[i] > a[j-1], 0 <= j <= i, then push j to j+1, j+1 to j+2, ... until j is empty. After then, a[j] = tempt.
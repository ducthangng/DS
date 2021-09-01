
## Binary Search Tree

Have the binary structure (smaller on the right, bigger on the left). Adding should not have any problems. Deleting, however, is more complicated.
When deleting, we check for cases:

1. If the node is a standalone node (no children nodes) then just delete and exits.
2. If the node is not a standalone node then we do the following step:
   a. If left node is existed, then save the left node, remove the current node and replace by the left node.
   b. If right node is existed, then do the same.

Note that we should find the parent node of the removed node and the child node of the replaced node. 

Complexity: 
1. Worst case: O(n) -> number of Node.
2. Average: O(h) -> h is the height of the node.
3. Searching time: worst case: O(log2n).
4. Adding time: worst case: O(log2n).
5. Deleting time: worst case: O(log2n)
## Hash Table

In programming languages, array or slice usually denote each of their entry by non-negative integer [0 ... n]; 
Then what if we can create a data structure that resemble array but index by anything from integer, string, bytes, address? HashTable it is.

The idea behind Hash Table is that we want to generate a unique corresponding index for each entry. The complexity after indexing should be O(1) (assume that it is uniquely indexed) instead of O(n*m) with m is the average size of data. 

Hash Table consists of 2 part:
1. The Hash Function: an algorithm to hash data into a fixed format [string, int]. The final product should be a fix format F that is unique. Given that H is the function, a & b are different input, then:
```
H(a) = H(b)
```
Appeared to be quite common. For some 64 bits encrypted algo, it is about 40%. However, with long encrypted algorithm such as SHA256 32 bytes is 4.3 * 10e(-60). The longer the hashing, the more reliable.

2. The Collision Protocol: Although the collision is not common when using the right algorithm, in real world situation is can be very common, for instant 2 person can have the same name. We have 2 simple ways to deal with this:
   
   a. Using Linked List: using array of linked list to store the value indexed, if collision happen then the entry can added the value to the list.
   ![ll](./ll.png)

   b. Using Linear Probing: if 2 entry collised, then place the entry to the nearest available position to the right of the array.
   ![lp](./collision.png)


### Complexity 

Complexity of Hash Function: depend;
Complexity of Searching: O(1 + log2m) with m is the length of corresponding linked list.
## Sample Test 

---
### 1. Answer the question:

a) What is the meaning of the output of code below?

b) n = 100, what is the output?

```cpp
#include <iostream>
#include <algorithm>

using namespace std;

int n;

int main() {
    cin >> n;
    if(n % 2 == 0) cout << (n-2)/2;
    else cout << n/2; 
}
```

---

### 2. Describe a sorting algorithm that you know and how it works. What is the average time complexity of the algorithm?

---

### 3. Give the pseudocode for coin-change problem.


--- 
### 4. You are given a weighted directed graph. The vertices are enumerated from 1 to n. Your task is to find the shortest path between the vertex 1 and the vertex n.

The first line contains two integers n and m, where n is the number of vertices and m is the number of edges. Following m lines contain one edge each in form a[i], b[i] and w[i], where a[i], b[i] are edge endpoints and w[i] is the weight of the edge.

It is possible that the graph has loops and multiple edges between pair of vertices.

```
Input:

5 6
1 2 2
2 5 5
2 3 4
1 4 1
4 3 3
3 5 1

Output: 
1 4 3 5 
```

---
### 5. What is the more efficient, Red-Black Tree or Binary Search Tree? Why?


---
### 6. VGU has n students, each has a virtual locker. Each locker can contain an unlimited number of data. Trung is the person in charge of all the locker, that said every student who want to access A STUFF in his locker must goes through Trung. Imagine that if all the student cum at Trung at 1 time, Trung's gonna die.
a. Each locker is denote by the student's name. If Trung knows the locker's owner's name, Trung can find it instantly. What is the best ways for Trung to handle all the locker?
Give the demonstration for the method.

b. Given that the address of the lockers is a string that do not exceed 10 characters, give a way so that trung can mark each locker's address from the student's name.

c. What is the time complexity for Trung to search for a chunk of data in a random locker in n chunks, known that the on average a locker contain m chunks of data.
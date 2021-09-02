1.
if n even, the output is the (n-2)/2, else it returns only n/2.
b. n = 100 => output: 49

3.

```
number_of_coint -> n;
desired_result -> d;

for ( i = 0 -> n): coins[i] = input;
for ( i = 0 -> n): a[i] = d + 1;

for (i = 0 -> d) {
	for (j = 0 -> n) {
		if (i >= coins[i])
		{
			a[i] = min_value(a[i - coins[j]] + 1, a[i])
		}
	}
}

final_result = a[d];
```

4. 
```cpp

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
const int infi = 10e6;

class Dijkstra
{
private:
	int edges[100][100];
	// D has the following space [vertex],  [distance], [Previous Vertex ].
	int distance[100];
	int previous[100];
	int visited[100];

public:
	Dijkstra()
	{
		for (int i = 0; i < 100; i++)
		{
			for (int j = 0; j < 100; j++)
			{
				this->edges[i][j] = -1;
			}
			this->distance[i] = infi;
			this->previous[i] = 0;
			this->visited[i] = 0;
		}
	}

	void AddEdge(int weight, int a, int b)
	{
		this->edges[a][b] = weight;
	}

	void ComputeShortestPath(int p, int n)
	{
		this->distance[p] = 0;
		this->previous[p] = p;

		this->Compute(p, n);
	}

	void Compute(int p, int n)
	{
		vector<pair<int, int>> unvisited;
		int next_vertex = -1;
		for (int i = 1; i <= n; i++)
		{
			int min = infi;
			if (this->edges[p][i] != -1)
			{
				if (this->distance[i] > (this->distance[p] + this->edges[p][i]))
				{
					this->distance[i] = this->distance[p] + this->edges[p][i];
					this->previous[i] = p;
				}

				if (this->visited[i] != 1)
				{
					unvisited.push_back({distance[i], i});
				}
			}
		}

		this->visited[p] = 1;
		sort(unvisited.begin(), unvisited.end());

		vector<pair<int, int>>::iterator i;
		for (i = unvisited.begin(); i != unvisited.end(); i++)
		{
			this->Compute(i->second, n);
		}
	}

	void ShortTestPathTo(int x)
	{
		int count = 0;
		int a[100];
		int d[100];

		cout << "Shortest Path Value to " << x << " is: " << this->distance[x] << endl;
		while (true)
		{
			d[count] = this->distance[x];
			a[count] = x;

			if (this->previous[x] == x)
			{
				break;
			}

			x = this->previous[x];
			count++;
		}

		for (int i = count; i >= 0; i--)
		{
			cout << "vertex " << a[i] << " - distance: " << d[i] << endl;
		}
	}
};

int main()
{
	int n = 6;

	Dijkstra d;

	// Weight, vertex a, vertex b that a->b.
	d.AddEdge(5, 1, 3);
	d.AddEdge(2, 1, 2);
	d.AddEdge(2, 2, 1);
	d.AddEdge(3, 2, 3);
	d.AddEdge(2, 3, 6);
	d.AddEdge(1, 4, 6);
	d.AddEdge(1, 4, 5);
	d.AddEdge(2, 5, 4);
	d.AddEdge(8, 5, 1);
	d.AddEdge(1, 6, 2);

	d.ComputeShortestPath(5, n);
	d.ShortTestPathTo(1);
}

```

5. 
Red Black Tree is more efficient than the Binary Search Tree. The RBTree is a self-balance tree, that is, the height of tree from the first
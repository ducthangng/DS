#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

// g++ -std=c++11 kruskal.cpp
// Creating shortcut for an integer pair
typedef pair<int, int> iPair;

class Graph
{
private:
	vector<pair<int, iPair>> edges;
	int kruskal[100];
	int kruskal_weight;
	int sumNode;
	int kruskalNode;

public:
	Graph(int sumNode)
	{
		this->kruskal_weight = 0;
		this->sumNode = sumNode;
		this->kruskalNode = 0;
		for (int i = 0; i < 100; i++)
		{
			kruskal[i] = -1;
		}
	};

	void AddEdge(int weight, int node_a, int node_b)
	{
		this->edges.push_back({weight, {node_a, node_b}});
	}

	int MST()
	{
		cout << "MST by Kruskal: " << endl;

		int size = this->sumNode;
		int a[size + 10];
		sort(this->edges.begin(), this->edges.end());

		vector<pair<int, iPair>>::iterator i;
		for (i = this->edges.begin(); i != this->edges.end(); i++)
		{
			int weight = i->first;
			int edge_a = i->second.first;
			int edge_b = i->second.second;

			// No cycle so we add this to the tree
			if (!this->detectCycle(edge_a, edge_b))
			{
				// Merge cycle together
				this->MergeEgde(edge_a, edge_b);
				cout << edge_a << " " << edge_b << "  (" << weight << ")" << endl;

				this->kruskal_weight += weight;
			}

			if (this->kruskalNode == this->sumNode)
			{
				return this->kruskal_weight;
			}
		}

		return this->kruskal_weight;
	}

	// Return true as a cycle existed
	// The idea is that each connected edge inside kruskal have the same indexed.
	bool detectCycle(int egde_a, int egde_b)
	{
		if ((this->kruskal[egde_a] == -2) && (this->kruskal[egde_b] == -2))
		{
			return true;
		}

		return false;
	}

	void MergeEgde(int egde_a, int egde_b)
	{
		if (this->kruskal[egde_a] != -2)
		{
			this->kruskalNode++;
			this->kruskal[egde_a] = -2;
		}
		if (this->kruskal[egde_b] != -2)
		{
			this->kruskalNode++;
			this->kruskal[egde_b] = -2;
		}
	};
};

int main()
{
	Graph mst(7);

	mst.AddEdge(1, 1, 2);
	mst.AddEdge(2, 1, 3);
	mst.AddEdge(5, 2, 3);
	mst.AddEdge(4, 2, 4);
	mst.AddEdge(6, 4, 7);
	mst.AddEdge(2, 4, 5);
	mst.AddEdge(4, 4, 6);
	mst.AddEdge(9, 1, 7);
	mst.AddEdge(0, 5, 7);

	int w = mst.MST();

	cout << "weight is " << w << endl;
}
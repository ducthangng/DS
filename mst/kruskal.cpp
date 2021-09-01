#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

// Creating shortcut for an integer pair
typedef pair<int, int> iPair;

class Graph
{
private:
	vector<pair<int, iPair>> edges;
	vector<pair<int, int>> kruskal;
	int kruskal_weight;
	int sumNode;

public:
	Graph(int sumNode)
	{
		this->kruskal_weight = 0;
		this->sumNode = sumNode;
	};

	void AddEdge(int weight, int node_a, int node_b)
	{
		this->edges.push_back({weight, {node_a, node_b}});
	}

	int MST()
	{
		cout << "MST by Kruskal: " << endl;

		// int sum
		int a[this->edges.size() + 10];
		sort(this->edges.begin(), this->edges.end());

		vector<pair<int, iPair>>::iterator i;
		for (i = this->edges.begin(); i != this->edges.end(); i++)
		{
			int weight = i->first;
			int edge_a = i->second.first;
			int edge_b = i->second.second;

			if (this->detectCycle(edge_a, edge_b))
			{
				// No cycle so we add this to the tree
				this->kruskal.push_back({edge_a, edge_b});

				cout << edge_a << " " << edge_b << "  (" << weight << ")" << endl;

				this->kruskal_weight += weight;
			}

			if (this->kruskal.size() == this->sumNode - 1)
			{
				return this->kruskal_weight;
			}
		}

		return this->kruskal_weight;
	}

	// False for cycle exist;
	// Otherwise true;
	bool detectCycle(int egde_a, int egde_b)
	{
		int a_invole = 0;
		int b_invole = 0;

		vector<pair<int, int>>::iterator i;
		for (i = this->kruskal.begin(); i != this->kruskal.end(); i++)
		{
			int f = i->first;
			int s = i->second;

			if ((egde_a == f) || (egde_a == s))
			{
				a_invole = 1;
			}

			if ((egde_b == f) || (egde_b == s))
			{
				a_invole = 1;
			}

			if ((b_invole == 1) && (a_invole == 1))
			{
				return false;
			}
		}

		return true;
	}
};

int main()
{
	Graph mst(7);

	mst.AddEdge(1, 1, 2);
	mst.AddEdge(5, 2, 3);
	mst.AddEdge(4, 2, 4);
	mst.AddEdge(2, 4, 5);
	mst.AddEdge(4, 4, 6);
	mst.AddEdge(9, 1, 7);
	mst.AddEdge(0, 5, 7);

	int w = mst.MST();

	cout << "weight is " << w << endl;
}
#include <iostream>
#include <string>

using namespace std;

class RBT
{
private:
	int value;
	//black is false; red is true;
	bool color;
	RBT *left;
	RBT *right;
	RBT *parent;

public:
	RBT() {}

	RBT(int value, RBT *parent)
	{
		this->parent = parent;
		this->value = value;
		this->color = true;
		this->left = nullptr;
		this->right = nullptr;
	}

	RBT *search(int value)
	{
		RBT *node = this;
		while (true)
		{
			switch (this->siding(node->value - value))
			{
			case true:
				if (node->value - value == 0)
				{
					return this;
				}
				else
				{
					node = node->right;
				}

				break;
			case false:
				node = node->left;

				break;
			default:
				cout << "Error" << endl;
				return nullptr;
			}
		}
	}

	// adding according to binary search tree format. Then go through cases.
	void addNode(int value)
	{
		RBT *node = this->findRoot();
		while (true)
		{
			switch (this->siding(node->value - value))
			{
			case 1:
				if (node->right != nullptr)
				{
					node = node->right;
				}
				else
				{
					node->right = new RBT(value, node);
					rePosition(node->right);
					return;
				}

				break;

			case 0:
				if (node->left != nullptr)
				{
					node = node->left;
				}
				else
				{
					node->left = new RBT(value, node);
					rePosition(node->left);
					return;
				}

				break;

			default:
				cout << "Error" << endl;
				return;
			}
		}
	}

	// returns the root of the tree.
	void rePosition(RBT *node)
	{
		if (node->parent == nullptr)
		{
			node->color = false;
			return;
		}

		// Single block code returns itself.
		if (node->parent != nullptr)
		{
			if (!node->parent->color)
			{
				return;
			}
		}

		RBT *uncle = node->parent->parent->findOtherNode(node->parent->value);

		// father is a leaf together with the child.
		if (uncle == nullptr)
		{
			uncle->color = false;
		}

		// Single block code returns itself.
		if ((node->parent->color) && (uncle->color))
		{

			uncle->color = false;
			node->parent->color = false;
			node->rePosition(node->parent->parent);
			//	cout << uncle->value << uncle->color << endl;
			return;
		}

		if ((node->color) && (node->parent->color) && (!uncle->color))
		{
			int child_pos = node->findPosition();
			int parent_pos = node->findPosition();

			// left_left case:
			if ((child_pos == 0) && (parent_pos == 0))
			{
				node->parent->parent->rotateRight();
				RBT *sibling = node->parent->findOtherNode(node->value);
				swapColor(node->parent, sibling);
			}

			if ((child_pos == 1) && (parent_pos == 0))
			{
				node->parent->rotateLeft();
				node->parent->rotateRight();
				swapColor(node, node->right);
			}

			if ((child_pos == 1) && (parent_pos == 1))
			{
				node->parent->parent->rotateLeft();
				RBT *sibling = node->parent->findOtherNode(node->value);
				swapColor(node->parent, sibling);
			}

			if ((child_pos == 0) && (parent_pos == 1))
			{
				node->parent->rotateRight();
				node->parent->rotateLeft();
				swapColor(node, node->left);
			}
		}

		return;
	}

	// this should be perform by the right node;
	void rotateRight()
	{
		RBT *left_node = this->left;

		this->left = this->left->right;

		left_node->parent = this->parent;

		left_node->right = this;
	}

	// this should be perform by the right node;
	void rotateLeft()
	{
		RBT *right_node = this->right;

		this->right = this->right->left;

		right_node->parent = this->parent;

		right_node->left = this;

		bool color = right_node->color;
		right_node->color = this->color;
		this->color = color;
	}

	void swapColor(RBT *node_1, RBT *node_2)
	{
		bool color = node_1->color;
		node_1->color = this->color;
		this->color = color;
	}

	// RBT is a parent and value is a child. Returns the sibling of value Node;
	RBT *findOtherNode(int value)
	{
		if (this->left != nullptr)
		{
			if (this->left->value == value)
			{
				return this->right;
			}
		}

		if (this->right != nullptr)
		{
			if (this->right->value == value)
			{
				return this->left;
			}
		}

		return nullptr;
	}

	// findPosition finds the position of a child to its parent position;
	// return 0 if the the node is on left; 1 if it is on right;
	// return 2 if no requirement meets.
	int findPosition()
	{
		if (this->parent != nullptr)
		{
			if (this->parent->left != nullptr)
			{
				if (this->parent->left->value == this->value)
				{
					return 0;
				}
			}

			if (this->parent->right != nullptr)
			{
				if (this->parent->right->value == this->value)
				{
					return 1;
				}
			}
		}

		return 2;
	}

	int siding(int value)
	{
		if (value >= 0)
		{
			return 1;
		}

		return 0;
	}

	RBT *findRoot()
	{
		RBT *node = this;
		while (node->parent != nullptr)
		{
			node = node->parent;
		}

		return node;
	}

	void deleteNode(int value) {}

	void deleteX(int value) {}

	void print(const string prefix = "", bool isBlack = false)
	{
		cout << prefix << (isBlack ? "R--" : "B--");
		cout << this->value << endl;

		if (this->left != NULL)
			this->left->print(prefix + (isBlack ? "|   " : "    "), this->left->color);
		if (this->right != NULL)
			this->right->print(prefix + (isBlack ? "|   " : "    "), this->right->color);
	}
};

int main()
{
	RBT root(10, nullptr);
	root.rePosition(&root);

	root.addNode(5);
	root.addNode(12);
	root.addNode(4);
	root.addNode(3);

	root.print();
}
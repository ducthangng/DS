#include <iostream>
#include <string>

using namespace std;

class RBT
{
private:
	int value;
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
		//black is false; red is true;
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
			case false:
				if (node->value - value == 0)
				{
					return this;
				}
				else
				{
					node = node->right;
				}

				break;
			case true:
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
			case false:
				if (node->right->value != -1)
				{
					node = node->right;
				}
				else
				{
					node->right->value = value;
					node->right->right = new RBT(-1, node->right);
					node->right->left = new RBT(-1, node->right);
					rePosition(node->right);
					return;
				}

				break;

			case true:
				if (node->left->value != -1)
				{
					node = node->left;
				}
				else
				{
					node->left->value = value;
					node->left->right = new RBT(-1, node->left);
					node->left->left = new RBT(-1, node->left);
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
		// root node;
		if (node->parent == nullptr)
		{
			node->color = false;

			if (this->left == nullptr)
			{
				this->left = new RBT(-1, this);
			}

			if (this->right == nullptr)
			{
				this->right = new RBT(-1, this);
			}

			return;
		}

		if (!node->parent->color)
		{
			return;
		}

		RBT *uncle = node->parent->parent->findOtherNode(node->parent->value);

		// Nullptr is black.
		if (uncle->value == -1)
		{
			uncle->color = false;
		}

		if ((node->parent->color) && (uncle->color))
		{
			uncle->color = false;
			node->parent->color = false;
			node->parent->parent->color = true;
			node->rePosition(node->parent->parent);
			return;
		}

		if ((node->color) && (node->parent->color) && (!uncle->color))
		{
			int child_pos = node->findPosition();
			int parent_pos = node->parent->findPosition();

			cout << child_pos << " x " << parent_pos << endl;
			// left_left case:
			if ((child_pos == 0) && (parent_pos == 0))
			{
				node->parent->parent->rotateRight(node->parent);
				swapColor(node->parent, node->parent->right);
			}

			// left_right case:
			if ((child_pos == 1) && (parent_pos == 0))
			{
				node->parent->rotateLeft(node);

				node->parent->rotateRight(node);
				swapColor(node, node->right);
			}

			// right_right case:
			if ((child_pos == 1) && (parent_pos == 1))
			{
				node->parent->parent->rotateLeft(node->parent);
				swapColor(node->parent, node->parent->left);
			}

			// left_right case:
			if ((child_pos == 0) && (parent_pos == 1))
			{
				node->parent->rotateRight(node);
				node->parent->rotateLeft(node);
				swapColor(node, node->left);
			}
		}

		return;
	}

	// this should be perform by the right node;
	void rotateRight(RBT *p)
	{

		p->parent = this->parent;
		if (this->parent != nullptr)
		{
			if (!this->siding(this->parent->value - this->value))
			{
				this->parent->right = p;
			}
			else
			{
				this->parent->left = p;
			}
		}

		this->left = p->right;
		p->right->parent = this;

		this->parent = p;
		p->right = this;
	}

	// "this" is the node that being called. P is the replaced node;
	void rotateLeft(RBT *p)
	{
		p->parent = this->parent;
		if (!this->siding(this->parent->value - this->value))
		{
			this->parent->right = p;
		}
		else
		{
			this->parent->left = p;
		}

		this->right = p->left;
		p->left->parent = this;

		this->parent = p;
		p->left = this;
	}

	void swapColor(RBT *node_1, RBT *node_2)
	{
		bool color = node_1->color;
		node_1->color = node_2->color;
		node_2->color = color;
	}

	// RBT is a parent and value is a child. Returns the sibling of arg value;
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

	// siding tells us if the node is on the right or the left of the parents.
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

	void print()
	{
		RBT *r = this->findRoot();
		r->prints();
	}

	void prints(const string prefix = "", bool isBlack = false)
	{
		if (this->value == -1)
		{
			return;
		}

		cout << prefix << (isBlack ? "R--" : "B--");
		cout << this->value << endl;

		if (this->left != NULL)
			this->left->prints(prefix + (isBlack ? "|   " : "    "), this->left->color);
		if (this->right != NULL)
			this->right->prints(prefix + (isBlack ? "|   " : "    "), this->right->color);
	}
};

int main()
{
	RBT root(20, nullptr);
	root.rePosition(&root);

	root.addNode(6);
	root.addNode(8);
	root.addNode(7);
	root.addNode(14);
	root.addNode(5);
	root.addNode(13);
	root.addNode(19);

	root.print();
}
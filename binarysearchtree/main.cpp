#include <iostream>
#include <algorithm>

using namespace std;

class Node
{
private:
	int value;
	Node *left;
	Node *right;

public:
	Node() {}

	Node(int value)
	{
		this->value = value;
		this->left = nullptr;
		this->right = nullptr;
	}

	Node *getNode()
	{
		return this;
	}

	int getValue()
	{
		return this->value;
	}

	Node *search(int value)
	{
		Node *pointer = this->getNode();
		while (true)
		{
			if (pointer->getValue() > value)
			{
				pointer = pointer->left;
				continue;
			}

			if (pointer->getValue() < value)
			{
				pointer = pointer->right;
				continue;
			}

			if (pointer->getValue() == value)
			{
				return pointer;
			}

			if ((pointer->right == nullptr) && (pointer->left == nullptr))
			{
				return nullptr;
			}
		}
	}

	void add(Node *node)
	{
		Node *pointer = this->getNode();
		while (true)
		{
			if (node->getValue() >= pointer->getValue())
			{

				if (pointer->right != nullptr)
				{
					pointer = pointer->right;
				}
				else
				{
					pointer->right = node;
					return;
				}
			}
			else
			{
				if (pointer->left != nullptr)
				{
					pointer = pointer->left;
				}
				else
				{
					pointer->left = node;
					return;
				};
			}
		}
	}

	void remove(int value)
	{
		this->deleteX(this->search(value));
	}

	void deleteX(Node *node)
	{
		if ((node->left == nullptr) && (node->right == nullptr))
		{

			Node *t = this->findPredecessor(node->getValue());
			cout << "precessor" << t->value << endl;
			if (t->left->getValue() == node->getValue())
			{
				t->left = nullptr;
			}
			else
			{
				t->right = nullptr;
			}
		}
		else if ((node->left == nullptr) || (node->right == nullptr))
		{
			if (node->right == nullptr)
			{
				int value = node->left->getValue();
				node->left = node->left->left;
				node->value = value;
			}
			else
			{
				int value = node->right->getValue();
				node->right = node->right->right;
				node->value = value;
			}
		}
		else
		{
			Node *t = this->findSuccessor(node);
			int value = t->getValue();
			cout << value << endl;
			this->remove(t->getValue());
			node->value = value;
		}
	}

	Node *findPredecessor(int value)
	{

		Node *pointer = this->getNode();
		Node *pre_node = this->getNode();
		while (true)
		{
			if (pointer->getValue() > value)
			{
				pre_node = pointer;
				pointer = pointer->left;
				continue;
			}

			if (pointer->getValue() < value)
			{
				pre_node = pointer;
				pointer = pointer->right;
				continue;
			}

			if (pointer->getValue() == value)
			{
				return pre_node;
			}

			if ((pointer->right == nullptr) && (pointer->left == nullptr))
			{
				return nullptr;
			}
		}
	}

	// Find the right-most of the left branch.
	Node *findSuccessor(Node *node)
	{
		node = node->left;
		while (true)
		{
			if (node->right != nullptr)
			{
				node = node->right;
			}
			else
			{
				return node;
			}
		}
	}

	void printTree(const string prefix = "", bool isLeft = false)
	{
		cout << prefix << (isLeft ? "L--" : "R--");
		cout << this->getValue() << endl;

		if (this->left != NULL)
			this->left->printTree(prefix + (isLeft ? "|   " : "    "), true);
		if (this->right != NULL)
			this->right->printTree(prefix + (isLeft ? "|   " : "    "), false);
	}
};

int main()
{
	Node root(10);

	Node a(5);
	Node b(5);
	Node c(7);
	Node d(12);
	Node e(11);
	Node f(19);

	root.add(&b);
	root.add(&a);
	root.add(&c);
	root.add(&d);
	root.add(&e);
	root.add(&f);
	// root.add(4);
	// root.add(0);

	root.remove(12);
	root.remove(11);

	root.printTree("", false);
}

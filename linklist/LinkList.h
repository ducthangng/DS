#include <iostream>
#include <algorithm>
#include <string.h>

using namespace std;

class Node
{
private:
	string value;
	Node *after;
	Node *before;

public:
	Node(){};

	Node(string value)
	{
		this->value = value;
		this->after = NULL;
		this->before = NULL;
	}

	void addAfter(Node *node)
	{
		this->after = node;
	}

	void addBefore(Node *node)
	{
		this->before = node;
	}

	void removeAfter()
	{
		this->after = NULL;
	}

	void removeBefore()
	{
		this->before = NULL;
	}

	Node *afterNode()
	{
		return this->after;
	}

	Node *beforeNode()
	{
		return this->before;
	}

	string getValue()
	{
		return this->value;
	}

	void print()
	{
		cout << "current value: " << this->value << endl;
	}
};

class LinkList
{
private:
	Node *firstNode;

public:
	LinkList()
	{
		this->firstNode = NULL;
	}

	void InsertLast(Node *node)
	{
		if (this->firstNode == NULL)
		{
			this->firstNode = node;
			return;
		}

		Node *current_node = this->firstNode;
		while (current_node->afterNode() != NULL)
		{
			current_node = current_node->afterNode();
		}

		current_node->addAfter(node);
	}

	void InsertBefore(Node *node)
	{
		if (this->firstNode == NULL)
		{
			this->firstNode = node;
			return;
		}

		node->addAfter(this->firstNode);
		this->firstNode = node;
	}

	void Remove(string keyword)
	{
		if (this->firstNode == NULL)
		{
			cout << "No current node in list" << endl;
			return;
		}

		Node *current = this->firstNode;
		while (current != NULL)
		{
			string value = current->getValue();
			if (value == keyword)
			{
				Node *afterNode = current->afterNode();
				Node *beforeNode = current->beforeNode();

				if (afterNode == NULL)
				{
					beforeNode->removeAfter();
				}
				else
				{
					afterNode->addBefore(beforeNode);
					beforeNode->addAfter(afterNode);
				}

				return;
			}

			current = current->afterNode();
		}
	}

	Node *search(string keyword)
	{
		Node *current = this->firstNode;
		while (current != NULL)
		{
			if (current->getValue() == keyword)
			{
				return current;
			}
			current = current->afterNode();
		}

		return NULL;
	}

	void print()
	{
		Node *current = this->firstNode;
		int count = 0;
		while (current != NULL)
		{
			string value = current->getValue();
			cout << "Node :" << count << " value: " << value << endl;
			current = current->afterNode();
			count += 1;
		}
	}
};
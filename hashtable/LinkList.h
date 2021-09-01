#include <iostream>
#include <algorithm>
#include <string.h>

using namespace std;

class Node
{
private:
	string key;
	string value;
	Node *after;
	Node *before;

public:
	Node(){};

	Node(string key, string value)
	{
		this->key = key;
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

	string getKey()
	{
		return this->key;
	}

	void print()
	{
		cout << "key - value: " << this->key << "-" << this->value << endl;
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

	string firstKey()
	{
		if (this->firstNode == nullptr)
		{
			return "";
		}

		return this->firstNode->getKey();
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
				if ((current->afterNode() == nullptr) && (current->beforeNode() != nullptr))
				{
					current->beforeNode()->removeAfter();
					return;
				}

				if (current->beforeNode() == nullptr)
				{
					this->firstNode = current->afterNode();
					return;
				}

				current->beforeNode()->addAfter(current->afterNode());
				return;
			}

			current = current->afterNode();
		}
	}

	string search(string keyword)
	{
		int count = 0;

		Node *current = this->firstNode;
		while (current != NULL)
		{
			if (current->getValue() == keyword)
			{
				return current->getValue();
			}
			current = current->afterNode();
		}

		return "";
	}

	void print()
	{
		Node *current = this->firstNode;
		int count = 0;
		while (current != NULL)
		{
			current->print();
			current = current->afterNode();
			count += 1;
		}
	}
};
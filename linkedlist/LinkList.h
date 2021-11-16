#include <iostream>
#include <algorithm>

using namespace std;

class LinkedList
{
private:
	LinkedList *next;
	LinkedList *before;
	string value;

public:
	LinkedList(string value)
	{
		this->next = nullptr;
		this->before = nullptr;
		this->value = value;
	}

	string getValue()
	{
		return this->value;
	}

	void push(string value)
	{
		this->value = value;
		return;
	}

	void push_back(string value)
	{
		if (this->next != nullptr)
		{
			this->next->push_back(value);
			return;
		}

		this->next = new LinkedList(value);
		this->next->before = this;
	}

	void push_front(string value)
	{
		if (this->before != nullptr)
		{
			this->before->push_front(value);
			return;
		}

		this->before = new LinkedList(value);
		this->before->next = this;
	}

	LinkedList *search(string value)
	{
		bool top = this->search2top(value);
		bool bot = this->search2bot(value);

		if ((top == true) || (bot == true))
		{
			return this;
		}

		return nullptr;
	}

	LinkedList *search2top(string value)
	{
		if (this->value != value)
		{
			return this->next == nullptr ? nullptr : this->next->search2top(value);
		}
		return this;
	}

	LinkedList *search2bot(string value)
	{
		if (this->value != value)
		{
			return this->next == nullptr ? nullptr : this->next->search2bot(value);
		}
		return this;
	}

	LinkedList *searchtop()
	{
		if (this->before != nullptr)
		{
			return this->before->searchtop();
		}

		return this;
	}

	void remove(string value)
	{
		LinkedList *currentNode = this->search(value);
		if (currentNode != nullptr)
		{
			currentNode->before->next = currentNode->next;
			currentNode->next->before = currentNode->before;
			delete currentNode;
		}
	}

	void print()
	{
		LinkedList *node = this->searchtop();

		cout << node->getValue() << " ";
		while (node->next != nullptr)
		{
			node = node->next;
			cout << node->getValue() << " ";
		}

		cout << endl;
	}
};
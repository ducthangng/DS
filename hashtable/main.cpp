#include <iostream>
#include <string>
#include <algorithm>
#include "LinkList.h"

#define A 54059	  /* a prime */
#define B 76963	  /* another prime */
#define C 101	  /* yet another prime */
#define FIRSTH 37 /* also prime */

using namespace std;

// Collision handler: linkedlist - open addressing.
// Rules: same key - same value is accepted. Push back into linkedlist or move up depends on the next value on open addressing.
// Insert-last cost: O(m). Insert-first cost: O(1).
// Retrive cost: O(1). Hashtable.list[(key.hashed.string)] = value. (linkedlist)

// Chaining Hashtable

class HashTable
{
private:
	LinkedList *list[110];

public:
	HashTable()
	{
		for (int i = 0; i < 101; i++)
		{
			this->list[i] = new LinkedList("");
		}
	}

	// Hash Function: convert string to int.
	unsigned hash_str(string key)
	{
		unsigned h = FIRSTH;
		for (int i = 0; i < key.length(); i++)
		{
			h = (h * A) ^ (key[i] * B);
		}
		return h % C;
	}

	void addEntry(string key, string value)
	{

		unsigned hashed_key = this->hash_str(key);

		string cvalue = this->list[hashed_key]->getValue();
		if (cvalue == "")
		{
			this->list[hashed_key]->push(value);
			return;
		}

		this->list[hashed_key]->push_back(value);
		return;
	}

	void remove(string key, string value)
	{
		unsigned hashed_key = this->hash_str(key);
		this->list[hashed_key]->remove(value);
	}

	void search(string key, string value)
	{
		unsigned hashed_key = this->hash_str(key);
		LinkedList *ll = this->list[hashed_key]->search(key);
		if (ll == nullptr)
		{
			cout << "no record found" << endl;
		}

		this->list[hashed_key]->print();
	}

	void print(string key)
	{
		unsigned hashed_key = this->hash_str(key);
		cout << key << " :";
		this->list[hashed_key]->print();
		cout << endl;
	}
};

class HashTablev2
{
private:
	string key[100];
	string val[100];

public:
	HashTablev2()
	{
		for (int i = 0; i < 100; i++)
		{
			key[i] = "";
			val[i] = "";
		}
	}

	// Hash Function: convert string to int.
	int hash_str(string key)
	{
		unsigned h = FIRSTH;
		for (int i = 0; i < key.length(); i++)
		{
			h = (h * A) ^ (key[i] * B);
		}
		return h % 23;
	}

	void addEntry(string key, string val)
	{
		int hashed_key = this->hash_str(key);

		if (this->key[hashed_key] == "")
		{
			this->key[hashed_key] = key;
			this->val[hashed_key] = val;
			return;
		}

		int count = hashed_key;
		while (this->key[count] != "")
		{
			count++;
		}

		this->key[count] = key;
		this->val[count] = val;
		return;
	}

	void search(string key)
	{
		int hashed_key = this->hash_str(key);

		if (this->key[hashed_key] == key)
		{
			cout << "key - value: " << key << ":" << val[hashed_key] << endl;
			return;
		}

		int count = hashed_key;
		while (this->key[count] != key)
		{
			count++;
		}

		cout << "key - value: " << key[count] << ":" << val[count] << endl;
	}

	void remove(string key)
	{
		int hashed_key = this->hash_str(key);
		int count = hashed_key;

		if (this->key[count] != key)
		{
			while (this->key[count] != key)
			{
				count++;
			}
		}

		this->key[count] = "";
		this->val[count] = "";
	}

	void print()
	{
		for (int i = 0; i < 100; i++)
		{
			if (this->key[i] != "")
			{
				cout << "key - value: " << key[i] << ":" << val[i] << endl;
			}
		}
	}
};

int main()
{
	// HashTable ht;
	// ht.addEntry("hfekasnlaskeffefeaew", "read");
	// ht.addEntry("hfekasnlaskeffefeaew", "write");
	// ht.addEntry("ducthang", "execute");
	// ht.addEntry("ducthang", "write");
	// ht.addEntry("ironman", "read");
	// ht.addEntry("ironman", "share");
	// ht.print("hfekasnlaskeffefeaew");
	// ht.print("ducthang");
	// ht.print("ironman");

	HashTablev2 hashv2;

	hashv2.addEntry("ducthang", "12");
	hashv2.addEntry("ducthang", "13");

	hashv2.search("ducthang");
	hashv2.remove("ducthang");

	hashv2.print();
}
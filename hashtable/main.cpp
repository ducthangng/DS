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
	LinkList list[110];
	int number_of_entry;

public:
	HashTable()
	{
		this->number_of_entry = 0;
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

	void addEntry(Node entry)
	{
		unsigned hashed_key = this->hash_str(entry.getKey());
		this->list[hashed_key].InsertLast(&entry);
	}

	void remove(Node entry)
	{
		unsigned hashed_key = this->hash_str(entry.getKey());
		this->list[hashed_key].Remove(entry.getValue());
	}

	void search(string key)
	{
		unsigned hashed_key = this->hash_str(key);

		string value = this->list[hashed_key].search(key);

		if (value.length() == 0)
		{
			cout << "no entries found" << endl;
		}

		this->list[hashed_key].print();
	}

	void print()
	{
		for (int i = 0; i < 110; i++)
		{
			if ((this->list[i].firstKey().length()) != 0)
			{
				unsigned hashed_key = this->hash_str(list[i].firstKey());
				cout << "hashed_key: " << hashed_key << endl;
				this->list[i].print();
			}
		}
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
	// HashTablev1 or HashTable does not have dynamic creation. It is prepered to used the hashtable v2.
	// Node thang("ducthang", "20");
	// Node replication_1("ducthang", "20");
	// Node replication_2("ducthang", "19");
	// Node replication_3("ducthang_3", "20");
	// Node replication_4("linhdan_1", "12");
	// Node replication_5("god", "100");
	// Node replication_6("bless", "12121");

	// HashTable hash;

	// hash.addEntry(thang);
	// hash.addEntry(replication_1);
	// hash.addEntry(replication_2);
	// hash.addEntry(replication_3);
	// hash.addEntry(replication_4);
	// hash.addEntry(replication_5);
	// hash.addEntry(replication_6);

	// // hash.search("ducthang");
	// hash.print();

	// cout << "delete ducthang" << endl;
	// hash.remove(thang);
	// hash.remove(thang);
	// hash.print();

	HashTablev2 hashv2;

	hashv2.addEntry("ducthang", "12");
	hashv2.addEntry("ducthang", "13");

	hashv2.search("ducthang");
	hashv2.remove("ducthang");

	hashv2.print();
}
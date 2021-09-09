#include <iostream>
#include "LinkList.h"

using namespace std;

int main()
{
	LinkedList list("nguyen");
	list.push_back("duc");
	list.push_back("thang");
	list.push_front("my name is");

	list.print();
}
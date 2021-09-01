#include <iostream>
#include "LinkList.h"

using namespace std;

int main()
{
	Node a("ducthang0");
	Node b("ducthang1");
	Node c("ducthang2");
	Node d("ducthang3");
	Node e("ducthang4");
	Node f("ducthang5");
	Node g("ducthang6");
	Node h("ducthang7");
	Node k("ducthang8");
	Node l("ducthang9");
	Node m("ducthang10");

	LinkList list;
	list.InsertLast(&a);
	list.InsertLast(&b);
	list.InsertLast(&c);
	list.InsertLast(&d);
	list.InsertLast(&e);
	list.InsertLast(&f);
	list.InsertLast(&g);
	list.InsertLast(&h);

	Node *result = list.search("ducthang4");
	result->print();
}
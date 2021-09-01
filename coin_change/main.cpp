#include <iostream>
#include <algorithm>

using namespace std;

int minvalue(int a, int b)
{
	if (a < b)
	{
		return a;
	}
	else
	{
		return b;
	}
}

int main()
{
	// the result 2D array
	int a[100];
	int coins[100];

	// number of coins and the desired result
	int n;
	int d;
	cin >> n >> d;

	for (int i = 0; i < n; i++)
	{
		cin >> coins[i];
	}

	a[0] = 0;
	for (int i = 1; i <= d; i++)
	{
		a[i] = d + 1;
	}

	for (int i = 0; i <= d; i++)
	{
		for (int j = 0; j < n; j++)
		{
			if (i >= coins[j])
			{
				a[i] = minvalue(a[i - coins[j]] + 1, a[i]);
			}
		}
	}

	for (int i = 0; i <= d; i++)
	{
		cout << a[i] << endl;
	}

	cout << "result: " << a[d] << endl;
}
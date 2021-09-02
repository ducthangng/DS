#include <iostream>
using namespace std;

void sort(int a[100], int r)
{
	for (int i = 1; i < r; i++)
	{
		int j = i - 1;
		while (true)
		{
			if (a[i] > a[j])
			{
				break;
			}

			if ((a[j] > a[i]) && (j != 0))
			{
				j = j - 1;
			}

			if ((a[i] >= a[j]) || ((a[j] >= a[i]) && (j == 0)))
			{
				if ((a[j] >= a[i]) && (j == 0))
				{
					j = -1;
				}
				int tempt = a[i];

				// shift everything to the right;
				for (int k = i; k > j + 1; k--)
				{
					a[k] = a[k - 1];
				}

				a[j + 1] = tempt;

				for (int i = 0; i < r; i++)
				{
					cout << a[i] << " ";
				}
				cout << endl;

				break;
			}
		}
	}

	// for (int i = 0; i < r; i++)
	// {
	// 	cout << a[i] << " ";
	// }
	// cout << endl;

	return;
}

int main()
{
	int n;
	int a[100];

	cin >> n;
	for (int i = 0; i < n; i++)
	{
		cin >> a[i];
	}

	sort(a, n);
}
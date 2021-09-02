#include <iostream>
#include <algorithm>
#include <math.h>

using namespace std;

int n;
float k, t, c;

int main()
{
	cin >> n >> k;
	for (int i = 1; i <= n; i++)
	{
		cin >> c;
		t += c;
	}
	c = n;
	if (t / n == k)
		cout << 0;
	else
	{
		while (round(t / n) != k)
		{
			n++;
			t += k;
			if (round(t / n) == k)
				break;
		}
		cout << n - c << endl;
	}
}

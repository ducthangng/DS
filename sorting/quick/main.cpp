#include <iostream>

using namespace std;

int n;
char a[100];
int b[100];

void quicksort(int l, int r)
{
    if (l <= r - 1)
    {
        // pivot is always r
        int i = l - 1;
        for (int k = l; k <= r - 1; k++)
        {
            if (a[k] <= a[r])
            {
                i++;
                char t = a[i];
                a[i] = a[k];
                a[k] = t;
            }
        }

        i++;
        char t = a[r];
        a[r] = a[i];
        a[i] = t;

        quicksort(l, i - 1);
        quicksort(i + 1, r);
    }
}

void quicksortb(int l, int r)
{
    if (l < r)
    {
        // pivot is always r
        int i = l - 1;
        for (int k = l; k <= r - 1; k++)
        {
            if (b[k] <= b[r])
            {
                i++;
                int t = b[i];
                b[i] = b[k];
                b[k] = t;
            }
        }

        i++;
        int t = b[r];
        b[r] = b[i];
        b[i] = t;

        for (int i = 0; i < n; i++)
        {
            cout << b[i] << " ";
        }
        cout << endl;
        quicksortb(l, i - 1);
        quicksortb(i + 1, r);
    }
}

int main()
{
    cout << "Choose Number, Press 1. Choose Letter, Press 2." << endl;
    cin >> n;

    if (n == 2)
    {
        cin >> n;
        for (int i = 0; i < n; i++)
        {
            cin >> a[i];
        }

        quicksort(0, n - 1);
        for (int i = 0; i < n; i++)
        {
            cout << a[i] << " ";
        }
        cout << endl;
    }

    if (n == 1)
    {
        cin >> n;
        for (int i = 0; i < n; i++)
        {
            cin >> b[i];
        }

        quicksortb(0, n - 1);
        for (int i = 0; i < n; i++)
        {
            cout << b[i] << " ";
        }
        cout << endl;
    }
}
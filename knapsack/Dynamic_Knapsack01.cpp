#include <iostream>

using namespace std;

int max(int a, int b)
{
    return (a > b) ? a : b;
}

void Knapsack01(int capacity, int weight[], int value[], int size)
{
    int i, j;
    int K[size + 1][capacity + 1]; //Create a table that store the value
    for (int i = 0; i <= size; i++)
    {
        for (int j = 0; j <= capacity; j++)
        {
            if (i == 0 || j == 0) //Base case
                K[i][j] = 0;
            else if (weight[i - 1] <= j)
                K[i][j] = max(value[i - 1] + K[i - 1][j - weight[i - 1]], K[i - 1][j]);
            else
                K[i][j] = K[i - 1][j];
        }
    }

    //Print out the result and the item's weight
    int res = K[size][capacity];
    cout << endl
         << res;
    j = capacity;
    for (int i = size; i > 0 && res > 0; i--)
    {
        if (res == K[i - 1][j])
            continue;
        else
        {
            cout << endl
                 << weight[i - 1];
            res -= value[i - 1];
            j -= weight[i - 1];
        }
    }
    cout << endl;

    //Print out the table
    for (int i = 1; i <= size; i++)
    {
        for (int j = 0; j <= capacity; j++)
        {
            cout << K[i][j] << " ";
        }
        cout << endl;
    }
}

int main()
{
    int weight[] = {6, 3, 5, 4, 6};
    int value[] = {6, 5, 4, 2, 2};
    int capacity = 10;
    int size = sizeof(value) / sizeof(value[0]);
    cout << "Optimal value coresponds to the capacity = ";
    Knapsack01(capacity, weight, value, size);
    return 0;
}
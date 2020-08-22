#include <stdio.h>

/**
 * brute force, time complexity: O(n^2), space complexity: O(1)
 */
int maxProfit(int *prices, int pricesSize)
{
  int maxVal = 0;
  if (prices == NULL || pricesSize <= 1)
  {
    return maxVal;
  }

  for (int i = 0; i < pricesSize; i++)
  {
    for (int j = i + 1; j < pricesSize; j++)
    {
      int minus = prices[j] - prices[i];
      if (minus > maxVal)
      {
        maxVal = minus;
      }
    }
  }
  return maxVal;
}

/**
 * two pointers, time complexity: O(n), space complexity: O(1)
 */
int maxProfit2(int *prices, int pricesSize)
{
  int maxVal = 0;
  if (prices == NULL || pricesSize <= 1)
  {
    return maxVal;
  }

  int minVal = prices[0];
  for (int i = 1; i < pricesSize; i++)
  {
    int minus = prices[i] - minVal;
    if (minus > maxVal)
    {
      maxVal = minus;
    }

    if (prices[i] < minVal)
    {
      minVal = prices[i];
    }
  }

  return maxVal;
}

/**
 * two pointers, time complexity: O(n), space complexity: O(1) ===> enhance maxProfit2
 */
int maxProfit3(int *prices, int pricesSize)
{
  int maxVal = 0;
  if (prices == NULL || pricesSize <= 1)
  {
    return maxVal;
  }

  int minVal = prices[0];
  for (int i = 1; i < pricesSize; i++)
  {
    if (prices[i] < minVal)
    {
      minVal = prices[i];
      continue;
    }

    int minus = prices[i] - minVal;
    if (minus > maxVal)
    {
      maxVal = minus;
    }
  }

  return maxVal;
}

int main()
{
  int a[] = {7, 1, 5, 3, 6, 4};
  printf("%d\n", maxProfit(a, (int)(sizeof(a) / sizeof(a[0]))));
  printf("%d\n", maxProfit2(a, (int)(sizeof(a) / sizeof(a[0]))));
  printf("%d\n", maxProfit3(a, (int)(sizeof(a) / sizeof(a[0]))));

  int b[] = {7, 6, 4, 3, 1};
  printf("%d\n", maxProfit(b, (int)(sizeof(b) / sizeof(b[0]))));
  printf("%d\n", maxProfit2(b, (int)(sizeof(b) / sizeof(b[0]))));
  printf("%d\n", maxProfit3(a, (int)(sizeof(a) / sizeof(a[0]))));
}
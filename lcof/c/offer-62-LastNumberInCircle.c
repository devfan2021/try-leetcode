#include <stdio.h>

/**
 * n = 1: f(n,m) = 0
 * n > 1: f(n, m) = (f(n-1, m) + m) % n
 */
int lastRemaining(int n, int m)
{
  if (n < 1 || m < 1)
  {
    return -1;
  }

  int last = 0;
  for (int i = 2; i <= n; i++)
  {
    last = (last + m) % i;
  }
  return last;
}

int main()
{
  printf("%d\n", lastRemaining(5, 3));   // 3
  printf("%d\n", lastRemaining(10, 17)); // 2
  printf("%d\n", lastRemaining(8, 6));
  return 0;
}
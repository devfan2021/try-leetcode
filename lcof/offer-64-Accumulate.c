#include <stdio.h>

int sumNums(int n)
{
  return n * (n + 1) / 2;
}

int sumNums1(int n)
{
  return n == 0 ? 0 : n + sumNums1(n - 1);
}

// use && shortcircuit
int sumNums2(int n)
{
  n && (n += sumNums2(n - 1));
  return n;
}

int main()
{
  printf("%d\n", sumNums(3));  // 6
  printf("%d\n", sumNums1(3)); // 6
  printf("%d\n", sumNums2(3)); // 6
  printf("%d\n", sumNums(9));  // 9
  printf("%d\n", sumNums1(9)); // 9
  printf("%d\n", sumNums2(9)); // 9
  return 0;
}
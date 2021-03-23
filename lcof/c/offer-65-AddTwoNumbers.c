#include <stdio.h>

/**
 * ^: 0^0=0, 1^1=0, 0^1=1, 1^0=1
 * &: 0&0=0, 1&0=0, 0&1=0, 1&1=1
 */
int add(int a, int b)
{
  int sum = 0, carry = 0;
  do
  {
    sum = a ^ b;
    carry = ((unsigned int)(a & b) << 1);

    a = sum;
    b = carry;

  } while (carry != 0);
  return a;
}

int main()
{
  printf("==%d\n", add(0, 0));
  printf("==%d\n", add(0, 1));
  printf("==%d\n", add(1, 1));
  printf("==%d\n", add(1, 0));
  printf("==%d\n", add(2, 3));
  return 0;
}
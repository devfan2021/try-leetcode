#include <stdio.h>
#include "minunit.h"

int hammingWeight(uint32_t n)
{
  int count = 0;
  for (; n != 0; n >>= 1)
  {
    if ((n & 1) == 1)
    {
      count++;
    }
  }
  return count;
}

int hammingWeight1(uint32_t n)
{
  int count = 0;
  while (n)
  {
    n = n & (n - 1);
    count++;
  }
  return count;
}

MU_TEST(test_case)
{
  mu_check(5 == 7);
}

MU_TEST_SUITE(test_suite)
{
  MU_RUN_TEST(test_case);
}

int main()
{
  MU_RUN_SUITE(test_suite);
  MU_REPORT();
  return MU_EXIT_CODE;
}
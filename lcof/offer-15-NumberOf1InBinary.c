#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/

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
  mu_assert_int_eq(3, hammingWeight(00000000000000000000000000001011));
  mu_assert_int_eq(3, hammingWeight1(00000000000000000000000000001011));
  mu_assert_int_eq(1, hammingWeight(00000000000000000000000010000000));
  mu_assert_int_eq(1, hammingWeight1(00000000000000000000000010000000));
  // mu_assert_int_eq(31, hammingWeight(11111111111111111111111111111101));
  // mu_assert_int_eq(31, hammingWeight1(11111111111111111111111111111101));
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
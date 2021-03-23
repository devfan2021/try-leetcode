#include <stdio.h>
#include "minunit.h"

int helper(int n, int *nums, int numsSize);

int fib(int n)
{
  if (n <= 1)
  {
    return n;
  }

  return (fib(n - 2) + fib(n - 1)) % 1000000007;
}

int fib1(int n)
{
  int cache[101] = {-1};
  cache[0] = 0;
  cache[1] = 1;
  return helper(n, cache, (int)sizeof(cache) / sizeof(cache[0]));
}

int helper(int n, int *nums, int numsSize)
{
  if (n > numsSize)
  {
    return 0;
  }

  if (n <= 1)
  {
    return n;
  }

  int val = nums[n];
  if (val > 0)
  {
    return val;
  }

  int retVal = (helper(n - 2, nums, numsSize) + helper(n - 1, nums, numsSize)) % 1000000007;
  nums[n] = retVal;
  return retVal;
}

int fib2(int n)
{
  if (n <= 1)
  {
    return n;
  }

  int dp0 = 0, dp1 = 1;
  int retVal = 0;
  for (int i = 2; i <= n; i++)
  {
    retVal = (dp0 + dp1) % 1000000007;
    dp0 = dp1;
    dp1 = retVal;
  }
  return retVal;
}

MU_TEST(test_case)
{
  mu_assert_int_eq(1, fib(2));
  mu_assert_int_eq(5, fib(5));
  mu_assert_int_eq(55, fib(10));
  mu_assert_int_eq(102334155, fib(40));
  mu_assert_int_eq(586268941, fib(50));

  mu_assert_int_eq(1, fib1(2));
  mu_assert_int_eq(5, fib1(5));
  mu_assert_int_eq(55, fib1(10));
  mu_assert_int_eq(102334155, fib1(40));
  mu_assert_int_eq(586268941, fib1(50));

  mu_assert_int_eq(1, fib2(2));
  mu_assert_int_eq(5, fib2(5));
  mu_assert_int_eq(55, fib2(10));
  mu_assert_int_eq(102334155, fib2(40));
  mu_assert_int_eq(586268941, fib2(50));
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
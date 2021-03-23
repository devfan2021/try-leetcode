#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

int max(int val1, int val2);

int cuttingRope(int n)
{
  if (n < 2)
  {
    return 0;
  }
  else if (n == 2)
  {
    return 1;
  }
  else if (n == 3)
  {
    return 2;
  }

  int *dp = (int *)malloc(sizeof(int) * (n + 1));
  dp[0] = 0;
  dp[1] = 0;
  dp[2] = 1;
  dp[3] = 2;

  for (int i = 4; i <= n; i++)
  {
    int curMax = 0;
    for (int j = 1; j <= i / 2; j++)
    {
      curMax = max(curMax, max(j * (i - j), j * dp[i - j]));
    }
    dp[i] = curMax;
  }

  return dp[n];
}

int max(int val1, int val2)
{
  if (val1 > val2)
  {
    return val1;
  }
  return val2;
}

MU_TEST(test_case)
{
  mu_assert_int_eq(1, cuttingRope(2));
  mu_assert_int_eq(36, cuttingRope(10));
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
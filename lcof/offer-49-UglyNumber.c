#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

#define min(a, b) a > b ? b : a

int nthUglyNumber(int n)
{
  int *dp = malloc(sizeof(int) * n);
  int twoStep = 0, threeStep = 0, fiveStep = 0;
  dp[0] = 1;
  for (int i = 1; i < n; i++)
  {
    int twoVal = dp[twoStep] * 2;
    int threeVal = dp[threeStep] * 3;
    int fiveVal = dp[fiveStep] * 5;
    int minVal = min(min(twoVal, threeVal), fiveVal);

    if (minVal == twoVal)
      twoStep++;
    if (minVal == threeVal)
      threeStep++;
    if (minVal == fiveVal)
      fiveStep++;

    dp[i] = minVal;
  }

  return dp[n - 1];
}

MU_TEST(test_case)
{
  mu_assert_int_eq(12, nthUglyNumber(10));
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
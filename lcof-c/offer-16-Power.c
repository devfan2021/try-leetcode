#include <stdio.h>
#include "minunit.h"

double myPow(double x, int n)
{
  if (n == 0)
  {
    return 1;
  }

  if (n < 0)
  {
    return myPow(1 / x, -n);
  }

  return (n % 2 == 0) ? myPow(x * x, n / 2) : x * myPow(x * x, n / 2);
}

double myPow1(double x, int n)
{
  if (n == 0)
  {
    return 1;
  }

  double half = myPow1(x, n / 2);
  if (n % 2 == 0)
  {
    return half * half;
  }

  if (n > 0)
  {
    return half * half * x;
  }
  else
  {
    return half * half / x;
  }
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
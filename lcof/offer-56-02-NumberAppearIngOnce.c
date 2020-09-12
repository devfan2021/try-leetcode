#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/

int singleNumber(int *nums, int numsSize)
{
  int result = 0;
  for (int i = 0; i < 32; i++)
  {
    int tmp = 0; //记录所有数字第i位的和
    for (int i = 0; i < numsSize; i++)
    {
      tmp += (nums[i] >> i) & 1; //求第i位之和
    }

    if (tmp % 3)
    {
      result += 1 << i; //置1
    }
  }

  return result;
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
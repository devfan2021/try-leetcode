#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// 摩尔投票法

int majorityElement(int *nums, int numsSize)
{
  if (nums == NULL || numsSize == 0)
  {
    return 0;
  }

  int retVal = 0, vote = 0;
  for (int i = 0; i < numsSize; i++)
  {
    if (vote == 0)
    {
      retVal = nums[i];
    }

    vote += (retVal == nums[i]) ? 1 : -1;
  }

  // 验证众数
  int count = 0;
  for (int i = 0; i < numsSize; i++)
  {
    if (nums[i] == retVal)
    {
      count++;
    }
  }

  return count > numsSize / 2 ? retVal : 0;
}

MU_TEST(test_case)
{
  int a[] = {1, 2, 3, 2, 2, 2, 5, 4, 2};
  mu_assert_int_eq(2, majorityElement(a, (int)sizeof(a) / sizeof(a[0])));
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
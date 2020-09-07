#include <stdio.h>
#include "minunit.h"

/**
 * https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
 */
int max(int val1, int val2)
{
  if (val1 > val2)
  {
    return val1;
  }
  return val2;
}

int maxSubArray(int *nums, int numsSize)
{
  if (nums == NULL || numsSize == 0)
  {
    return 0;
  }

  int previewSum = 0;
  int maxSum = nums[0];
  for (int i = 0; i < numsSize; i++)
  {
    previewSum = max(nums[i], previewSum + nums[i]);
    maxSum = max(previewSum, maxSum);
  }

  return maxSum;
}

int maxSubArray1(int *nums, int numsSize)
{
  if (nums == NULL || numsSize == 0)
  {
    return 0;
  }

  int previewSum = 0;
  int maxSum = nums[0];
  for (int i = 0; i < numsSize; i++)
  {
    if (previewSum + nums[i] < nums[i])
    {
      previewSum = nums[i];
    }
    else
    {
      previewSum = previewSum + nums[i];
    }

    if (maxSum < previewSum)
    {
      maxSum = previewSum;
    }
  }

  return maxSum;
}

MU_TEST(test_case)
{
  int a[10] = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
  mu_assert_int_eq(6, maxSubArray(a, (int)sizeof(a) / sizeof(a[0])));
  mu_assert_int_eq(6, maxSubArray1(a, (int)sizeof(a) / sizeof(a[0])));
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
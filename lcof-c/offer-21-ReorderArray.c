#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *exchange(int *nums, int numsSize, int *returnSize)
{
  if (nums == NULL || numsSize == 0)
  {
    *returnSize = 0;
    return NULL;
  }

  int begin = 0, end = numsSize - 1;
  while (begin < end)
  {
    if (nums[begin] % 2 != 0)
    {
      begin++;
      continue;
    }

    if (nums[end] % 2 == 0)
    {
      end--;
      continue;
    }
    else
    {
      int tmp = nums[begin];
      nums[begin] = nums[end];
      nums[end] = tmp;
      begin++;
    }
  }

  *returnSize = numsSize;
  return nums;
}

MU_TEST(test_case)
{
  int a[] = {1, 2, 3, 4};
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
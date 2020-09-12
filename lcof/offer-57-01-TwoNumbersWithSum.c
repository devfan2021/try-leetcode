#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *twoSum(int *nums, int numsSize, int target, int *returnSize)
{
  if (nums == NULL || numsSize == 0)
  {
    *returnSize = 0;
    return NULL;
  }

  int begin = 0, end = numsSize - 1;
  while (begin < end)
  {
    if (target - nums[begin] == nums[end])
    {
      *returnSize = 2;
      int *retVals = malloc(sizeof(int) * 2);
      retVals[0] = nums[begin];
      retVals[1] = nums[end];
      return retVals;
    }
    else if (target - nums[begin] > nums[end])
    {
      begin++;
    }
    else
    {
      end--;
    }
  }

  *returnSize = 0;
  return NULL;
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
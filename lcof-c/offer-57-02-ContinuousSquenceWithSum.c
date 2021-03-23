#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int **findContinuousSequence(int target, int *returnSize, int **returnColumnSizes)
{
  if (target <= 0)
  {
    *returnSize = 0;
    *returnColumnSizes = 0;
    return NULL;
  }

  int **output = (int **)malloc(sizeof(int *) * target);
  *returnColumnSizes = (int *)malloc(sizeof(int) * target);
  *returnSize = 0;

  int left = 1, right = 1, sum = 0, len = target / 2;
  while (left <= len)
  {
    if (sum < target)
    {
      sum += right;
      right++;
    }
    else if (sum > target)
    {
      sum -= left;
      left++;
    }
    else
    {
      int cols = right - left;
      output[*returnSize] = malloc(sizeof(int) * cols);
      (*returnColumnSizes)[*returnSize] = cols;
      for (int i = left; i < right; i++)
      {
        output[*returnSize][i - left] = i;
      }
      (*returnSize)++;
      sum -= left;
      left++;
    }
  }

  return output;
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
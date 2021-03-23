#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

int comp(const void *a, const void *b)
{
  return *(int *)a - *(int *)b;
}

int *getLeastNumbers(int *arr, int arrSize, int k, int *returnSize)
{
  if (arr == NULL || arrSize == 0 || k == 0)
  {
    *returnSize = 0;
    return NULL;
  }

  qsort(arr, arrSize, sizeof(int), comp);
  int *output = (int *)malloc(sizeof(int) * k);
  for (int i = 0; i < k; i++)
  {
    output[i] = arr[i];
  }
  *returnSize = k;
  return output;
}

int partition(int *a, int left, int right)
{
  int pivot = a[left];
  while (left < right)
  {
    while (left < right && a[right] >= pivot)
    {
      right--;
    }

    a[left] = a[right];

    while (left < right && a[left] <= pivot)
    {
      left++;
    }
    a[right] = a[left];
  }
  a[left] = pivot;
  return left;
}

int *getLeastNumbers2(int *arr, int arrSize, int k, int *returnSize)
{
  if (arr == NULL || arrSize == 0 || k == 0)
  {
    *returnSize = 0;
    return NULL;
  }

  int left = 0, right = arrSize - 1, pivotPos = partition(arr, left, right);
  while (k != (pivotPos + 1))
  {
    if (k > (pivotPos + 1))
    {
      left = pivotPos + 1;
    }
    else
    {
      right = pivotPos - 1;
    }
    pivotPos = partition(arr, left, right);
  }

  *returnSize = k;
  return arr;
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
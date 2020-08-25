#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// question similar offer-03-01, don't modify input array
int getDuplication(int *nums, int numsSize)
{
  if (nums == NULL || numsSize == 0)
  {
    return -1;
  }

  int *hash = (int *)malloc(sizeof(int) * numsSize);
  for (int i = 0; i < numsSize; i++)
  {
    if (hash[nums[i]] == 1)
    {
      return nums[i];
    }
    else
    {
      hash[nums[i]] = 1;
    }
  }

  return -1;
}

int countRange(int *nums, int numsSize, int begin, int end)
{
  if (nums == NULL || numsSize == 0)
  {
    return 0;
  }

  int count = 0;
  for (int i = 0; i < numsSize; i++)
  {
    if (nums[i] >= begin && nums[i] <= end)
    {
      count++;
    }
  }
  return count;
}

// binary search
int getDuplication1(int *nums, int numsSize)
{
  if (nums == NULL || numsSize == 0)
  {
    return -1;
  }

  int begin = 1, end = numsSize - 1;
  while (begin <= end)
  {
    int mid = begin + (end - begin) / 2;
    int count = countRange(nums, numsSize, begin, end);
    if (end == begin)
    {
      if (count > 1)
      {
        return begin;
      }
      else
      {
        break;
      }
    }

    if (count > (mid - begin + 1))
    {
      end = mid;
    }
    else
    {
      begin = mid + 1;
    }
  }

  return -1;
}

MU_TEST(test_case)
{
  int a[] = {2, 3, 5, 4, 3, 2, 6, 7};
  mu_assert_int_eq(3, getDuplication(a, (int)sizeof(a) / sizeof(a[0])));
  mu_assert_int_eq(2, getDuplication1(a, (int)sizeof(a) / sizeof(a[0])));
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
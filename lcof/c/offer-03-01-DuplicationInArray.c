#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

int cmp(const void *_a, const void *_b)
{
  int *a = (int *)_a;
  int *b = (int *)_b;
  return *a - *b;
}

int findRepeatNumber(int *nums, int numsSize)
{
  qsort(nums, numsSize, sizeof(int), cmp);
  for (int i = 0; i < numsSize - 1; i++)
  {
    if (nums[i] == nums[i + 1])
    {
      return nums[i];
    }
  }
  return -1;
}

// put i at index i of nums
int findRepeatNumber1(int *nums, int numsSize)
{
  for (int i = 0; i < numsSize; i++)
  {
    while (i != nums[i])
    {
      if (nums[i] == nums[nums[i]])
      {
        return nums[i];
      }

      int tmp = nums[nums[i]];
      nums[nums[i]] = nums[i];
      nums[i] = tmp;
    }
  }

  return -1;
}

int findRepeatNumber2(int *nums, int numsSize)
{
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

MU_TEST(test_check)
{
  int a[] = {1, 2, 3, 4, 2};
  mu_assert_int_eq(2, findRepeatNumber(a, (int)(sizeof(a) / sizeof(a[0]))));
  mu_assert_int_eq(2, findRepeatNumber1(a, (int)(sizeof(a) / sizeof(a[0]))));
  mu_assert_int_eq(2, findRepeatNumber2(a, (int)(sizeof(a) / sizeof(a[0]))));

  int b[] = {1, 2, 3, 4, 6, 3};
  mu_assert_int_eq(3, findRepeatNumber(b, (int)(sizeof(b) / sizeof(b[0]))));
  mu_assert_int_eq(3, findRepeatNumber1(b, (int)(sizeof(b) / sizeof(b[0]))));
  mu_assert_int_eq(3, findRepeatNumber2(b, (int)(sizeof(b) / sizeof(b[0]))));
}

MU_TEST_SUITE(test_suite)
{
  MU_RUN_TEST(test_check);
}

int main()
{
  MU_RUN_SUITE(test_suite);
  MU_REPORT();
  return MU_EXIT_CODE;
}
#include <stdio.h>
#include <stdlib.h>

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
    if (i != nums[i])
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

int main()
{
  int a[] = {1, 2, 3, 4, 2};
  printf("%d", findRepeatNumber(a, (int)(sizeof(a) / sizeof(a[0]))));
  printf("%d", findRepeatNumber1(a, (int)(sizeof(a) / sizeof(a[0]))));
  printf("%d", findRepeatNumber2(a, (int)(sizeof(a) / sizeof(a[0]))));

  int b[] = {1, 2, 3, 4, 6, 3};
  printf("%d", findRepeatNumber(b, (int)(sizeof(b) / sizeof(b[0]))));
  printf("%d", findRepeatNumber1(b, (int)(sizeof(b) / sizeof(b[0]))));
  printf("%d", findRepeatNumber2(a, (int)(sizeof(a) / sizeof(a[0]))));

  return 0;
}
#include <stdio.h>

int getNumberSameAsIndex(int *nums, int numsSize)
{
  if (nums == NULL || numsSize <= 0)
  {
    return -1;
  }

  for (int i = 0; i < numsSize; i++)
  {
    if (nums[i] == i)
    {
      return i;
    }
  }

  return -1;
}

int getNumberSameAsIndex1(int *nums, int numsSize)
{
  if (nums == NULL || numsSize <= 0)
  {
    return -1;
  }

  int begin = 0, end = numsSize - 1;
  while (begin <= end)
  {
    int mid = begin + (end - begin) / 2;
    if (nums[mid] == mid)
    {
      return mid;
    }
    else if (nums[mid] > mid)
    {
      end = mid - 1;
    }
    else
    {
      begin = mid + 1;
    }
  }

  return -1;
}

int main()
{
  int a[] = {-3, -1, 1, 3, 5};
  printf("==%d\n", getNumberSameAsIndex(a, (int)(sizeof(a) / sizeof(a[0]))));
  printf("==%d\n", getNumberSameAsIndex1(a, (int)(sizeof(a) / sizeof(a[0]))));

  int b[] = {-3, -1, 1, 2, 4, 5};
  printf("==%d\n", getNumberSameAsIndex(b, (int)(sizeof(b) / sizeof(b[0]))));
  printf("==%d\n", getNumberSameAsIndex1(b, (int)(sizeof(b) / sizeof(b[0]))));

  int c[] = {0, 1, 3, 5, 6};
  printf("==%d\n", getNumberSameAsIndex(c, (int)(sizeof(c) / sizeof(c[0]))));
  printf("==%d\n", getNumberSameAsIndex1(c, (int)(sizeof(c) / sizeof(c[0]))));

  int d[] = {0};
  printf("==%d\n", getNumberSameAsIndex(d, (int)(sizeof(d) / sizeof(d[0]))));
  printf("==%d\n", getNumberSameAsIndex1(d, (int)(sizeof(d) / sizeof(d[0]))));

  int e[] = {10};
  printf("==%d\n", getNumberSameAsIndex(e, (int)(sizeof(e) / sizeof(e[0]))));
  printf("==%d\n", getNumberSameAsIndex1(e, (int)(sizeof(e) / sizeof(e[0]))));

  return 0;
}
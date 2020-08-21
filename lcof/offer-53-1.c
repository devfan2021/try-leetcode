#include <stdio.h>

int search(int *nums, int numsSize, int target)
{
  if (nums == NULL)
  {
    return 0;
  }

  int begin = -1, end = -1;
  for (int i = 0; i < numsSize; i++)
  {
    if (nums[i] == target && begin == -1)
    {
      begin = i;
    }
    else if (nums[i] != target && begin >= 0)
    {
      end = i;
      break;
    }
    else if (i == numsSize - 1)
    {
      end = i + 1;
    }
  }
  if (begin == -1)
  {
    return 0;
  }
  else
  {
    return end - begin;
  }
}

int main()
{
  int a[] = {1, 2, 3};
  printf("%d\n", search(a, (int)(sizeof(a) / sizeof(a[0])), 2));

  int b[] = {1, 2, 3, 3, 3, 4, 5};
  printf("%d\n", search(b, (int)(sizeof(b) / sizeof(b[0])), 3));

  int c[] = {1, 2, 3, 3, 3, 3, 3};
  printf("%d\n", search(c, (int)(sizeof(c) / sizeof(c[0])), 3));

  int d[] = {1, 2, 3, 3, 3, 3, 3};
  printf("%d\n", search(d, (int)(sizeof(d) / sizeof(d[0])), 10));
  return 0;
}
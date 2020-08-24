#include <stdio.h>

// using two pointers
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
  else if (begin >= 0 && end <= 0)
  {
    return 1;
  }
  else
  {
    return end - begin;
  }
}

// simple solution, concise
int search1(int *nums, int numsSize, int target)
{
  int cnt = 0;
  for (int i = 0; i < numsSize; i++)
  {
    cnt = (nums[i] == target ? cnt + 1 : cnt);
  }
  return cnt;
}

int main()
{
  int a[] = {1, 2, 3};
  printf("%d\n", search(a, (int)(sizeof(a) / sizeof(a[0])), 2));
  printf("%d\n", search1(a, (int)(sizeof(a) / sizeof(a[0])), 2));

  int b[] = {1, 2, 3, 3, 3, 4, 5};
  printf("%d\n", search(b, (int)(sizeof(b) / sizeof(b[0])), 3));
  printf("%d\n", search1(b, (int)(sizeof(b) / sizeof(b[0])), 3));

  int c[] = {1, 2, 3, 3, 3, 3, 3};
  printf("%d\n", search(c, (int)(sizeof(c) / sizeof(c[0])), 3));
  printf("%d\n", search1(c, (int)(sizeof(c) / sizeof(c[0])), 3));

  int d[] = {1, 2, 3, 3, 3, 3, 3};
  printf("%d\n", search(d, (int)(sizeof(d) / sizeof(d[0])), 10));
  printf("%d\n", search1(d, (int)(sizeof(d) / sizeof(d[0])), 10));

  int f[] = {1};
  printf("%d\n", search(f, (int)(sizeof(f) / sizeof(f[0])), 1));
  printf("%d\n", search1(f, (int)(sizeof(f) / sizeof(f[0])), 1));

  return 0;
}
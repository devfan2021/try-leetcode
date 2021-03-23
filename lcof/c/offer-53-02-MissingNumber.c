#include <stdio.h>

// iterator, time complexity:O(n)
int missingNumber(int *nums, int numsSize)
{
  if (nums == NULL || numsSize == 0)
  {
    return 0;
  }

  int min = nums[0];
  for (int i = 1; i < numsSize; i++)
  {
    if (min + i != nums[i])
    {
      return min + i;
    }
  }

  // check miss first element or last element
  if (min == 0)
  {
    return nums[numsSize - 1] + 1;
  }
  else
  {
    return min - 1;
  }
}

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/solution/mian-shi-ti-53-ii-0n-1zhong-que-shi-de-shu-zi-er-f/
int missingNumber1(int *nums, int numsSize)
{
  if (nums == NULL || numsSize == 0)
  {
    return 0;
  }

  int begin = 0, end = numsSize - 1;
  while (begin <= end)
  {
    int mid = begin + (end - begin) / 2;
    if (nums[mid] == mid) // key point
    {
      begin = mid + 1;
    }
    else
    {
      end = mid - 1;
    }
  }
  return begin;
}

//b^b=0 => a^b^b=a
int missingNumber2(int *nums, int numsSize)
{
  int val = 0;
  for (int i = 0; i < numsSize; i++)
  {
    val ^= nums[i] ^ i;
  }
  return val ^ numsSize;
}

int main()
{
  int a[] = {0, 1, 2};
  printf("==%d\n", missingNumber(a, (int)(sizeof(a) / sizeof(a[0]))));
  printf("==%d\n", missingNumber1(a, (int)(sizeof(a) / sizeof(a[0]))));
  printf("==%d\n", missingNumber2(a, (int)(sizeof(a) / sizeof(a[0]))));

  int b[] = {1, 2};
  printf("==%d\n", missingNumber(b, (int)(sizeof(b) / sizeof(b[0]))));
  printf("==%d\n", missingNumber1(b, (int)(sizeof(b) / sizeof(b[0]))));
  printf("==%d\n", missingNumber2(b, (int)(sizeof(b) / sizeof(b[0]))));

  int c[] = {0, 1, 2, 3, 4, 5, 6, 7, 9};
  printf("==%d\n", missingNumber(c, (int)(sizeof(c) / sizeof(c[0]))));
  printf("==%d\n", missingNumber1(c, (int)(sizeof(c) / sizeof(c[0]))));
  printf("==%d\n", missingNumber2(c, (int)(sizeof(c) / sizeof(c[0]))));

  int d[] = {0, 1, 2, 3, 4};
  printf("==%d\n", missingNumber(d, (int)(sizeof(d) / sizeof(d[0]))));
  printf("==%d\n", missingNumber1(d, (int)(sizeof(d) / sizeof(d[0]))));
  printf("==%d\n", missingNumber2(d, (int)(sizeof(d) / sizeof(d[0]))));

  int e[] = {};
  printf("==%d\n", missingNumber(e, (int)(sizeof(e) / sizeof(e[0]))));
  printf("==%d\n", missingNumber1(e, (int)(sizeof(e) / sizeof(e[0]))));
  printf("==%d\n", missingNumber2(e, (int)(sizeof(e) / sizeof(e[0]))));

  return 0;
}
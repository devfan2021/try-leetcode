#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int compare(const void *arg1, const void *arg2)
{
  return *(int *)arg1 - *(int *)arg2;
}

/**
 * check max and min diff less 4
 */
bool isStraight(int *nums, int numsSize)
{
  if (nums == NULL || numsSize < 1)
  {
    return false;
  }

  bool numFlags[14] = {false};
  int min = 14;
  int max = 0;
  for (int i = 0; i < numsSize; i++)
  {
    int curVal = nums[i];
    if (curVal == 0)
    {
      continue;
    }

    if (numFlags[curVal])
    {
      return false;
    }
    numFlags[curVal] = true;
    if (min > curVal)
    {
      min = curVal;
    }
    if (curVal > max)
    {
      max = curVal;
    }
  }

  if (max - min <= 4)
  {
    return true;
  }

  return false;
}

bool isStraight1(int *nums, int numsSize)
{
  if (nums == NULL || numsSize < 1)
  {
    return false;
  }

  qsort(nums, numsSize, sizeof(int), compare);

  int numberOfZero = 0;
  int numberOfGap = 0;

  //caculate zero number
  for (int i = 0; i < numsSize && nums[i] == 0; i++)
  {
    numberOfZero++;
  }

  // caculate number gap, compare each two element
  int begin = numberOfZero;
  while (begin < numsSize - 1)
  {
    // have repeat number, return false
    if (nums[begin] == nums[begin + 1])
    {
      return false;
    }

    numberOfGap += nums[begin + 1] - nums[begin] - 1;
    begin++;
  }
  return (numberOfGap > numberOfZero) ? false : true;
}

int main()
{
  int a[] = {1, 2, 3, 4, 5};
  printf("%d\n", isStraight(a, (int)(sizeof(a) / sizeof(a[0]))));  // true
  printf("%d\n", isStraight1(a, (int)(sizeof(a) / sizeof(a[0])))); // true
  int b[] = {0, 0, 1, 2, 5};
  printf("%d\n", isStraight(b, (int)(sizeof(b) / sizeof(b[0]))));  // true
  printf("%d\n", isStraight1(b, (int)(sizeof(b) / sizeof(b[0])))); // true
  return 0;
}
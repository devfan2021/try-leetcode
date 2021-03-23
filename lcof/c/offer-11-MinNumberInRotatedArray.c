#include <stdio.h>
#include "minunit.h"

int minArray(int *numbers, int numbersSize)
{
  if (numbers == NULL || numbersSize == 0)
  {
    return -1;
  }

  int minVal = numbers[0];

  for (int i = 1; i < numbersSize; i++)
  {
    if (numbers[i] < minVal)
    {
      minVal = numbers[i];
    }
  }
  return minVal;
}

int minArray1(int *numbers, int numbersSize)
{
  if (numbers == NULL || numbersSize == 0)
  {
    return -1;
  }

  int right = numbersSize - 1;
  while (right > 0)
  {
    if (numbers[right - 1] <= numbers[right])
    {
      right--;
    }
    else
    {
      break;
    }
  }
  return numbers[right];
}

int minArray2(int *numbers, int numbersSize)
{
  if (numbers == NULL || numbersSize == 0)
  {
    return -1;
  }

  int low = 0, high = numbersSize - 1;
  while (low < high)
  {
    int mid = low + (high - low) / 2;
    if (numbers[mid] < numbers[high])
    {
      high = mid;
    }
    else if (numbers[mid] > numbers[high])
    {
      low = mid + 1;
    }
    else
    {
      high = high - 1;
    }
  }
  return numbers[low];
}

MU_TEST(test_case)
{
  int a[] = {3, 4, 5, 1, 2};
  mu_assert_int_eq(1, minArray(a, (int)sizeof(a) / sizeof(a[0])));
  mu_assert_int_eq(1, minArray1(a, (int)sizeof(a) / sizeof(a[0])));
  mu_assert_int_eq(1, minArray2(a, (int)sizeof(a) / sizeof(a[0])));

  int b[] = {2, 2, 2, 0, 1};
  mu_assert_int_eq(0, minArray(b, (int)sizeof(b) / sizeof(b[0])));
  mu_assert_int_eq(0, minArray1(b, (int)sizeof(b) / sizeof(b[0])));
  mu_assert_int_eq(0, minArray2(b, (int)sizeof(b) / sizeof(b[0])));

  int c[] = {3, 1};
  mu_assert_int_eq(1, minArray(c, (int)sizeof(c) / sizeof(c[0])));
  mu_assert_int_eq(1, minArray1(c, (int)sizeof(c) / sizeof(c[0])));
  mu_assert_int_eq(1, minArray2(c, (int)sizeof(c) / sizeof(c[0])));

  int d[] = {3, 1, 3};
  mu_assert_int_eq(1, minArray(d, (int)sizeof(d) / sizeof(d[0])));
  mu_assert_int_eq(1, minArray1(d, (int)sizeof(d) / sizeof(d[0])));
  mu_assert_int_eq(1, minArray2(d, (int)sizeof(d) / sizeof(d[0])));

  int f[] = {1, 3, 3};
  mu_assert_int_eq(1, minArray(f, (int)sizeof(f) / sizeof(f[0])));
  mu_assert_int_eq(1, minArray1(f, (int)sizeof(f) / sizeof(f[0])));
  mu_assert_int_eq(1, minArray2(f, (int)sizeof(f) / sizeof(f[0])));
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
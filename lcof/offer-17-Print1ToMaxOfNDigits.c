#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *printNumbers(int n, int *returnSize)
{
  if (n == 0)
  {
    *returnSize = 0;
    return NULL;
  }

  int length = 0;
  while (n--)
  {
    length = 10 * length + 9;
  }

  int *retVal = (int *)malloc(sizeof(int) * length);
  for (int i = 1; i <= length; i++)
  {
    retVal[i - 1] = i;
  }

  *returnSize = length;
  return retVal;
}

int incrementOne(char *number)
{
  int isOverflow = 0;
  int nTakeOver = 0;
  int nLength = strlen(number);
  int i;
  for (i = nLength; i >= 0; i--)
  {
    int nSum = number[i] - '0' + nTakeOver;

    // add one in last
    if (i == nLength - 1)
    {
      nSum++;
    }

    if (nSum >= 10)
    {
      if (i == 0)
      {
        isOverflow = 1;
      }
      else
      {
        nSum -= 10;
        nTakeOver = 1;
        number[i] = '0' + nSum;
      }
    }
    else
    {
      number[i] = '0' + nSum;
      break;
    }
  }

  return isOverflow;
}

void printNumber(char *number)
{
  bool isBeginning0 = true;
  int nLength = strlen(number);
  int i;
  for (i = 0; i < nLength; i++)
  {
    // avoid begin zero, such as: 078, 003
    if (isBeginning0 && number[i] != '0')
      isBeginning0 = false;

    if (!isBeginning0)
      printf("%c", number[i]);
  }
  printf("\t");
}

// 处理大数问题，字符串
void print1ToMaxOfNDigitsOne(int n)
{
  if (n < 0)
  {
    return;
  }

  char number[n + 1];
  memset(number, '0', n);
  number[n] = '\0';

  while (!incrementOne(number))
  {
    printNumber(number);
  }
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
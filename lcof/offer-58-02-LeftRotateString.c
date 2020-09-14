#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

//https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

char *reverseLeftWords(char *s, int n)
{
  if (s == NULL || n == 0)
  {
    return s;
  }

  int len = strlen(s), i = 0;
  char *output = (char *)malloc(sizeof(char) * (len + 1));
  while (i < len)
  {
    *(output++) = s[(n + i) % len];
    i++;
  }
  *output = '\0';

  // 指向数组起始位置
  return output - len;
}

char *reverseLeftWords1(char *s, int n)
{
  if (s == NULL || n == 0)
  {
    return s;
  }

  int m = n;
  char *output = (char *)calloc(strlen(s) + 1, sizeof(char));

  char *b = output;
  char *t = s;
  while (n--)
    t++;
  while (*t)
    *b++ = *t++; // 拷贝第一段
  while (m--)
    *b++ = *s++; // 拷贝第二段

  return output;
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
#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/

void reverse(char *pBegin, char *pEnd)
{
  if (pBegin == NULL || pEnd == NULL)
    return;

  while (pBegin < pEnd)
  {
    char tmp = *pBegin;
    *pBegin = *pEnd;
    *pEnd = tmp;

    pBegin++;
    pEnd++;
  }
}

// 前后空格有问题
char *reverseWords(char *s)
{
  if (s == NULL)
    return NULL;

  char *pBegin = s;
  char *pEnd = s;
  while (*pEnd != '\0')
    pEnd++;

  pEnd--;

  //反转整个句子
  reverse(pBegin, pEnd);

  //反转句子中的每个单词
  pBegin = s;
  pEnd = s;
  while (*pBegin != '\0')
  {
    if (*pBegin == ' ')
    {
      pBegin++;
      pEnd++;
    }
    else if (*pEnd == ' ' || *pEnd == '\0')
    {
      reverse(pBegin, --pEnd);
      pBegin = ++pEnd;
    }
    else
    {
      pEnd++;
    }
  }

  return s;
}

char *reverseWords1(char *s)
{
  // 去掉尾部的空格。最终至少留下一个空格，除非本身长度就为0
  int n = strlen(s);
  while (n > 0 && s[n - 1] == ' ')
    n--;

  // 去掉头部空格
  int front = 0;
  if (n > 0)
  {
    while (s[front] == ' ')
      front++;
  }

  if (n - front == 0)
    return "";

  char *output = (char *)calloc(n - front + 1, sizeof(char));
  int index = 0; // 新字符串下标

  // 最字符末尾进行遍历
  for (int begin = n - 1, end = n - 1; begin >= front; begin--)
  {
    if (begin == front || (s[begin] == ' ' && begin != end))
    {
      int tmpIndex = begin + 1;
      if (begin == front) // 头部，不用加一
        tmpIndex = begin;

      for (; tmpIndex <= end; tmpIndex++, index++)
      {
        output[index] = s[tmpIndex];
      }

      // 单词后面要加上空格
      if (begin != front)
      {
        output[index] = s[begin];
      }

      // 进入下一个单词
      end = begin - 1;
      index++;
    }
    // 中间空格多时要跳过
    else if (s[begin] == ' ')
    {
      end--;
    }
  }

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
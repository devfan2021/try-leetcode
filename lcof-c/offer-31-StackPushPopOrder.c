#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

bool validateStackSequences(int *pushed, int pushedSize, int *popped, int poppedSize)
{
  if (pushed == NULL || popped == NULL)
  {
    return false;
  }
  if (pushedSize < poppedSize)
  {
    return false;
  }

  int *stack = (int *)malloc(sizeof(int) * pushedSize);
  int stack_top = -1;
  int pop_top = 0;
  for (int i = 0; i < pushedSize; i++)
  {
    stack[++stack_top] = pushed[i];
    while (stack_top >= 0 && stack[stack_top] == popped[pop_top])
    {
      stack_top--;
      pop_top++;
    }
  }

  return stack_top < 0;
}

bool validateStackSequences1(int *pushed, int pushedSize, int *popped, int poppedSize)
{
  int *stack = (int *)malloc(sizeof(int) * pushedSize);
  int size = 0;

  for (int push_index = 0, pop_index = 0; push_index < pushedSize || pop_index < poppedSize;)
  {
    if (size == 0 || stack[size - 1] != popped[pop_index])
    {
      // push
      if (push_index == pushedSize)
        break;

      stack[size++] = pushed[push_index++];
    }
    else
    {
      // pop
      --size;
      ++pop_index;
    }
  }

  return size == 0;
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

// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。 假设压入栈的所有数字均不相等。例如，序列{1, 2, 3, 4, 5} 是某栈的压栈序列，
// 序列{4, 5, 3, 2, 1} 是该压栈序列对应的一个弹出序列，但{4, 3, 5, 1, 2} 就不可能是该压栈序列的弹出序列。
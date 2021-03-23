#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/

#define max_size 10000

typedef struct
{
  int *data_stack;
  int data_top;
  int *min_stack;
  int min_top;

  int min_val;

} MinStack;

/** initialize your data structure here. */

MinStack *minStackCreate()
{
  MinStack *obj = (MinStack *)malloc(sizeof(MinStack));
  obj->data_stack = (int *)malloc(sizeof(int) * max_size);
  obj->min_stack = (int *)malloc(sizeof(int) * max_size);
  obj->data_top = -1;
  obj->min_top = -1;
  obj->min_val = INT_MAX;

  return obj;
}

void minStackPush(MinStack *obj, int x)
{
  if (obj->data_top < max_size)
  {
    if (obj->data_top == -1)
      obj->min_val = x;

    if (obj->min_val > x)
      obj->min_val = x;
    obj->min_stack[++(obj->min_top)] = obj->min_val;

    obj->data_stack[++(obj->data_top)] = x;
  }
}

void minStackPop(MinStack *obj)
{
  if (obj->data_top != -1)
  {
    --(obj->data_top);
    --(obj->min_top);
  }

  if (obj->min_top >= 0)
    obj->min_val = obj->min_stack[obj->min_top];
}

int minStackTop(MinStack *obj)
{
  if (obj->min_top < 0)
    return INT_MAX;

  return obj->data_stack[obj->data_top];
}

int minStackMin(MinStack *obj)
{
  if (obj->min_top < 0)
    return INT_MAX;

  return obj->min_stack[obj->min_top];
}

void minStackFree(MinStack *obj)
{
  free(obj->data_stack);
  free(obj->min_stack);
  free(obj);
}

/**
 * Your MinStack struct will be instantiated and called as such:
 * MinStack* obj = minStackCreate();
 * minStackPush(obj, x);
 
 * minStackPop(obj);
 
 * int param_3 = minStackTop(obj);
 
 * int param_4 = minStackMin(obj);
 
 * minStackFree(obj);
*/

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
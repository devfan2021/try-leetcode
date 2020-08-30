#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

typedef struct
{
  int len;
  int *push_stack;
  int push_top;
  int *pop_stack;
  int pop_top;

} CQueue;

CQueue *cQueueCreate()
{
  CQueue *q = (CQueue *)malloc(sizeof(CQueue));
  q->len = 10000;
  q->push_stack = (int *)malloc(sizeof(int) * q->len);
  q->push_top = -1;
  q->pop_stack = (int *)malloc(sizeof(int) * q->len);
  q->pop_top = -1;

  return q;
}

void cQueueAppendTail(CQueue *obj, int value)
{
  obj->push_stack[++(obj->push_top)] = value;
}

int cQueueDeleteHead(CQueue *obj)
{
  if (obj->pop_top == -1) // pop stack empty
  {
    if (obj->push_top == -1) // push stack empty, error, return -1
    {
      return -1;
    }

    while (obj->push_top > -1)
    {
      obj->pop_stack[++(obj->pop_top)] = obj->push_stack[(obj->push_top)--];
    }
  }

  return obj->pop_stack[(obj->pop_top)--];
}

void cQueueFree(CQueue *obj)
{
  free(obj->push_stack);
  obj->push_stack = NULL;
  free(obj->pop_stack);
  obj->pop_stack = NULL;
  free(obj);
  obj = NULL;
}

/**
 * Your CQueue struct will be instantiated and called as such:
 * CQueue* obj = cQueueCreate();
 * cQueueAppendTail(obj, value);
 
 * int param_2 = cQueueDeleteHead(obj);
 
 * cQueueFree(obj);
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
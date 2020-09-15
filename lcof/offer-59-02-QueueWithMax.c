#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/
// https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/solution/ru-he-jie-jue-o1-fu-za-du-de-api-she-ji-ti-by-z1m/

#define MAX_SIZE 10000

typedef struct
{
  int *queue;
  int front;
  int rear;

  int *max_data;
  int max_front;
  int max_rear;

} MaxQueue;

MaxQueue *maxQueueCreate()
{
  MaxQueue *obj = malloc(sizeof(MaxQueue));
  obj->queue = malloc(sizeof(int) * MAX_SIZE);
  obj->front = 0;
  obj->rear = 0;

  obj->max_data = malloc(sizeof(int) * MAX_SIZE);
  obj->max_front = 0;
  obj->max_rear = 0;

  return obj;
}

int maxQueueMax_value(MaxQueue *obj)
{
  if (obj == NULL || obj->front == obj->rear)
  {
    return -1;
  }

  return obj->max_data[obj->max_front];
}

void maxQueuePush_back(MaxQueue *obj, int value)
{
  if (obj == NULL)
  {
    return;
  }

  obj->queue[obj->rear++] = value;
  while (obj->max_front < obj->max_rear)
  {
    if (obj->max_data[obj->max_rear - 1] < value)
    {
      (obj->max_rear)--;
    }
    else
    {
      break;
    }
  }
  obj->max_data[obj->max_rear++] = value;
}

int maxQueuePop_front(MaxQueue *obj)
{
  if (obj == NULL || obj->front == obj->rear)
  {
    return -1;
  }

  int value = obj->queue[obj->front++];
  if (value == obj->max_data[obj->max_front])
  {
    (obj->max_front)++;
  }
  return value;
}

void maxQueueFree(MaxQueue *obj)
{
  if (obj != NULL)
  {
    if (obj->queue != NULL)
    {
      free(obj->queue);
    }
    if (obj->max_data != NULL)
    {
      free(obj->max_data);
    }
    free(obj);
  }
}

/**
 * Your MaxQueue struct will be instantiated and called as such:
 * MaxQueue* obj = maxQueueCreate();
 * int param_1 = maxQueueMax_value(obj);
 
 * maxQueuePush_back(obj, value);
 
 * int param_3 = maxQueuePop_front(obj);
 
 * maxQueueFree(obj);
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
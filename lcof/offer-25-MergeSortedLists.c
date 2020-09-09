#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode
{
  int val;
  struct ListNode *next;
};

struct ListNode *mergeTwoLists(struct ListNode *l1, struct ListNode *l2)
{
  if (l1 == NULL && l2 == NULL)
    return NULL;
  else if (l1 == NULL)
    return l2;
  else if (l2 == NULL)
    return l1;
  struct ListNode *head;
  struct ListNode *tail;
  if (l1->val <= l2->val) // 找出两个链表第一个值最小值选择做头部开始
  {
    head = l1;
    tail = l2;
  }
  else
  {
    head = l2;
    tail = l1;
  }
  struct ListNode *ret = head; // 用来返回的指针
  struct ListNode *p;          // 循环里用来转换的
  while (head->next != NULL && tail != NULL)
  {
    if (tail->val <= head->next->val)
    {
      p = tail->next;
      tail->next = head->next;
      head->next = tail;
      tail = p;
      head = head->next;
    }
    else
      head = head->next;
  }
  if (head->next == NULL) // 如果头部指针先走完，意味着尾部指针剩下的值全部大于头部最大值
  {
    head->next = tail;
  }
  return ret;
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
#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

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

struct ListNode *reverseList(struct ListNode *head)
{
  if (head == NULL)
  {
    return head;
  }

  struct ListNode *p = NULL;
  struct ListNode *pre = NULL;
  struct ListNode *cur = head;
  while (cur != NULL)
  {
    struct ListNode *next = cur->next;
    if (next == NULL) // end of link
    {
      p = cur;
    }

    cur->next = pre;
    pre = cur;
    cur = next;
  }

  return p;
}

struct ListNode *reverseList2(struct ListNode *head)
{
  struct ListNode *prev = NULL;
  struct ListNode *curr = head;

  while (curr != NULL)
  {
    struct ListNode *nextNode = curr->next;

    curr->next = prev;
    prev = curr;
    curr = nextNode;
  }

  return prev;
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
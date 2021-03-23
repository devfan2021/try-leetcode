#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

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

struct ListNode *getKthFromEnd(struct ListNode *head, int k)
{
  if (head == NULL)
  {
    return NULL;
  }

  struct ListNode *nextNode = head;
  while (k > 0 && nextNode != NULL)
  {
    nextNode = nextNode->next;
    k--;
  }

  if (k > 0)
  {
    return NULL;
  }

  struct ListNode *curNode = head;
  while (nextNode != NULL)
  {
    nextNode = nextNode->next;
    curNode = curNode->next;
  }

  return curNode;
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
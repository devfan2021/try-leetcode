#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
/**
 * Definition for singly-linked list.
 */
struct ListNode
{
  int val;
  struct ListNode *next;
};

struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB)
{
  if (headA == NULL || headB == NULL)
  {
    return NULL;
  }

  struct ListNode *nodeA = headA;
  struct ListNode *nodeB = headB;
  while (nodeA != nodeB)
  {
    if (nodeA != NULL)
    {
      nodeA = nodeA->next;
    }
    else
    {
      nodeA = headB;
    }

    if (nodeB != NULL)
    {
      nodeB = nodeB->next;
    }
    else
    {
      nodeB = headA;
    }
  }

  return nodeA;
}

struct ListNode *getIntersectionNode1(struct ListNode *headA, struct ListNode *headB)
{
  struct ListNode *headToA = headA, *headToB = headB;
  int countA = 0, countB = 0;

  while (headToA != NULL)
  {
    headToA = headToA->next;
    countA++;
  }

  while (headToB != NULL)
  {
    headToB = headToB->next;
    countB++;
  }

  headToA = headA, headToB = headB;
  if (countA > countB)
  {
    int gap = countA - countB;
    for (int i = 0; i < gap; i++)
    {
      headToA = headToA->next;
    }
  }
  else if (countB > countA)
  {
    int gap = countB - countA;
    for (int i = 0; i < gap; i++)
    {
      headToB = headToB->next;
    }
  }

  while (headToA != headToB)
  {
    headToA = headToA->next;
    headToB = headToB->next;
  }
  return headToA;
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
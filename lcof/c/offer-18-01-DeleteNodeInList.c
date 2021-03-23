#include <stdio.h>
#include "minunit.h"

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

struct ListNode *deleteNode(struct ListNode *head, int val)
{
  if (head == NULL)
  {
    return NULL;
  }

  if (head->val == val)
  {
    head = head->next;
  }
  else
  {
    struct ListNode *root = head;
    while (root->next != NULL)
    {
      if (root->next->val == val)
      {
        if (root->next->next == NULL)
        {
          root->next = NULL;
        }
        else
        {
          struct ListNode *nextNode = root->next->next;
          root->next->val = nextNode->val;
          root->next->next = nextNode->next;
        }
        break;
      }
      else
      {
        root = root->next;
      }
    }
  }
  return head;
}

struct ListNode *deleteNode1(struct ListNode *head, int val)
{
  if (head == NULL)
  {
    return NULL;
  }

  if (head->val == val)
  {
    return head->next;
  }

  struct ListNode *pre = head;
  while (pre->next != NULL && pre->next->val != val)
  {
    pre = pre->next;
  }

  if (pre->next != NULL)
  {
    pre->next = pre->next->next;
  }
  return head;
}

// two point
struct ListNode *deleteNode2(struct ListNode *head, int val)
{
  struct ListNode *pre = NULL;
  struct ListNode *cur = head;
  while (cur != NULL && cur->val != val)
  {
    pre = cur;
    cur = cur->next;
  }

  // head is the val
  if (pre == NULL)
  {
    return cur->next;
  }

  if (cur->val == val)
  {
    pre->next = cur->next;
  }
  return head;
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
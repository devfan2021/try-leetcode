#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
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

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *reversePrint(struct ListNode *head, int *returnSize)
{
  if (head == NULL)
  {
    *returnSize = 0;
    return NULL;
  }

  struct ListNode *root = head;
  int sum = 0;
  while (root != NULL)
  {
    sum++;
    root = root->next;
  }
  int *retVals = (int *)malloc(sizeof(int) * sum);

  *returnSize = sum;
  root = head;
  while (root != NULL)
  {
    retVals[--sum] = root->val;
    root = root->next;
  }

  return retVals;
}

MU_TEST(test_case)
{
  // 输入：head = [ 1, 3, 2 ]
  // 输出：[2, 3, 1]
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
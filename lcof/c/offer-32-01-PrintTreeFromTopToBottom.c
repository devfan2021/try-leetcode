#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

/**
 * Definition for a binary tree node.
 */
struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#define MAXLINE 2000
int *levelOrder(struct TreeNode *root, int *returnSize)
{
  *returnSize = 0;
  if (root == NULL)
  {
    return NULL;
  }
  int front = 0, tail = 0;
  struct TreeNode *queue[MAXLINE];
  int *res = (int *)malloc(sizeof(int) * MAXLINE);
  queue[tail++] = root;
  while (front < tail)
  {
    struct TreeNode *tmp = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    tmp = queue[front++];

    res[(*returnSize)++] = tmp->val;
    if (tmp->left != NULL)
    {
      queue[tail++] = tmp->left;
    }

    if (tmp->right != NULL)
    {
      queue[tail++] = tmp->right;
    }
  }

  return res;
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
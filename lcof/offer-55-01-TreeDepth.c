#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
/**
 * Definition for a binary tree node.
 */
struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

typedef struct TreeNode TreeNode;

int dfs(struct TreeNode *root, int level)
{
  if (root == NULL)
  {
    return level;
  }

  int leftVal = level, rightVal = level;

  if (root->left != NULL)
  {
    leftVal = dfs(root->left, level + 1);
  }

  if (root->right != NULL)
  {
    rightVal = dfs(root->right, level + 1);
  }

  return leftVal > rightVal ? leftVal : rightVal;
}

int maxDepth(TreeNode *root)
{
  if (root == NULL)
  {
    return 0;
  }

  return dfs(root, 1);
}

int maxDepth1(TreeNode *root)
{
  if (root == NULL)
  {
    return 0;
  }

  int left = maxDepth1(root->left) + 1;
  int right = maxDepth1(root->right) + 1;
  return left > right ? left : right;
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
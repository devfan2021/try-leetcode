#include <stdio.h>
#include <stdbool.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/

#define max(a, b) (a > b ? a : b)

/**
 * Definition for a binary tree node.
 */
struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

int dfs(struct TreeNode *root)
{
  if (root == NULL)
  {
    return 0;
  }

  int left = dfs(root->left);
  int right = dfs(root->right);

  return max(left, right) + 1;
}

int absVal(int a, int b)
{
  if (a > b)
  {
    return a - b;
  }
  else
  {
    return b - a;
  }
}

bool isBalanced(struct TreeNode *root)
{
  if (root == NULL)
  {
    return true;
  }

  int left = dfs(root->left);
  int right = dfs(root->right);

  return absVal(left, right) <= 1 && isBalanced(root->left) && isBalanced(root->right);
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
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

bool isBalanced(struct TreeNode *root)
{
  if (root == NULL)
  {
    return true;
  }

  int left = dfs(root->left);
  int right = dfs(root->right);
  if (left - right > 1 || right - left < -1)
  {
    return false;
  }

  return isBalanced(root->left) && isBalanced(root->right);
}

bool helper(struct TreeNode *pRoot, int *pDepth)
{
  if (pRoot == NULL)
  {
    *pDepth = 0;
    return true;
  }

  int left, right;
  if (helper(pRoot->left, &left) && helper(pRoot->right, &right))
  {
    int diff = left - right;
    if (diff <= 1 && diff >= -1)
    {
      *pDepth = 1 + (left > right ? left : right);
      return true;
    }
  }

  return false;
}

bool isBalanced2(struct TreeNode *root)
{
  int depth = 0;
  return helper(root, &depth);
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
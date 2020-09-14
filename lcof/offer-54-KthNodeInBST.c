#include <stdio.h>
#include <stdbool.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
/**
 * Definition for a binary tree node.
 */
struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

bool dfs(struct TreeNode *root, int k, int *val, int *count)
{
  if (root == NULL)
  {
    return false;
  }

  if (dfs(root->right, k, val, count))
  {
    return true;
  }
  if (++(*count) == k)
  {
    *val = root->val;
    return true;
  }

  return dfs(root->left, k, val, count);
}

int kthLargest(struct TreeNode *root, int k)
{
  int retVal, count = 0;
  dfs(root, k, &retVal, &count);
  return retVal;
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
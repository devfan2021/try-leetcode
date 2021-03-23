#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
/**
 * Definition for a binary tree node.
 */
struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

struct TreeNode *lowestCommonAncestor(struct TreeNode *root, struct TreeNode *p, struct TreeNode *q)
{
  if (root == NULL || p == root || q == root)
  {
    return root;
  }

  struct TreeNode *leftNode = lowestCommonAncestor(root->left, p, q);
  struct TreeNode *rightNode = lowestCommonAncestor(root->right, p, q);

  if (leftNode != NULL && rightNode != NULL)
  {
    return root;
  }

  if (leftNode != NULL)
  {
    return leftNode;
  }

  if (rightNode != NULL)
  {
    return rightNode;
  }
  return NULL;
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
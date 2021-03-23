#include <stdio.h>
#include "minunit.h"

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
  if (root == NULL)
  {
    return root;
  }
  if (p->val > root->val && q->val > root->val)
  {
    return lowestCommonAncestor(root->right, p, q);
  }
  else if (p->val < root->val && q->val < root->val)
  {
    return lowestCommonAncestor(root->left, p, q);
  }

  return root;
}

struct TreeNode *lowestCommonAncestor1(struct TreeNode *root, struct TreeNode *p, struct TreeNode *q)
{
  if (root == NULL)
  {
    return root;
  }

  while (root)
  {
    if (p->val > root->val && q->val > root->val)
    {
      root = root->right;
    }
    else if (p->val < root->val && q->val < root->val)
    {
      root = root->left;
    }
    else
    {
      return root;
    }
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
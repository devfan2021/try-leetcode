#include <stdio.h>
#include <stdbool.h>
#include "minunit.h"

struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

bool isMirror(struct TreeNode *node1, struct TreeNode *node2);

bool isSymmetric(struct TreeNode *root)
{
  if (root == NULL)
  {
    return true;
  }

  return isMirror(root, root);
}

bool isMirror(struct TreeNode *node1, struct TreeNode *node2)
{
  if (node1 == NULL && node2 == NULL)
  {
    return true;
  }

  if (node1 == NULL || node2 == NULL)
  {
    return false;
  }

  return node1->val == node2->val && isMirror(node1->left, node2->right) && isMirror(node1->right, node2->left);
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
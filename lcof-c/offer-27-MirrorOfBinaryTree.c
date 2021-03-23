#include <stdio.h>
#include "minunit.h"

struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

struct TreeNode *mirrorTree(struct TreeNode *root)
{
  if (root == NULL)
  {
    return root;
  }

  if (root->left == NULL && root->right == NULL)
  {
    return root;
  }

  struct TreeNode *leftTemp = root->left;
  root->left = mirrorTree(root->right);
  root->right = mirrorTree(leftTemp);
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
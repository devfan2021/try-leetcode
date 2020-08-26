#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

struct TreeNode *buildTree(int *preorder, int preorderSize, int *inorder, int inorderSize)
{
  if (preorder == NULL || preorderSize == 0 || inorder == NULL || inorderSize == 0)
  {
    return NULL;
  }

  int rootIndex = 0;
  for (int i = 0; i < inorderSize; i++)
  {
    if (inorder[i] == preorder[0])
    {
      rootIndex = i;
      break;
    }
  }

  struct TreeNode *root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
  root->val = preorder[0];

  root->left = buildTree(&preorder[1], rootIndex, &inorder[0], rootIndex);
  root->right = buildTree(&preorder[rootIndex + 1], preorderSize - rootIndex - 1, &inorder[rootIndex + 1], inorderSize - rootIndex - 1);
  return root;
}

MU_TEST(test_case)
{
  int preOrder[] = {3, 9, 20, 15, 7};
  int inOrder[] = {9, 3, 15, 20, 7};
  buildTree(preOrder, (int)sizeof(preOrder) / sizeof(preOrder[0]), inOrder, (int)sizeof(inOrder) / sizeof(inOrder[0]));
  // todo: add print tree method
  // [3,9,20,null,null,15,7]
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
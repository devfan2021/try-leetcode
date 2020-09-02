#include <stdio.h>
#include <stdbool.h>
#include "minunit.h"

struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

bool checkNode(struct TreeNode *A, struct TreeNode *B);

bool isSubStructure(struct TreeNode *A, struct TreeNode *B)
{
  if (A == NULL || B == NULL)
  {
    return false;
  }

  if (A->val == B->val && checkNode(A->left, B->left) && checkNode(A->right, B->right))
  {
    return true;
  }

  return isSubStructure(A->left, B) || isSubStructure(A->right, B);
}

///////////////////////////
//////// method 2 ////////
//////////////////////////
bool checkNode(struct TreeNode *A, struct TreeNode *B)
{
  if (B == NULL)
  {
    return true;
  }

  if (A == NULL)
  {
    return false;
  }

  return A->val == B->val && checkNode(A->left, B->left) && checkNode(A->right, B->right);
}

// 需要依次遍历每一个节点
bool searchTree(struct TreeNode *A, struct TreeNode *B)
{
  if (A == NULL)
  {
    return false;
  }

  if (checkNode(A, B))
  {
    return true;
  }

  if (searchTree(A->left, B))
  {
    return true;
  }

  if (searchTree(A->right, B))
  {
    return true;
  }

  return false;
}

bool isSubStructure2(struct TreeNode *A, struct TreeNode *B)
{
  if (A == NULL || B == NULL)
  {
    return false;
  }
  return searchTree(A, B);
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
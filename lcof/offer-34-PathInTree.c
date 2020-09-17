#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

#define MAX_SIZE 100
/**
 * Definition for a binary tree node.
 */
struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

void dfs(struct TreeNode *root, int target, int *returnSize, int **returnColumnSizes, int **output, int *nums, int numSize)
{
  if (root->left == NULL && root->right == NULL)
  {
    if (target - root->val == 0)
    {
      // copy cur path
      nums[numSize++] = root->val;
      int *path = (int *)malloc(sizeof(int) * (numSize + 1));
      memcpy(path, nums, sizeof(int) * numSize);
      (*returnColumnSizes)[*returnSize] = numSize;
      output[(*returnSize)++] = path;
    }
  }
  else
  {
    nums[numSize++] = root->val;
    int *path = (int *)malloc(sizeof(int) * (numSize + 1));
    memcpy(path, nums, sizeof(int) * numSize);

    if (root->left != NULL)
    {
      dfs(root->left, target - root->val, returnSize, returnColumnSizes, output, path, numSize);
    }
    if (root->right != NULL)
    {
      dfs(root->right, target - root->val, returnSize, returnColumnSizes, output, path, numSize);
    }
  }
}

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int **pathSum(struct TreeNode *root, int sum, int *returnSize, int **returnColumnSizes)
{
  if (root == NULL)
  {
    *returnSize = 0;
    **returnColumnSizes = 0;
    return NULL;
  }

  int **output = (int **)malloc(sizeof(int *) * MAX_SIZE);
  *returnColumnSizes = (int *)malloc(sizeof(int) * MAX_SIZE);
  int *nums = (int *)malloc(sizeof(int) * 2);
  *returnSize = 0;
  dfs(root, sum, returnSize, returnColumnSizes, output, nums, 0);
  return output;
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
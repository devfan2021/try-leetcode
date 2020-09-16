#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

/**
 * Definition for a binary tree node.
 */
struct TreeNode
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
#define MAX_SIZE 1500

int **levelOrder(struct TreeNode *root, int *returnSize, int **returnColumnSizes)
{
  if (root == NULL)
  {
    *returnSize = 0;
    **returnColumnSizes = 0;
    return NULL;
  }

  int **output = (int **)malloc(sizeof(int *) * MAX_SIZE);
  struct TreeNode *queue[MAX_SIZE];
  int level = 0, front = 0, rear = 0;
  queue[rear++] = root;
  while (front != rear)
  {
    int curColSize = rear - front;
    (*returnColumnSizes)[level] = curColSize;
    output[level] = (int *)malloc(sizeof(int) * curColSize);
    for (int i = 0; i < curColSize; i++)
    {
      struct TreeNode *curNode = queue[front];
      output[level][i] = curNode->val;
      front++;

      if (curNode->left)
      {
        queue[rear++] = curNode->left;
      }

      if (curNode->right)
      {
        queue[rear++] = curNode->right;
      }
    }

    // 换位
    if (level % 2 != 0)
    {
      for (int i = 0, j = curColSize - 1; i < j; i++, j--)
      {
        int tmp = output[level][i];
        output[level][i] = output[level][j];
        output[level][j] = tmp;
      }
    }

    level++;
  }

  *returnSize = level;
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
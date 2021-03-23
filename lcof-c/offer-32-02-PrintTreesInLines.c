#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/

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
    (*returnColumnSizes)[level] = rear - front;
    output[level] = (int *)malloc(sizeof(int) * (rear - front));
    for (int i = 0; i < (*returnColumnSizes)[level]; i++)
    {
      struct TreeNode *node = queue[front];
      output[level][i] = node->val;
      front++;

      if (node->left)
      {
        queue[rear++] = node->left;
      }
      if (node->right)
      {
        queue[rear++] = node->right;
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
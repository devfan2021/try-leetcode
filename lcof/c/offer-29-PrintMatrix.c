#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include "minunit.h"

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *spiralOrder(int **matrix, int matrixSize, int *matrixColSize, int *returnSize)
{
  if (matrixSize == 0 || *matrixColSize == 0)
  {
    *returnSize = 0;
    return NULL;
  }

  int col = (*matrixColSize), total = matrixSize * col;
  int *retVals = (int *)malloc(sizeof(int) * total);
  *returnSize = total;

  int rowBegin = 0, rowEnd = matrixSize - 1;
  int colBegin = 0, colEnd = col - 1;

  int index = 0;
  while (true)
  {
    // from left to right
    for (int i = colBegin; i <= colEnd; i++)
    {
      retVals[index++] = matrix[rowBegin][i];
    }
    rowBegin++;
    if (rowBegin > rowEnd)
    {
      break;
    }

    // from top to bottom
    for (int i = rowBegin; i <= rowEnd; i++)
    {
      retVals[index++] = matrix[i][colEnd];
    }
    colEnd--;
    if (colEnd < colBegin)
    {
      break;
    }

    // from right to left
    for (int i = colEnd; i >= colBegin; i--)
    {
      retVals[index++] = matrix[rowEnd][i];
    }
    rowEnd--;
    if (rowEnd < rowBegin)
    {
      break;
    }

    // from bootm to up
    for (int i = rowEnd; i >= rowBegin; i--)
    {
      retVals[index++] = matrix[i][colBegin];
    }
    colBegin++;
    if (colBegin > colEnd)
    {
      break;
    }
  }

  return retVals;
}

int *spiralOrder1(int **matrix, int matrixSize, int *matrixColSize, int *returnSize)
{
  if (matrix == NULL || matrixSize == 0)
  {
    *returnSize = 0;
    return NULL;
  }

  int m = matrixSize;
  int n = matrixColSize[0];
  int *ans = malloc(sizeof(int) * m * n);
  int count = 0;
  int t = 0, b = m - 1, l = 0, r = n - 1;

  while (1)
  {
    for (int i = l; i <= r; i++)
      ans[count++] = matrix[t][i]; // from left to right
    if (++t > b)
      break;
    for (int i = t; i <= b; i++)
      ans[count++] = matrix[i][r]; // from top to bottom
    if (--r < l)
      break;
    for (int i = r; i >= l; i--)
      ans[count++] = matrix[b][i]; // from right to left
    if (--b < t)
      break;
    for (int i = b; i >= t; i--)
      ans[count++] = matrix[i][l]; // from bottom to top
    if (++l > r)
      break;
  }

  *returnSize = m * n;
  return ans;
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
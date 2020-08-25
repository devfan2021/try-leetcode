#include <stdio.h>
#include <stdbool.h>
#include "minunit.h"

bool findNumberIn2DArray(int **matrix, int matrixSize, int *matrixColSize, int target)
{
  if (matrix == NULL || matrixSize == 0 || *matrixColSize == 0)
  {
    return false;
  }

  for (int row = 0, col = *matrixColSize - 1; row < matrixSize && col >= 0;)
  {
    if (target < matrix[row][col])
    {
      col--;
    }
    else if (target > matrix[row][col])
    {
      row++;
    }
    else
    {
      return true;
    }
  }

  return false;
}

MU_TEST(test_case)
{
  int a[5][5] = {
      {1, 4, 7, 11, 15},
      {2, 5, 8, 12, 19},
      {3, 6, 9, 16, 22},
      {10, 13, 14, 17, 24},
      {18, 21, 23, 26, 30}};
  int cols = 5;

  // https://cloud.tencent.com/developer/article/1451114,  array in function parameters
  int *p[5];
  for (int i = 0; i < 5; i++)
  {
    p[i] = &a[i][0];
  }
  mu_check(true == findNumberIn2DArray(p, 5, &cols, 5));
  mu_check(false == findNumberIn2DArray(p, 5, &cols, 20));
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
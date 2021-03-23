#include <stdio.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
#define max(a, b) (a > b ? a : b)

int maxValue(int **grid, int gridSize, int *gridColSize)
{
  if (grid == NULL || gridSize == 0 || *gridColSize == 0)
    return 0;

  for (int i = 1; i < *gridColSize; i++)
  {
    grid[0][i] += grid[0][i - 1];
  }

  for (int i = 1; i < gridSize; i++)
  {
    grid[i][0] += grid[i - 1][0];
  }

  for (int i = 1; i < gridSize; i++)
  {
    for (int j = 1; j < *gridColSize; j++)
    {
      if (grid[i - 1][j] > grid[i][j - 1])
      {
        grid[i][j] += grid[i - 1][j];
      }
      else
      {
        grid[i][j] += grid[i][j - 1];
      }
    }
  }

  return grid[gridSize - 1][(*gridColSize) - 1];
}

int maxValue1(int **grid, int gridSize, int *gridColSize)
{
  int dp[201][201] = {0};

  for (int i = 1; i <= gridSize; i++)
  {
    for (int j = 1; j <= *gridColSize; j++)
    {
      dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]) + grid[i - 1][j - 1];
    }
  }

  return dp[gridSize][*gridColSize];
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
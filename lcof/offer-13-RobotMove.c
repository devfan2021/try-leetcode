#include <stdio.h>
#include "minunit.h"

int cal(int x, int y)
{
  int ret = 0;
  int tmp = 0;
  while (x)
  {
    tmp += x % 10;
    x = x / 10;
  }
  while (y)
  {
    tmp += y % 10;
    y = y / 10;
  }
  return tmp;
}

int dfs(int m, int n, int k, int i, int j, int *visited)
{
  if (i < 0 || i == m || j < 0 || j == n || visited[i * n + j] == 1 || cal(i, j) > k)
  {
    return 0;
  }

  visited[i * n + j] = 1;
  return dfs(m, n, k, i + 1, j, visited) + dfs(m, n, k, i, j + 1, visited) + 1;
}

int movingCount(int m, int n, int k)
{
  int visited[m * n];
  memset(visited, 0, m * n * sizeof(int));
  return dfs(m, n, k, 0, 0, visited);
}

MU_TEST(test_case)
{
  mu_assert_int_eq(3, movingCount(2, 3, 1));
  mu_assert_int_eq(1, movingCount(3, 1, 0));
  mu_assert_int_eq(15, movingCount(16, 8, 4));
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
#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/

//n >= 5 2*(n-2) > n   3*(n-3) > n  且3*(n-3) >= 2*(n-2)
//n = 4 2 * 2 > 1 * 3
//2和3不能再分了  分了就变小了 且3优于2
int cuttingRope(int n)
{
  if (n <= 3)
    return n - 1;

  long rs = 1;
  while (n > 4)
  {
    //3最优
    rs *= 3;
    rs %= 1000000007;
    n -= 3;
  }
  //循环结束 n只剩下1, 2 ,3,4
  //1不能再分
  //2，3再分会标小
  //4 可以分成1 * 3  2 * 2,所以还是4最优
  //所以 剩下的1 2 3 4 都不能再分了
  return (rs * n) % 1000000007;
}

MU_TEST(test_case)
{
  mu_assert_int_eq(1, cuttingRope(2));
  mu_assert_int_eq(36, cuttingRope(10));
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
#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
/**
 * 利用异或运算 ^ 的特点
 * 1.0和任何数异或等于该数 
 * 2.任何数和自身异或等于0
 * 3.异或运算满足交换律
 * 
 * 相同的数异或为0，不同的异或为1。0和任何数异或等于这个数本身
 */

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *singleNumbers(int *nums, int numsSize, int *returnSize)
{
  if (nums == NULL || numsSize == 0)
  {
    *returnSize = 0;
    return NULL;
  }

  int mask = 0;
  for (int i = 0; i < numsSize; i++)
  {
    mask ^= nums[i];
  }

  mask = mask & (-1 * mask); // 末尾为1
  int m1 = 0, m2 = 0;
  for (int i = 0; i < numsSize; i++)
  {
    if ((mask & nums[i]) == 0)
    {
      m1 ^= nums[i];
    }
    else
    {
      m2 ^= nums[i];
    }
  }

  *returnSize = 2;
  int *retVals = malloc(sizeof(int) * 2);
  retVals[0] = m1;
  retVals[1] = m2;

  return retVals;
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
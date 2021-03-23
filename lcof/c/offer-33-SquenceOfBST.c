#include <stdio.h>
#include <stdbool.h>
#include "minunit.h"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/

bool verifyPostorder(int *postorder, int postorderSize)
{
  if (postorderSize <= 1)
  {
    return true;
  }

  int left_up = 0, up = postorderSize - 1;
  int root_value = postorder[up];
  while (left_up < up && postorder[left_up] < root_value)
  {
    // try to find left child tree
    left_up++;
  }

  int right_up = left_up;
  while (right_up < up && postorder[right_up] > root_value)
  {
    // check if there is some node in right child tree litter than root
    right_up++;
  }

  if (right_up != up)
  {
    // check if there is some node in right child tree litter than root
    return false;
  }

  return verifyPostorder(postorder, left_up) && verifyPostorder(postorder + left_up, up - left_up);
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
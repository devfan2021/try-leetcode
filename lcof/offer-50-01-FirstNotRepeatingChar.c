#include <stdio.h>
#include "minunit.h"

char firstUniqChar(char *s)
{
  if (s == NULL)
  {
    return ' ';
  }

  int hash[26] = {0};
  int len = strlen(s);
  for (int i = 0; i < len; i++)
  {
    ++hash[s[i] - 'a'];
  }

  for (int i = 0; i < len; i++)
  {
    if (hash[s[i] - 'a'] == 1)
    {
      return s[i];
    }
  }
  return ' ';
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
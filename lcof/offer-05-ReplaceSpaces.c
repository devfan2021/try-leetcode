#include <stdio.h>
#include <stdlib.h>
#include "minunit.h"

// malloc new space to store retval
char *replaceSpace(char *s)
{
  if (s == NULL)
  {
    return NULL;
  }
  int count = 0;
  for (int i = 0; s[i] != '\0'; i++)
  {
    count++;
  }

  char *retVal = (char *)malloc((3 * count + 1) * sizeof(char));
  int index = 0;
  for (int i = 0; s[i] != '\0'; i++)
  {
    if (s[i] == ' ')
    {
      retVal[index++] = '%';
      retVal[index++] = '2';
      retVal[index++] = '0';
    }
    else
    {
      retVal[index++] = s[i];
    }
  }
  retVal[index] = '\0';

  return retVal;
}

char *replaceSpace2(char *s)
{
  if (s == NULL)
  {
    return NULL;
  }
  int count = 0, whiteSpace = 0;
  for (int i = 0; s[i] != '\0'; i++)
  {
    count++;
    if (s[i] == ' ')
    {
      whiteSpace++;
    }
  }
  if (whiteSpace == 0)
  {
    return s;
  }

  s = (char *)realloc(s, sizeof(char) * (count + 2 * whiteSpace + 1));
  s[count + 2 * whiteSpace] = '\0'; // end
  int index = count + 2 * whiteSpace - 1;

  for (int i = count - 1; i >= 0; i--)
  {
    if (s[i] == ' ')
    {
      s[index--] = '0';
      s[index--] = '2';
      s[index--] = '%';
    }
    else
    {
      s[index--] = s[i];
    }
  }

  return s;
}

MU_TEST(test_case)
{
  // 输入：s = "We are happy." 输出：
  // "We%20are%20happy."
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
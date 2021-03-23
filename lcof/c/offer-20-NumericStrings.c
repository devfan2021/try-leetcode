#include <stdio.h>
#include <stdbool.h>
#include "minunit.h"

/**
 * https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
 * https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/solution/biao-shi-shu-zhi-de-zi-fu-chuan-by-leetcode-soluti/
 */
enum State
{
  STATE_INITIAL,
  STATE_INT_SIGN,
  STATE_INTEGER,
  STATE_POINT,
  STATE_POINT_WITHOUT_INT,
  STATE_FRACTION,
  STATE_EXP,
  STATE_EXP_SIGN,
  STATE_EXP_NUMBER,
  STATE_END,
  STATE_ILLEGAL
};

enum CharType
{
  CHAR_NUMBER,
  CHAR_EXP,
  CHAR_POINT,
  CHAR_SIGN,
  CHAR_SPACE,
  CHAR_ILLEGAL
};

enum CharType toCharType(char ch)
{
  if (ch >= '0' && ch <= '9')
  {
    return CHAR_NUMBER;
  }
  else if (ch == 'e' || ch == 'E')
  {
    return CHAR_EXP;
  }
  else if (ch == '.')
  {
    return CHAR_POINT;
  }
  else if (ch == '+' || ch == '-')
  {
    return CHAR_SIGN;
  }
  else if (ch == ' ')
  {
    return CHAR_SPACE;
  }
  else
  {
    return CHAR_ILLEGAL;
  }
}

enum State transfer(enum State st, enum CharType typ)
{
  switch (st)
  {
  case STATE_INITIAL:
  {
    switch (typ)
    {
    case CHAR_SPACE:
      return STATE_INITIAL;
    case CHAR_NUMBER:
      return STATE_INTEGER;
    case CHAR_POINT:
      return STATE_POINT_WITHOUT_INT;
    case CHAR_SIGN:
      return STATE_INT_SIGN;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_INT_SIGN:
  {
    switch (typ)
    {
    case CHAR_NUMBER:
      return STATE_INTEGER;
    case CHAR_POINT:
      return STATE_POINT_WITHOUT_INT;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_INTEGER:
  {
    switch (typ)
    {
    case CHAR_NUMBER:
      return STATE_INTEGER;
    case CHAR_EXP:
      return STATE_EXP;
    case CHAR_POINT:
      return STATE_POINT;
    case CHAR_SPACE:
      return STATE_END;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_POINT:
  {
    switch (typ)
    {
    case CHAR_NUMBER:
      return STATE_FRACTION;
    case CHAR_EXP:
      return STATE_EXP;
    case CHAR_SPACE:
      return STATE_END;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_POINT_WITHOUT_INT:
  {
    switch (typ)
    {
    case CHAR_NUMBER:
      return STATE_FRACTION;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_FRACTION:
  {
    switch (typ)
    {
    case CHAR_NUMBER:
      return STATE_FRACTION;
    case CHAR_EXP:
      return STATE_EXP;
    case CHAR_SPACE:
      return STATE_END;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_EXP:
  {
    switch (typ)
    {
    case CHAR_NUMBER:
      return STATE_EXP_NUMBER;
    case CHAR_SIGN:
      return STATE_EXP_SIGN;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_EXP_SIGN:
  {
    switch (typ)
    {
    case CHAR_NUMBER:
      return STATE_EXP_NUMBER;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_EXP_NUMBER:
  {
    switch (typ)
    {
    case CHAR_NUMBER:
      return STATE_EXP_NUMBER;
    case CHAR_SPACE:
      return STATE_END;
    default:
      return STATE_ILLEGAL;
    }
  }
  case STATE_END:
  {
    switch (typ)
    {
    case CHAR_SPACE:
      return STATE_END;
    default:
      return STATE_ILLEGAL;
    }
  }
  default:
    return STATE_ILLEGAL;
  }
}

bool isNumber(char *s)
{
  int len = strlen(s);
  enum State st = STATE_INITIAL;

  for (int i = 0; i < len; i++)
  {
    enum CharType type = toCharType(s[i]);
    enum State nextState = transfer(st, type);
    if (nextState == STATE_ILLEGAL)
    {
      return false;
    }
    st = nextState;
  }
  return st == STATE_INTEGER || st == STATE_POINT || st == STATE_FRACTION || st == STATE_EXP_NUMBER || st == STATE_END;
}

bool isNumber2(char *s)
{
  int i = 0;
  int end = strlen(s) - 1;
  int dot = 0;
  int num = 0;
  int numE = 0;
  while (s[i] == ' ')
  {
    i++;
  }

  while (end >= 0 && s[end] == ' ')
  {
    end--;
  }

  if (s[i] == '+' || s[i] == '-')
  {
    i++;
  }

  while (i <= end)
  {
    if (s[i] >= '0' && s[i] <= '9')
    {
      num++;
    }
    else if (s[i] == '.')
    {
      dot++;
    }
    else if (s[i] == 'e' || s[i] == 'E')
    {
      if (dot > 0 || num == 0 || numE > 0)
      {
        return false;
      }
      numE++;
      dot = 0;
      num = 0;

      if (i + 1 < end && (s[i + 1] == '+' || s[i + 1] == '-'))
      {
        i++;
      }
    }
    else
    {
      return false;
    }
    i++;
  }

  if ((numE == 1 && dot > 0) || (numE == 0 && dot > 1))
  {
    return false;
  }

  return num > 0;
}

MU_TEST(test_case)
{
  mu_check(isNumber("+100") == true);
  mu_check(isNumber("5e2") == true);
  mu_check(isNumber("-123") == true);
  mu_check(isNumber("3.1416") == true);
  mu_check(isNumber("-1E-16") == true);
  mu_check(isNumber("0123") == true);
  mu_check(isNumber("12e") == false);
  mu_check(isNumber("1a3.14") == false);
  mu_check(isNumber("1.2.3") == false);
  mu_check(isNumber("+-5") == false);
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
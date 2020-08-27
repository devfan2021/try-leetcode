#include <stdio.h>
#include <stdbool.h>
#include "minunit.h"

bool dfs(char **board, int boardSize, int *boardColSize, char *word, int row, int col, int len)
{
  if (row < 0 || col < 0 || row >= boardSize || col >= *boardColSize || board[row][col] != word[len])
  {
    return false;
  }

  if (strlen(word) - 1 == len)
  {
    return true;
  }

  char tmp = board[row][col];
  // mark this element, avoid repeat check
  board[row][col] = '/';

  bool res = dfs(board, boardSize, boardColSize, word, row, col + 1, len + 1) ||
             dfs(board, boardSize, boardColSize, word, row + 1, col, len + 1) ||
             dfs(board, boardSize, boardColSize, word, row, col - 1, len + 1) ||
             dfs(board, boardSize, boardColSize, word, row - 1, col, len + 1);

  // roback reset
  board[row][col] = tmp;
  return res;
}

bool exist(char **board, int boardSize, int *boardColSize, char *word)
{
  for (int i = 0; i < boardSize; i++)
  {
    for (int j = 0; j < *boardColSize; j++)
    {
      if (dfs(board, boardSize, boardColSize, word, i, j, 0))
      {
        return true;
      }
    }
  }

  return false;
}

MU_TEST(test_case)
{
  char atrix[3][4] = {{'A', 'B', 'C', 'E'}, {'S', 'T', 'C', 'S'}, {'A', 'D', 'E', 'E'}};
  char *p[3];
  for (int i = 0; i < 3; i++)
  {
    p[i] = &atrix[i][0];
  }
  // char atrix = "ABTGCFCSJDEH";
  char *word = "ABCCED";
  int colSize = 4;
  mu_check(true == exist(p, 3, &colSize, word));

  char atrix2[2][2] = {{'a', 'b'}, {'c', 'd'}};
  char *p2[2];
  for (int i = 0; i < 3; i++)
  {
    p2[i] = &atrix2[i][0];
  }
  // char atrix = "ABTGCFCSJDEH";
  char *word2 = "abcd";
  int colSize2 = 2;
  mu_check(false == exist(p2, 2, &colSize2, word2));
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
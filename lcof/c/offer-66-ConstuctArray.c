#include <stdio.h>
#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 * C[i]=A[0]*A[1]*A[2]*...*A[i-1]  D[i]=A[i+1]*A[i+1]*...*A[n-1]
 * C[i]=C[i-1]*A[i-1]  D[i]=D[i+1]*A[i+1]
 * B[i]=C[i]*D[i]
 */
int *constructArr(int *a, int aSize, int *returnSize)
{
  if (a == NULL || aSize <= 1)
  {
    *returnSize = 0;
    return a;
  }

  *returnSize = aSize;
  int *arr = (int *)malloc(sizeof(int) * aSize);
  arr[0] = 1;

  // caculate left bottom area
  for (int i = 1; i < aSize; i++)
  {
    arr[i] = arr[i - 1] * a[i - 1];
  }

  // iterator from last element
  // B[i]=C[i]*D[i]  => arr[i]= arr[i] * (D[i+1]*A[i+1])
  int temp = 1;
  for (int i = aSize - 2; i >= 0; i--)
  {
    temp *= a[i + 1];
    arr[i] *= temp;
  }

  return arr;
}

int *constructArr1(int *a, int aSize, int *returnSize)
{
  if (a == NULL || aSize <= 1)
  {
    *returnSize = aSize;
    return a;
  }

  int *arrayA = (int *)malloc(sizeof(int) * aSize);
  arrayA[0] = 1;
  int *arrayB = (int *)malloc(sizeof(int) * aSize);
  arrayB[0] = 1;

  for (int i = 1; i < aSize; i++)
  {
    arrayA[i] = arrayA[i - 1] * a[i - 1];
    arrayB[i] = arrayB[i - 1] * a[aSize - i];
  }

  int *retArray = arrayA;
  for (int i = 0; i < aSize; i++)
  {
    retArray[i] = arrayA[i] * arrayB[aSize - i - 1];
  }
  *returnSize = aSize;

  return retArray;
}

int main()
{
  int a[] = {1, 2, 3, 4, 5};
  int bSize;
  int *b = constructArr(a, (int)(sizeof(a) / sizeof(a[0])), &bSize);
  if (b != NULL)
  {
    for (int i = 0; i < bSize; i++)
    {
      printf("%d\n", b[i]);
    }
  }

  printf("======\n");
  b = constructArr1(a, (int)(sizeof(a) / sizeof(a[0])), &bSize);
  if (b != NULL)
  {
    for (int i = 0; i < bSize; i++)
    {
      printf("%d\n", b[i]);
    }
  }
  free(b);

  return 0;
}
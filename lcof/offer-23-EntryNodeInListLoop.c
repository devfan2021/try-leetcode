#include <stdio.h>
#include "minunit.h"

/**
 * 入口节点即遍历过程第一个再次到达的节点
 * 第一步：确定链表中包含环 
 * 方法：两个指针同时从头节点出发，一个每次一步，一个每次两步。如果走的快的指针追上了走的慢的，说明有环。如果走的快的指针走到了链表尾，那么说明没环。 
 * 
 * 第二步：找到环中的节点数
 * 方法：因为上述两个节点相遇的节点一定在环中，所以从这个节点出发，再次走到这个节点，即可计算出环中的节点数。 
 * 
 * 第三步：找到链表中环的入口节点
 * 方法：两个指针从链表头节点出发，一个先走环的节点数步，然后两个指针同时往前走，两个指针相遇的节点即是环的入口节点
 */

typedef struct ListNode
{
  int val;
  struct ListNode *next;
} ListNode;

ListNode *meetingNode(ListNode *pHead)
{
  if (pHead == NULL)
  {
    return NULL;
  }

  ListNode *pSlow = pHead->next;
  if (pSlow == NULL)
  {
    return NULL;
  }

  ListNode *pFast = pSlow->next;
  while (pFast != NULL && pSlow != NULL)
  {
    if (pFast == pSlow)
      return pFast;

    pSlow = pSlow->next;

    pFast = pFast->next;
    if (pFast != NULL)
      pFast = pFast->next;
  }

  return NULL;
}

ListNode *entryNodeOfLoop(ListNode *pHead)
{
  ListNode *meetingNode = meetingNode(pHead);
  if (NULL == meetingNode)
    return NULL;

  // 得到环中节点的数目
  int nodesInLoop = 1;
  ListNode *pNode1 = meetingNode;
  while (pNode1->next != meetingNode)
  {
    pNode1 = pNode1->next;
    ++nodesInLoop;
  }

  // 先移动pNode1，次数为环中节点的数目
  pNode1 = pHead;
  for (int i = 0; i < nodesInLoop; i++)
  {
    pNode1 = pNode1->next;
  }

  // 再移动pNode1和pNode2
  ListNode *pNode2 = pHead;
  while (pNode1 != pNode2)
  {
    pNode1 = pNode1->next;
    pNode2 = pNode2->next;
  }

  return pNode1;
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
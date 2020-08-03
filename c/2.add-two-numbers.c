#include <stdio.h>
/*
 * @lc app=leetcode id=2 lang=c
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (33.51%)
 * Likes:    8176
 * Dislikes: 2084
 * Total Accepted:    1.4M
 * Total Submissions: 4.2M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 * 
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 * 
 * Example:
 * 
 * 
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *addTwoNumbers2(struct ListNode *l1, struct ListNode *l2)
{
	struct ListNode anchor = {.next = NULL}, *curr = &anchor;
	int carry = 0;
	while (l1 || l2)
	{
		int sum = 0, d1 = 0, d2 = 0;

		if (l1)
		{
			d1 = l1->val;
			l1 = l1->next;
		}
		if (l2)
		{
			d2 = l2->val;
			l2 = l2->next;
		}

		sum = carry + d1 + d2;
		carry = sum > 9;
		sum = sum % 10;

		curr = curr->next = malloc(sizeof(struct ListNode));
		curr->val = sum;
		curr->next = NULL;
	}

	if (carry)
	{
		curr = curr->next = malloc(sizeof(struct ListNode));
		curr->val = 1;
		curr->next = NULL;
	}

	return anchor.next;
}

struct ListNode *addTwoNumbers1(struct ListNode *l1, struct ListNode *l2)
{
	int carry = 0;
	struct ListNode *dummy = (struct ListNode *)malloc(sizeof(struct ListNode));
	dummy->next = NULL;
	struct ListNode *curNode = dummy;
	while (l1 != NULL || l2 != NULL || carry > 0)
	{
		if (l1 != NULL)
		{
			carry += l1->val;
			l1 = l1->next;
		}

		if (l2 != NULL)
		{
			carry += l2->val;
			l2 = l2->next;
		}
		struct ListNode *node = (struct ListNode *)malloc(sizeof(struct ListNode));
		node->val = carry % 10;
		node->next = NULL;

		curNode->next = node;
		carry = carry / 10;
		curNode = node;
	}
	return dummy->next;
}

struct ListNode *addTwoNumbers(struct ListNode *l1, struct ListNode *l2)
{
	return addTwoNumbers2(l1, l2);
}
// @lc code=end
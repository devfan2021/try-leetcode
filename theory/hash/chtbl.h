/* chtbl.h */
#ifndef CHTBL_H
#define CHTBL_H
#include <stdlib.h>
#include "../list/list.h"

/* Define a structure for chained hash tables. */
typedef struct CHTBL_
{
  int buckets;
  int (*h)(const void *key);
  int (*match)(const void *key1, const void *key2);
  void (*destory)(void *data);
  int size;
  List *table;
} CHTBL;

/* Public Interface */
int chtbl_init(CHTBL *htbl, int buckets, int (*h)(const void *key), int (*match)(const void *key1, const void *key2), void (*destory)(void *data));
void chtbl_destory(CHTBL *htbl);
int chtbl_insert(CHTBL *htbl, const void *data);
int chtbl_remove(CHTBL *htbl, const void **data);
int chtbl_lookup(const CHTBL *htbl, void **data);

#define chtbl_size(htbl) ((htbl)->size)

#endif
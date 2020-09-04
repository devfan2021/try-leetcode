/* frame.h */
#ifndef FRAME_H
#define FRAME_H
#include "list.h"

/* alloc_frame */
int alloc_frame(List *frames);

/* free_frame */
int free_frame(List *frames, int frame_number);

#endif
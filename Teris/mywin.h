//declaration
#ifndef __MYWIN_H
#define __MYWIN_H
#include <termios.h>
#include <unistd.h>
#include <stdlib.h>
#include <curses.h>
#include <signal.h>
#include <setjmp.h>
#include <sys/time.h>
#include <pthread.h>


#define  X 20                    // game_window_size
#define  Y 20

#define MATRIX_ROW 18
#define MATRIX_COL 20

jmp_buf env;

WINDOW *game_win_ptr;
WINDOW *next_win_ptr;
WINDOW *score_win_ptr;

typedef struct itimerval LEVEL;
static LEVEL level_00 = { {0,      0}, {0,      0} };
static LEVEL level_01 = { {0, 800000}, {1,      0} };
static LEVEL level_02 = { {0, 500000}, {0, 500000} };
static LEVEL level_03 = { {0, 400000}, {0, 300000} };
static LEVEL level_04 = { {0, 300000}, {0, 300000} };
static LEVEL level_05 = { {0, 200000}, {0, 300000} };
static LEVEL level_06 = { {0, 150000}, {0, 300000} };
static LEVEL level_07 = { {0, 100000}, {0, 300000} };
static LEVEL level_08 = { {0, 80000 }, {0, 300000} };
static LEVEL level_09 = { {0, 60000 }, {0, 300000} };

struct block{
	int kind;
	int mode;
	int color;
	int col;
	int row;
};
pthread_mutex_t cur_mutex;
pthread_mutex_t save_mutex;
pthread_mutex_t next_mutex;

struct block cur_block;
struct block save_block;
struct block next_block;
struct block new_block;

int flag_erase;
int game_score;
int game_level;


struct termios initial_setting,new_setting;
int peek_character;

static  int matirx[MATRIX_ROW][MATRIX_COL] = { 
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //1	
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //2
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //3
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //4
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //5
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //6
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //7
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //8
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, //9
	1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //10
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //11
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //12
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //13
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //14
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //15
	1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //16
	1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //17
	1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, //18
	
};



static const int shape[7][4][18] = {
                              {//{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,     2, 1},
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1,     2, 1},   //        
                               {0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0,     1, 2},   //   []   []    [][][]     []
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0,     2, 1},   // [][][] [][]    []     [][]
                               {0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1,     1, 2}}, //         []               []
 
                              {{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1,     2, 1},   //
                               {0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1,     1, 2},  //          []            [][]
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0,     2, 1},   //      [] []    [][][]    []
                               {0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1,     1, 2}},   // [][][] [][]  []        []
 
                              {{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1,     1, 2},   //
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1,     2, 1},   //          [][]            []
                               {0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0,     1, 2},   // []       []    [][][]    []
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1,     2, 1}},  // [][][]   []        []  [][]
 
                              {{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1,     1, 2},   // 
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0,     2, 1},   //    []
                               {0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1,     1, 2},   //    [][]     [][]
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0,     2, 1}}, //       []   [][]
 
                              {{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0,     1, 2},   //
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1,     2, 1},   //      []
                               {0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0,     1, 2},   //    [][]   [][]
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1,     2, 1}},  //    []       [][]
 
                              {{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1,     2, 2},   //
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1,     2, 2},   //
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1,     2, 2},   //     [][]
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1,     2, 2}},  //     [][]
 
                              {{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,     0, 3},   //     []
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,     3, 0},   //     []     [][][][]
                               {0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,     0, 3},   //     [] 
                               {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,     3, 0}}   //     []     
                             };
void welcome(void);
void print_matix(void);
void print_next_block(void);
void print_block(WINDOW * winptr,struct block dis_b);
void clear_block(WINDOW * winptr,struct block dis_b);
void move_down(int signo );
int judge_by_color(struct block judge);
struct block create_block(void);
struct block next_to_cur(struct  block  block_in);
void destroy_line(void);
void game_over(void);
void store_flag_color(struct block arg);
void print_matrix(void);
void print_level(void);
void print_score(void);
void key_init(void);
void key_close(void);
int kbhit(void);
char readch(void);
void move_left(void);
void move_right(void);
void change_shape(void);
void key_deal(char key);
#endif
//this file define the function 
#include "mywin.h"

void welcome(void)
{
	
    initscr();
	refresh();
	//start_color();
	game_win_ptr = newwin(20,22,5,5);
	box(game_win_ptr,'|','=');
	wrefresh(game_win_ptr);
	next_win_ptr = newwin(7,12, 5,27);
	mvwprintw(next_win_ptr,1,1,"%s","next:");
	box(next_win_ptr,'|','-');
	touchwin(next_win_ptr);
	wrefresh(next_win_ptr);
	score_win_ptr =newwin(13,12, 12,27);
	touchwin(score_win_ptr);
	printf("\33[32m\n");
	mvwprintw(score_win_ptr,2,2,"%s","score:");
	touchwin(score_win_ptr);
	wrefresh(score_win_ptr);
	printf("\33[0m\n");

	mvwprintw(score_win_ptr,5,2,"%s","level:");
	box(score_win_ptr,'|','-');
	wrefresh(score_win_ptr);
	
    flag_erase=1;
	next_block=create_block();
	cur_block = next_to_cur(next_block);
	cur_block.kind=6;
	cur_block.mode=1;
    next_block=create_block();
    flag_erase=1;
	print_block(next_win_ptr,next_block);
	pthread_mutex_init(&cur_mutex,NULL);
	
	game_score=0;
	game_level=1;
	peek_character = -1;
}


void clear_block(WINDOW * winptr,struct block dis_b)
{
	int row=0;
	int col=0;
	int i=0;
	pthread_mutex_lock(&cur_mutex);
	for(i=0;i<16;i++)
	{
		if(i%4==0)
		{
			row++;
			col=0;
		}
		if(shape[dis_b.kind][dis_b.mode][i]==1)
		{	
			mvwprintw(winptr,dis_b.row + row,dis_b.col + 2*col, "  ");
		}
		col++;		
	}	
	pthread_mutex_unlock(&cur_mutex);
	touchwin(winptr);
	wrefresh(winptr);
	

}
void print_block(WINDOW * winptr,struct block dis_b)
{
	int row=0;
	int col=0;
	int i=0;
	pthread_mutex_lock(&cur_mutex);
	if(flag_erase==0)
	{
		pthread_mutex_unlock(&cur_mutex);
		    clear_block(game_win_ptr,save_block);
		pthread_mutex_lock(&cur_mutex);
	}

	for(i=0;i<16;i++)
	{
		if(i%4==0)
		{
			row++;
			col=0;
		}
		if(shape[dis_b.kind][dis_b.mode][i]==1)
		{	
			mvwprintw(winptr,dis_b.row + row,dis_b.col + 2*col, "[]");
		}
		col++;		
	}	

	touchwin(winptr);
	wrefresh(winptr);
	flag_erase=0;
	pthread_mutex_unlock(&cur_mutex);
    fflush(stdout);
}




void store_flag_color(struct block arg)
{
	int i=0;
	int row=0;
	int col=0;
	pthread_mutex_lock(&cur_mutex);
	for (i=0;i<16;i++)
	{
		if(i%4==0)
		{
			row++;
			col=0;
		}
		if(shape[save_block.kind][save_block.mode][i]==1)
		{
			matirx[save_block.row+row-1][save_block.col+2*col-1]=1;
			matirx[save_block.row+row-1][save_block.col+2*col]=1;
		}
		col++;
	}
	pthread_mutex_unlock(&cur_mutex);
	print_matrix();

	touchwin(game_win_ptr);
	wrefresh(game_win_ptr);
	
}


void 
game_over(void)
{
	if(judge_by_color(cur_block))
	{
		longjmp(env, 2);
	}


}
int judge_by_color(struct block arg)
{
	int i=0;
	int row =0;
	int col=0;

	for(i=0;i<16;i++)
	{
		if(i%4==0)
		{
			row++;
			col=0;
		}
		if(shape[arg.kind][arg.mode][i]==1)
		{
			if(matirx[arg.row+row-1][arg.col+2*col-1]==1)
			{
				return 1;
			}
		}
		col++;
	}

	return 0;
}


// move down
void 
move_down(int signo )
{
	pthread_mutex_lock(&cur_mutex);
    cur_block.row++;
    pthread_mutex_unlock(&cur_mutex);
    if (cur_block.row >= Y - 5 || judge_by_color(cur_block)==1 )
    {
    	store_flag_color(save_block);
        cur_block = next_to_cur(next_block);
    	clear_block(next_win_ptr,next_block);
        next_block=create_block(); 
        flag_erase = 1;    
        print_block(next_win_ptr,next_block);  
        save_block=cur_block;          
        print_block(game_win_ptr,cur_block);    
        destroy_line();
        game_over();        
        fflush(stdout);
        return;
    }

    print_block(game_win_ptr,cur_block);
    pthread_mutex_lock(&cur_mutex);
    save_block=cur_block;
    pthread_mutex_unlock(&cur_mutex);
    fflush(stdout);
}


void print_level(void)
{

}
void print_score(void)
{

}
void print_matrix(void)
{
	
	int x_cnt=0;
	int y_cnt=0;
	pthread_mutex_lock(&cur_mutex);
	for(x_cnt=0; x_cnt<X; x_cnt+=2)
		for(y_cnt=0; y_cnt<Y-2; y_cnt++)
		{
			if(matirx[y_cnt][x_cnt]==0)
				mvwprintw(game_win_ptr,y_cnt+1,x_cnt+1, "  ");
			else 
				mvwprintw(game_win_ptr,y_cnt+1,x_cnt+1, "[]");
		}

	touchwin(game_win_ptr);
	wrefresh(game_win_ptr);

	pthread_mutex_unlock(&cur_mutex);
	fflush(stdout);
	
}

void
destroy_line(void)
{
	
	int row =0;
	int col=0;
	int i=0;
	int j=0;
	int full=1;
	int a=0;
	int b=0;
	for(i=0;i<Y-2;i++)
	{
		full=1;
		for(j=0;j<X;j++)
		{
			if(matirx[i][j]==0)
				full=0;
		}
		if(full)
		{
			for(a=0;a<i;a++)
			{
				for(b=0;b<X;b++)
				{
					matirx[i-a][b]=matirx[i-a-1][b];
				}
			}
			print_matrix();
			game_score++;
			if(game_score>=100)
			{
				game_score=0;
				game_level++;
				print_level();
			}
			print_score();
			full=0;
		}
	}
	print_matrix();
}
 

struct block create_block(void)
{
	
	struct block block_out;
	pthread_mutex_lock(&cur_mutex);
    block_out.kind = random() % 7;
    block_out.mode = random() % 4;
    block_out.color = random() % 7 + 41;
    block_out.row = 1;
    block_out.col = 1;
    pthread_mutex_unlock(&cur_mutex);
    return block_out;
}
struct block next_to_cur(struct  block  block_in)
{
	struct block block_out;
	pthread_mutex_lock(&cur_mutex);
	block_out.kind = block_in.kind;
    block_out.mode = block_in.mode;
    block_out.color = block_in.color;
    block_out.row = -shape[block_in.kind][block_in.mode][16];
    block_out.col = X/2-3;
    pthread_mutex_unlock(&cur_mutex);
    return block_out;
}
void key_init(void)
{
	tcgetattr(0,&initial_setting);
	new_setting = initial_setting;
	new_setting.c_lflag &= ~ICANON;
	new_setting.c_lflag &= ~ECHO;
	new_setting.c_lflag &= ~ISIG;
	new_setting.c_cc[VMIN]=1;
	new_setting.c_cc[VTIME]=0;
	tcsetattr(0,TCSANOW,&new_setting);

}

void key_close(void)
{
	tcsetattr(0,TCSANOW,&initial_setting);
}

int kbhit(void)
{
	char ch;
	int nread;

	if(peek_character !=-1)
		return 1;

//	new_setting.c_cc[VMIN]=0;
//	tcsetattr(0,TCSANOW,&new_setting);
	nread = read(0,&ch,1);
//	sleep(1);
//	new_setting.c_cc[VMIN]=1;
//	tcsetattr(0,TCSANOW,&new_setting);

	if(nread == 1)
	{
		peek_character=ch;
		return 1;
	}
	return 0;
}
char readch(void)
{
	char ch;
	if(peek_character !=-1)
	{
		ch = peek_character;
		peek_character = -1;
		return ch;
	}
	read(0,&ch,1);
	return ch;
}

void key_deal(char key)
{
	switch(key)
	{
		case 'a':
		    move_left();
		    break;
		case 's':
		    move_down(1);
		    break;
		case 'd':
		    move_right();
		    break;
		case 'w':
		    change_shape();
		    break;
		default:
		    break;
	}
}

void move_left(void)
{	
	pthread_mutex_lock(&cur_mutex);
	cur_block.col -= 2;
	if (cur_block.col <  -2*shape[cur_block.kind][cur_block.mode][17] || judge_by_color(cur_block) == 1) 
	{
		cur_block.col += 2;
		pthread_mutex_unlock(&cur_mutex);
        return;
    }
    pthread_mutex_unlock(&cur_mutex);
    print_block(game_win_ptr,cur_block);
    pthread_mutex_lock(&cur_mutex);
    save_block =cur_block;
    pthread_mutex_unlock(&cur_mutex);
    fflush(stdout);
}

void move_right(void)
{
	pthread_mutex_lock(&cur_mutex);
	cur_block.col += 2;
	if (cur_block.col >=  X-6 || judge_by_color(cur_block) == 1) 
	{		
		cur_block.col -= 2;
        pthread_mutex_unlock(&cur_mutex);	
        return;
    }
    pthread_mutex_unlock(&cur_mutex);
    print_block(game_win_ptr,cur_block);
    pthread_mutex_lock(&cur_mutex);
    save_block =cur_block;
    pthread_mutex_unlock(&cur_mutex);
    fflush(stdout);

}
void change_shape(void)
{
    struct block  tmp_block;
    tmp_block=cur_block;    
    cur_block.mode = (cur_block.mode + 1) % 4;
    int i, n;
    int row_limit =shape[cur_block.kind][cur_block.mode][17];
    int col_limit = shape[cur_block.kind][cur_block.mode][16];
    if (( cur_block.col +6 - 2 * (col_limit - row_limit) <= 1 && cur_block.col <= -1 && row_limit < col_limit ) ||
        ( cur_block.row +6 - 2 * (row_limit - col_limit) <= 1 && cur_block.row <= -1 && row_limit > col_limit ) ||
    	judge_by_color(cur_block) == 1) 
    {
    	cur_block = tmp_block;
        return;
    }
    fflush(stdout);
    pthread_mutex_unlock(&cur_mutex);
    print_block(game_win_ptr,cur_block);
    pthread_mutex_lock(&cur_mutex);
    save_block =cur_block;
    pthread_mutex_unlock(&cur_mutex);
    fflush(stdout);
}


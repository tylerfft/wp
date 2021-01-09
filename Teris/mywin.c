// this file if the main
#include "mywin.h"

int main(int agrc,char** argv)
{
	char ch;
	key_init();	
	printf("\033[?25l"); 

	initscr();
	welcome();
	print_matrix();
	destroy_line();	
	setitimer(ITIMER_REAL, &level_01, NULL);    //init one leve ;interval 800ms
    signal(SIGALRM, move_down); 
    if (setjmp(env) == 0) 
    {
    	while(ch != 'q')
    	{
    		if(kbhit())
    		{
    			ch =readch();
    			key_deal(ch);
    		}
    	}
    	
    }

	endwin();
    key_close();
    printf("\033[?25h"); 
	exit(EXIT_SUCCESS);
}


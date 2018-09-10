package Clib

/*
//#include <windows.h>
//#include <conio.h>
//#include <curses.h>
// 使用了WinAPI来移动控制台的光标
//void gotoxy(int x,int y)
//{
//   COORD c;
//   c.X=x,c.Y=y;
//   SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
//   //printf("\033[%d;%dH",x,y)
//}
//
//// 从键盘获取一次按键，但不显示到控制台
//int direct()
//{
//   return _getch();
//}
////去掉控制台光标
//void hideCursor()
//{
//	CONSOLE_CURSOR_INFO  cci;
//	cci.bVisible = FALSE;
//	cci.dwSize = sizeof(cci);
//	SetConsoleCursorInfo(GetStdHandle(STD_OUTPUT_HANDLE), &cci);
//}
//#cgo CFLAGS : -I /usr/include
#cgo LDFLAGS: -L/usr/ -lcurses
#include <curses.h>
void initial()
{
    initscr();                    //开启curses模式

    cbreak();                     //开启cbreak模式，除 DELETE 或 CTRL 等仍被视为特殊控制字元外一切输入的字元将立刻被一一读取

    nonl();                       //用来决定当输入资料时，按下 RETURN 键是否被对应为 NEWLINE 字元

    noecho();                     //echo() and noecho(): 此函式用来控制从键盘输入字元时是否将字元显示在终端机上

    intrflush(stdscr,false);

    keypad(stdscr,true);          //当开启 keypad 後, 可以使用键盘上的一些特殊字元, 如上下左右>等方向键

    refresh();                    //将做清除萤幕的工作

}
void gotoxy(int x,int y)
{
   move(y,x);
}
int direct()
{
   return getch();
}
*/
import "C" // go中可以嵌入C语言的函数
func Cinit() {
	C.initscr()
}

//设置控制台光标位置
func GotoPostion(X int, Y int) {
	//调用C语言函数
	C.gotoxy(C.int(X), C.int(Y))
}

//无显获取键盘输入的字符
func Direction() (key int) {
	key = int(C.direct())
	return
}

////设置控制台光标隐藏
//func HideCursor() {
//	C.hideCursor()
//}

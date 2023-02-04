# GracefullyExit

优雅的退出go程序
**在终端执行的程序才会生效**

每个平台的信号定义或许有些不同。下面列出了POSIX中定义的信号。
Linux 使用34-64信号用作实时系统中。
命令man 7 signal提供了官方的信号介绍。

在POSIX.1-1990标准中定义的信号列表

<table>

<thead>

<tr>

<th style="text-align:left">信号</th>

<th style="text-align:left">值</th>

<th style="text-align:left">动作</th>

<th style="text-align:left">说明</th>

</tr>

</thead>

<tbody>

<tr>

<td style="text-align:left">SIGHUP</td>

<td style="text-align:left">1</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">终端控制进程结束(终端连接断开)</td>

</tr>

<tr>

<td style="text-align:left">SIGINT</td>

<td style="text-align:left">2</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">用户发送INTR字符(Ctrl+C)触发</td>

</tr>

<tr>

<td style="text-align:left">SIGQUIT</td>

<td style="text-align:left">3</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">用户发送QUIT字符(Ctrl+/)触发</td>

</tr>

<tr>

<td style="text-align:left">SIGILL</td>

<td style="text-align:left">4</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">非法指令(程序错误、试图执行数据段、栈溢出等)</td>

</tr>

<tr>

<td style="text-align:left">SIGABRT</td>

<td style="text-align:left">6</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">调用abort函数触发</td>

</tr>

<tr>

<td style="text-align:left">SIGFPE</td>

<td style="text-align:left">8</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">算术运行错误(浮点运算错误、除数为零等)</td>

</tr>

<tr>

<td style="text-align:left">SIGKILL</td>

<td style="text-align:left">9</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">无条件结束程序(不能被捕获、阻塞或忽略)</td>

</tr>

<tr>

<td style="text-align:left">SIGSEGV</td>

<td style="text-align:left">11</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)</td>

</tr>

<tr>

<td style="text-align:left">SIGPIPE</td>

<td style="text-align:left">13</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)</td>

</tr>

<tr>

<td style="text-align:left">SIGALRM</td>

<td style="text-align:left">14</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">时钟定时信号</td>

</tr>

<tr>

<td style="text-align:left">SIGTERM</td>

<td style="text-align:left">15</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">结束程序(可以被捕获、阻塞或忽略)</td>

</tr>

<tr>

<td style="text-align:left">SIGUSR1</td>

<td style="text-align:left">30,10,16</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">用户保留</td>

</tr>

<tr>

<td style="text-align:left">SIGUSR2</td>

<td style="text-align:left">31,12,17</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">用户保留</td>

</tr>

<tr>

<td style="text-align:left">SIGCHLD</td>

<td style="text-align:left">20,17,18</td>

<td style="text-align:left">Ign</td>

<td style="text-align:left">子进程结束(由父进程接收)</td>

</tr>

<tr>

<td style="text-align:left">SIGCONT</td>

<td style="text-align:left">19,18,25</td>

<td style="text-align:left">Cont</td>

<td style="text-align:left">继续执行已经停止的进程(不能被阻塞)</td>

</tr>

<tr>

<td style="text-align:left">SIGSTOP</td>

<td style="text-align:left">17,19,23</td>

<td style="text-align:left">Stop</td>

<td style="text-align:left">停止进程(不能被捕获、阻塞或忽略)</td>

</tr>

<tr>

<td style="text-align:left">SIGTSTP</td>

<td style="text-align:left">18,20,24</td>

<td style="text-align:left">Stop</td>

<td style="text-align:left">停止进程(可以被捕获、阻塞或忽略)</td>

</tr>

<tr>

<td style="text-align:left">SIGTTIN</td>

<td style="text-align:left">21,21,26</td>

<td style="text-align:left">Stop</td>

<td style="text-align:left">后台程序从终端中读取数据时触发</td>

</tr>

<tr>

<td style="text-align:left">SIGTTOU</td>

<td style="text-align:left">22,22,27</td>

<td style="text-align:left">Stop</td>

<td style="text-align:left">后台程序向终端中写数据时触发</td>

</tr>

</tbody>

</table>

**在SUSv2和POSIX.1-2001标准中的信号列表:**

<table>

<thead>

<tr>

<th style="text-align:left">信号</th>

<th style="text-align:left">值</th>

<th style="text-align:left">动作</th>

<th style="text-align:left">说明</th>

</tr>

</thead>

<tbody>

<tr>

<td style="text-align:left">SIGTRAP</td>

<td style="text-align:left">5</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">Trap指令触发(如断点，在调试器中使用)</td>

</tr>

<tr>

<td style="text-align:left">SIGBUS</td>

<td style="text-align:left">0,7,10</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">非法地址(内存地址对齐错误)</td>

</tr>

<tr>

<td style="text-align:left">SIGPOLL</td>

<td style="text-align:left"></td>

<td style="text-align:left">Term</td>

<td style="text-align:left">Pollable event (Sys V). Synonym for SIGIO</td>

</tr>

<tr>

<td style="text-align:left">SIGPROF</td>

<td style="text-align:left">27,27,29</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">性能时钟信号(包含系统调用时间和进程占用CPU的时间)</td>

</tr>

<tr>

<td style="text-align:left">SIGSYS</td>

<td style="text-align:left">12,31,12</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">无效的系统调用(SVr4)</td>

</tr>

<tr>

<td style="text-align:left">SIGURG</td>

<td style="text-align:left">16,23,21</td>

<td style="text-align:left">Ign</td>

<td style="text-align:left">有紧急数据到达Socket(4.2BSD)</td>

</tr>

<tr>

<td style="text-align:left">SIGVTALRM</td>

<td style="text-align:left">26,26,28</td>

<td style="text-align:left">Term</td>

<td style="text-align:left">虚拟时钟信号(进程占用CPU的时间)(4.2BSD)</td>

</tr>

<tr>

<td style="text-align:left">SIGXCPU</td>

<td style="text-align:left">24,24,30</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">超过CPU时间资源限制(4.2BSD)</td>

</tr>

<tr>

<td style="text-align:left">SIGXFSZ</td>

<td style="text-align:left">25,25,31</td>

<td style="text-align:left">Core</td>

<td style="text-align:left">超过文件大小资源限制(4.2BSD)</td>

</tr>

</tbody>

</table>

**kill pid 与 kill -9 pid的区别**

kill
pid的作用是向进程号为pid的进程发送SIGTERM（这是kill默认发送的信号），该信号是一个结束进程的信号且可以被应用程序捕获。若应用程序没有捕获并响应该信号的逻辑代码，则该信号的默认动作是kill掉进程。这是终止指定进程的推荐做法。

`kill -9 pid` 则是向进程号为pid的进程发送
SIGKILL（该信号的编号为9），从本文上面的说明可知，SIGKILL既不能被应用程序捕获，也不能被阻塞或忽略，其动作是立即结束指定进程。通俗地说，应用程序根本无法“感知”SIGKILL信号，它在完全无准备的情况下，就被收到SIGKILL信号的操作系统给干掉了，显然，在这种“暴力”情况下，应用程序完全没有释放当前占用资源的机会。事实上，SIGKILL信号是直接发给init进程的，它收到该信号后，负责终止pid指定的进程。在某些情况下（如进程已经hang死，无法响应正常信号），就可以使用 `kill -9`
来结束进程。

**应用程序如何优雅退出?**

Linux
Server端的应用程序经常会长时间运行，在运行过程中，可能申请了很多系统资源，也可能保存了很多状态，在这些场景下，我们希望进程在退出前，可以释放资源或将当前状态dump到磁盘上或打印一些重要的日志，也就是希望进程优雅退出（exit
gracefully）。

## Go中的Signal发送和处理

* golang中对信号的处理主要使用os/signal包中的两个方法：
* notify方法用来监听收到的信号
* stop方法用来取消监听

### 编写代码

- 1.监听全部信号
- 2.监听指定信号
- 3.优雅退出go守护进程

### 扩展

- [grace](https://github.com/facebookgo/grace)
# 异步编程

## P1：线程（Thread）：创建线程

### 什么是线程 Thread

- 线程是一个可执行路径，它可以独立其他线程执行
- 每个线程都在操作系统的进程（Process）内执行，而操作系统进程提供了程序运行的独立环境
- 单线程应用，在进程的独立环境里只跑一个线程，所以该线程拥有独占权
- 多线程应用，单个进程中会跑多个线程，它们会共享当前的执行环境（尤其是内存）
  - 例如，一个线程在后台读取数据，另一个线程在数据到达后进行展示。
  - 这个数据就被称作是共享的状态

### 线程的一些属性

- 线程一旦开始执行，IsAlive 就是 true，线程结束就变成 false。
- 线程结束的条件就是：线程构造函数传入的委托结束了执行。
- 线程一旦结束，就无法再重启。
- 每个线程都有个 Name 属性，通常用于调试。
  - 线程 Name 属性只能设置一次，以后更改会抛出异常。
- 静态的 Thread.CurrentThread 属性，会返回当前执行的线程

## P2:Join and Sleep

- 调用 Join 方法，就可以等待另一个线程结束。

### 添加超时

- 调用 Join 的时候，可以设置一个超时，用毫秒或者 TimeSpan 都可以。
  - 如果返回 true，那就是线程结束了；如果超时了，就返回 false。
- Thread.Sleep()方法会暂停当前的线程，并等一段时间。
- 注意：
  - Thread.Sleep(0) 这样调用会导致线程立即放弃本身当前的时间片，自动将 CPU 移交给其他线程。
  - Thread.Yield()做同样的事情，但是它只会把执行交给同一处理器上的其他线程
  - 当等待 Sleep 或 Join 的时候，线程处于阻塞的状态。

## P3:阻塞 Blocking

- 如果线程的执行由于某种原因导致暂停，那么就认为该线程被阻塞了。
  - 例如在 Sleep 或者通过 Join 等待其他线程结束。
- 被阻塞的线程会立即将其处理器的时间片生成给其他线程，从此就不再消耗处理器时间，直到满足其阻塞条件为止。
- 可以通过 ThreadState 这个属性来判断线程是否处于被阻塞的状态：
  ```C#
  bool block = (someThread.ThreadState & ThreadState.WaitSleepJoin) != 0
  ```

### ThreadState

- ThreadState 是一个 flags enum，通过按位的形式，可以合并数据的选项。

|                   |     |                                                                                              |
| ----------------- | --- | -------------------------------------------------------------------------------------------- |
| Aborted           | 256 | 线程状态包括 AbortRequested 并且该线程现在已死，但其状态尚未更改为 Stopped。                 |
| AbortRequest      | 128 | 已对线程调用了 Abort（Object）方法，但线程尚未收到试图终止它的挂起的 ThreadAbortException。  |
| Background        | 4   | 线程正作为后台程序执行（相对于前台线程而言）。此状态可以通过设置 IsBackground 属性来控制。   |
| Running           | 0   | 线程已启动尚未停止。                                                                         |
| Stopped           | 16  | 线程已停止。                                                                                 |
| StopRequested     | 1   | 正在请求线程停止，这仅用于内部。                                                             |
| Suspended         | 64  | 线程已挂起。                                                                                 |
| Suspended Request | 2   | 正在请求线程挂起                                                                             |
| Unstarted         | 8   | 尚未对线程调用 Start（）方法                                                                 |
| WaitSleepJoin     | 32  | 线程已被阻止，这可能是调用 Sleep（Int32）或 Join（），请求锁定或在线程同步对象上等待的结果。 |

- 但是它大部分的枚举值都没什么用，下面代码将 ThreadState 剥离为四个最有用的值之一：Unstarted、Running、WaitSleepJoin 和 Stopped
  ```C#
  public static ThreadState SimpleThreadState (ThreadState ts)
  {
      return ts & (ThreadState.Unstarted | ThreadState.WaitSleepJoin | ThreadState.Stopped);
  }
  ```
- 通常用于诊断

### 解除阻塞 Unblocking

- 当遇到下列四种情况的时候，就会解除阻塞：
  - 阻塞条件被满足
  - 操作超时（如果设置超时的话）
  - 通过 Thread.Interrupt()进行打断
  - 通过 Thread.About() 进行中止

### 上下文切换

- 当线程阻塞或解除阻塞时，操作系统将执行上下文切换。这会产生少量开销，通常为 1 或 2 微秒。

### I/O-bound vs Compute-bound(or CPU-Bound)

- 一个花费大部分时间等待某事发生的操作称为 I/O-bound
  - I/O 绑定操作通常涉及输入或输出，但这不是硬性需求：Thread.Sleep() 也被视为 I/O-bound
- 相反，一个花费大部分时间执行 CPU 密集型工作的操作称为 Compute-bound。

### 阻塞 vs 忙等待（自旋） Blocking vs Spinning

- IO-bound 操作的工作方式有两种：
  - 在当前线程上同步的等待
    - Console.ReadLine()、Thread.Sleep()、Thread.Join()...
  - 异步的操作，在稍后操作完成时触发一个回调动作。
- 同步等待的 I/O-bound 操作将大部分时间花在阻塞线程上。
- 它们也可以周期性的在一个循环里“打转（自旋）”
- 在忙等待和阻塞方面有一些细微差别
  - 首先，如果您希望条件会快得到满足（可能在几微秒之内），则短暂自旋可能会很有效，因为它避免了上下文切换的开销和延迟。
    - .Net Framework 提供了特殊的方法和类来提供帮助 SpinLock 和 SpinWait。
  - 其次，阻塞也不是零成本。这是因为每个线程在生存期间会占用大约 1MB 的内存，并会给 CLR 和操作系统带来持续的管理开销。
    - 因此，在需要处理成百上千个并发操作的大量 I/O-bound 程序的上下文中，阻塞可能会很麻烦
    - 所以，此类程序需要使用基于回调的方法，在等待时完全撤销其线程。

## P4:本地 vs 共享的状态 Local vs Shared State

### Local 本地独立

- CLR 为每个线程分配自己的内存栈（Stack），以便使本地变量保持独立。

### Shared 共享

- 如果多个线程都引用到同一个对象的实例，那么它们就共享里数据。
- 被 Lambda 表达式或匿名委托所捕获的本地变量，会被编译器转化为字段（field），所以也会被共享。
- 静态字段（field）也会在线程间共享数据。

### 线程安全 Thread Safety

### 锁定与线程安全 简介 Locking & Thread Safety

- 在读取和写入共享数据的时候，通过使用一个互斥锁（exclusive lock）
- C# 使用 lock 语句来加锁
- 当两个线程同时竞争一个锁的时候（锁可以基于任何引用类型对象），一个线程会等待或阻塞，直到锁变成可用状态。
- 在多线程上下文中，以这种方式避免不确定性的代码就叫做线程安全。
- Lock 不是线程安全的银弹，很容易忘记对字段加锁，lock 也会引起一些问题（死锁）

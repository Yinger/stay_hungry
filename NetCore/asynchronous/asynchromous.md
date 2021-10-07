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

## P5:向线程传递数据 & 异常处理

### 向线程传递数据

- 如果你想往线程的启动方法里传递参数，最简单的方式是使用 lambda 表达式，在里面使用参数调用方法。
- 甚至可以把整个逻辑都放在 lambda 里面。

#### Lambda 表达式与被捕获的变量

- 使用 Lambda 表达式可以很简单的给 Thread 传递参数。但是线程开始后，可能会不小心修改了被捕获的变量，这要多加注意。

#### 异常处理

- 创建线程时在作用范围内的 try/catch/finally 块，在线程开始执行后就与线程无关了。
- 在 WPF、WinForm 里，可以订阅全局异常处理事件：
  - Application.DispatcherUnhandledException
  - Application.ThreadException
  - 在通过消息循环调用的程序的任何部分发生未处理的异常后，将触发这些异常。
  - 但是非 UI 线程上的未处理异常，并不会触发它。
- 而任何线程有任何未处理的异常都会触发 AppDomain.CurrentDomain.UnhandledException

## P6:前台线程 vs 后台线程

- 默认情况下，你手动创建的线程就是前台线程。
- 只要有前台线程在运行，那么应用程序就会一直处于活动状态。
  - 但是后台线程却不行。
  - 一旦所有的前台线程停止，那么应用程序就停止了
  - 任何的后台线程也会突然终止
  - 注意：线程的前台、后台状态与它的优先级无关（所分配的执行时间）
- 可以通过 IsBackground 属性判断线程是否是后台线程。
- 进程以这种形式终止的时候，后台线程执行栈中的 finally 块就不会被执行了。
  - 如果想让它执行，可以在退出程序时使用 Join 来等待后台线程（如果是你自己创建的线程），或者使用 signal construct，如果是线程池...
- 应用程序无法正常退出的一个常见原因是还有活跃的前台线程

## P7：线程优先级

- 线程的优先级 （Thread 的 Priority 属性）决定了相对于操作系统中其他活跃线程所占的执行时间
- 优先级分为：
  - `enum ThreadPriority { Lowest, BelowNormal, Normal, AboveNormal, Highest}`

### 提升线程优先级

- 提升线程优先级的时候需特别注意，因为它可能“饿死”其他线程。
- 如果想让某线程（Thread）的优先级比其他进程（Process）中的线程（Thread）高，那就必须提升进程（Process）的优先级
  - 使用 System.Diagnositics 下的 Process 类。
    ```C#
    using (Process p = Process.GetCurrentProcess())
        p.PriorityClass = ProcessPriorityClass.High;
    ```
- 这可以很好地用于只做少量工作且需要较低延迟的非 UI 进程。
- 对于需要大量计算的应用程序（尤其是有 UI 的应用程序），提高进程优先级可能会使其他进程饿死，从而降低整个计算机的速度。

## P8:信号简介 Signaling

- 有时，你需要让某线程一直处于等待的状态，直至接收到其他线程发来的通知。这就叫做 signaling （发送信号）
- 最简单的信号结构就是 ManualResetEvent
  - 调用它上面的 WaitOne 方法会阻塞当前的线程，直到另一个线程通过调用 Set 方法来开启信号。

## P9:富客户端应用处理耗时操作的一种办法

在 WPF，UWP，WinForm 等类型的程序中，如果在主线程执行耗时的操作，就会导致整个程序无响应。因为主线程同时还需要处理消息循环，而渲染和鼠标键盘事件处理等工作都是消息循环来执行的。

- 针对这种耗时的操作，一种流行的做法是启用一个 worker 线程。
  - 执行完操作后，再更新到 UI
- 富客户端应用到线程模型通常是：
  - UI 元素和控件只能从创建它们的线程来进行访问（通常是主 UI 线程）
  - 当想从 worker 线程更新 UI 的时候，你必须把请求交给 UI 线程
- 比较底层的实现是：
  - WPF，在元素的 Dispatcher 对象上调用 BeginInvoke 或 Invoke。
  - WinForm，调用控件的 BeginInvoke 或 Invoke。
  - UWP，调用 Dispatcher 对象上的 RunAsync 或 Invoke
- 所有这些方法都接收一个委托。
- BeginInvoke 或 RunAsync 通过将委托排队到 UI 线程的消息队列来执行工作。
- Invoke 执行相同的操作，但随后会进行阻塞，直到 UI 线程读取并处理消息。
  - 因此，Invoke 允许您从方法中获取返回值。
  - 如果不需要返回值，BeginInvoke/RunAsync 更可取，因为它们不会阻塞调用方，也不会引入死锁的可能性

## P10:同步上下文 Synchronization Context

- 在 System.ComponentModel 下有一个抽象类：SynchronizationContext，它使得 Thread Marshaling 得到泛化。
- 针对移动、桌面（WPF，UWP，WinForms）等富客户端应用等 API，它们都定义和实例化了 Synchronization Context 的子类
  - 可以通过静态属性 SynchronizationContext.Current 来获得（当运行在 UI 线程时）
  - 捕获该属性让你可以在稍后等时候从 worker 线程向 UI 线程发送数据
  - 调用 Post 就相当于调用 Dispatch 或 Control 上面的 BeginInvoke 方法
  - 还有一个 Send 方法，它等价于 Invoke 方法

## P11:线程池 Thread Pool

- 当开始一个线程的时候，将花费几百微秒来组织类似以下的内容：
  - 一个新的局部变量栈（Stack）
- 线程池就可以节省这种开销：
  - 通过预先创建一个可循环使用线程的池来减少这一开销。
- 线程池对于高效的并行编程和细粒度并发是必不可少的
- 它允许在不被线程启动的开销淹没的情况下运行短期操作

### 使用线程池线程需要注意到几点：

- 不可以设置池线程的 Name
- 池线程都是后台线程
- 阻塞池线程可使性能降级

你可以自由的更改池线程的优先级

- 当它释放回池的时候优先级将还原为正常状态
  可以通过 Thread.CurrentThread.IsThreadPoolThread 属性来判断是否执行在池线程上

### 进入线程池

- 最简单的，显式的在池线程运行代码的方式就是使用 Task.Run

### 谁使用了线程池

- WCF、Remoting、ASP.NET、ASMX Web Services 应用服务器
- System.Timers.Timer、System.Threading.Timer
- 并行编程结构
- BackgroundWorker 类（现在很多余）
- 异步委托（现在很多余）

### 线程池中的整洁

- 线程池提供了另一个功能，即确保临时超出 计算-Bound 的工作不会导致 CPU 超额订阅
- CPU 超额订阅：活跃的线程超过 CPU 的核数，操作系统就需要对线程进行时间切片
- 超额订阅对性能影响很大，时间切片需要昂贵的上下文切换，并且可能使 CPU 缓存失效，而 CPU 缓存对于现代处理器的性能至关重要

#### CLR 的策略

- CLR 通过对任务排队并对其启动进行节流限制来避免线程池中的超额订阅。
- 它首先运行尽可能多的并发任务（只要还有 CPU 核），然后通过爬山算法调整并发级别，并在特定方向上不断调整工作负载。
  - 如果吞吐量提高，它将继续朝同一方向（否则将反转）
- 这确保它始终追随最佳性能曲线，即使面对计算机上竞争的进程活动时也是如此
- 如果下面两点能够满足，那么 CLR 的策略将发挥出最佳效果：
  - 工作项大多是短时间运行的（<250 毫秒，或者理想情况下<100 毫秒），因此 CLR 有很多机会进行测量和调整
  - 大部分时间都被阻塞的工作项不会主宰线程池

## P12:开始一个 Task

### Thread 的问题

- 线程（Thread）是用来创建并发（concurrency）的一种低级别工具，它有一些限制，尤其是：
  - 虽然开始线程的时候可以方便的传入数据，但是当 Join 的时候，很难从线程获得返回值。
    - 可能需要设置一些共享字段。
    - 如果操作抛出异常，捕获和传播该异常都很麻烦。
  - 无法告诉线程在结束时开始做另外的工作，你必须进行 Join 操作（在进程中阻塞当前的线程）
- 很难使用较小的并发（concurrent）来组建大型的并发。
- 导致了对手动同步的更大依赖以及随之而来的问题

### Task class

- Task 类可以很好的解决上述问题
- Task 是一个相对高级的抽象：它代表了一个并发操作（concurrent）
  - 该操作可能由 Thread 支持，或不由 Thread 支持
- Task 是可组合的（可使用 Continuation 把它们串成链）
  - Tasks 可以使用线程池来减少启动延迟
  - 使用 TaskCompletionSource，Tasks 可以利用回调的方式，在等待 I/O 绑定操作时完全避免线程。

### 开始一个 Task -- Task.Run

- Task 类在 System.Threading.Tasks 命名空间下。
- 开始一个 Task 最简单的办法就是使用 Task.Run （.NET 4.5,4.0 的时候是 Task.Factory.StartNew)这个静态方法：
  - 传入一个 Action 委托即可
- Task 默认使用线程池，也就是后台线程：
  - 当主线程结束时，你创建的所有 tasks 都会结束
- Task.Run 返回一个 Task 对象，可以使用它来监视其过程
  - 在 Task.Run 之后，我们没有调用 Start，因为该方法创建的是 “热” 任务 （hot task）
    - 可以通过 Task 的构造函数创建“冷”任务（cold task），但是很少这样做
- 可以通过 Task 的 Status 属性来跟踪 task 的执行状态。

### Wait 等待

- 调用 task 的 Wait 方法会进行阻塞直到操作完成
  - 相当于调用 thread 上的 Join 方法
- Wait 也可以让你指定一个超时时间和一个取消令牌来提前结束等待

### Long-running tasks 长时间运行的任务

- 默认情况下，CLR 在线程池中运行 Task，这非常适合短时间运行的 Compute-Bound 类工作
- 针对长时间运行的任务或者阻塞操作，你可以不采用线程池
- 如果同时运行多个 long-running tasks （尤其是其中有处于阻塞状态的），那么性能将会受很大影响，这时有比 TaskCreationOptions.LongRunning 更好的办法：
  - 如果任务是 IO-Bound，TaskCompletionSource 和异步函数可以让你用回调（Coninuations）代替线程来实现并发
  - 如果任务是 Compute-Bound，生产者/消费者队列允许你对任务的并发性进行限流，避免把其他线程和进程饿死。

## P13:Task 的返回值

## P14:Task 的异常

## P15:Coninuation

## P16:TaskCompletionSource

## P17:同步和异步

## P18:异步和 continuation 以及语言的支持

## P19:await

## P20:编写异步函数

## P21:异步中的同步上下文

## P22:优化：同步完成

## P23:ValueTask<T>

## P24：取消

## P25：进度报告

## P26：TAP

## P27:Task 组合器

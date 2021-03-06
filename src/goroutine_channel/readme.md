#### 并发和并行

并发： 
    - 立刻处理多个任务的能力(需要切换，从a切换到b开始处理，一次只能处理一个事物)
      a -> b -> a -> b
并行： 
    - 同时处理多个任务，如果运行在多核处理器上，多件事情在不同的内核上同时运行
      cpu1 a
      cpu2 b
    - 并行并不代表拥有更快的执行时间，因为有可能存在依赖关系 a依赖b
    
#### Goroutines（协程）
- Go 中，应用程序并发处理的部分被称作 goroutines（协程），它可以进行更有效的并发运算
- Go 使用 channels 来同步协程,通信
当 main() 函数返回的时候，程序退出：它不会等待任何其他非 main 协程的结束
- 启动一个新的协程，协程的调用会立即返回;和函数不同，主程序不会等待go协程执行完毕，
  在调用go协程后,程序控制会立即返回到代码的下一行，并且忽略该协程的任何返回值
- 如果希望运行其他go协程，go主协程必须继续运行，如果go主协程终止，则程序终止，其他go协程也不会继续运行
- 不要通过共享内存来通信，而通过通信来共享内存

#### channel
通道（channel），就像一个可以用于发送类型化数据的管道，由其负责协程之间的通信
- 可用于在其他协程结束执行之前，阻塞go主协程
- 先进先出（FIFO）的结构
- make() 函数来给它分配内存

#### 死锁
当 Go 协程给一个信道发送数据时，会有其他 Go 协程来接收数据。如果没有的话，程序就会在运行时触发 panic，形成死锁







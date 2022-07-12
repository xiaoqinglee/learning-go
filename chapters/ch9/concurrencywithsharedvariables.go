package ch9

/*
避免共享变量出现race condition的几类方法:
1.在main goroutine一次完成变量的初始化之后就不要再写该变量了, 所有的goroutine只读取该变量
2.不允许多个goroutine访问同一个变量
	(1)为变量设置监护人goroutine, 所有的读写操作均交由监护人goroutine处理
	(2)多goroutine直接读写共享变量的时候, 串行化进行, 不要出现interleaving
3.允许多个goroutine访问同一个变量, 但是使用互斥机制
	(1)使用sync.Mutex (或使用buffer size为1的buffered channel实现二元信号量), 注意防止死锁
	(2)使用sync.RWMutex
	(3)使用关联了某个mutex的sync.Cond
	(4)延迟初始化的场景使用sync.Once
4.使用并发安全的数据结构(如 sync.Map)和操作(如 atomic包)
*/

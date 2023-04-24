package main

// 1. 同一个package中，每个文件中，可以重复定义 init

// 2. 多文件中，依据文件名。文件中，依据书写顺序

// 3. 不相互依赖的package，依据 main 包的导入书写顺序

// 4. 有依赖关系的package，调用和依赖顺序相反，即：最后依赖，最先执行

// 5. 所有 init 函数都在同⼀个 goroutine 内执⾏，所有 init 函数结束后才会执⾏ main.main 函数。

// 6. package 执行  const -> var -> init

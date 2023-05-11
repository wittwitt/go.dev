package main

// body数据的操作
// bodyIO多数Method只操作一次，不存储body

// 如：先 c.GetRawData()

// 再 c.Getxxx，，就空了，，

// # bind

// ```golang
// binding:"required"`
// // int 不能为０
// // string 不能为""
// ```

// bindWithBody,,读取IO还，保存body reader的上下文

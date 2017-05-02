#How to use is 
func main() {
    urls := []string{"a", "b", "c"}
    pool := NewPool(20, len(urls)) // 初始化一个容量为20的并发控制池
    for _, v := range urls {
        go func(url string) {
            pool.AddOne() // 向并发控制池中添加一个, 一旦池满则此处阻塞
            err := Download(url)
            if nil != err {
                println(err)
            }
            pool.DelOne() // 从并发控制池中释放一个, 之后其他被阻塞的可以进入池中
        }(v)
    }
    pool.wg.Wait()  // 等待所有下载全部完成
}
func Download(s string) error {
    // do download logic
    println(s)
    return nil
}
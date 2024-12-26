func main() {
    var wg sync.WaitGroup
    ch := make(chan int, 5) // Buffered channel

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            val := <-ch
            fmt.Printf("Goroutine %d received: %d\n", i, val)
        }(i)
    }

    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
    wg.Wait()
} 
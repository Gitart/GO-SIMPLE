type Worker struct {
    tomb.Tomb
}

func (w *Worker) run() {
    defer w.Tomb.Done()
    for {
        switch v := <-conn; v {
        case 0:
            w.Tomb.Killf("error! got a zero")
            return
        case 1:
            return
        default:
            // continue
        }
    }
}

func main() {
    for i := 0; i < 10; i++ {
        var w Worker
        go w.run()
        fmt.Printf("Worker exited: %v\n", w.Wait())
    }

package main
import "fmt"


type Queue struct{
 	 items []int
}

func (q *Queue) Enqueu(i int) {
      q.items = append(q.items, i )
}

func (q *Queue) Dequeu() int {
	toRemove := q.items[0]
      q.items = q.items[:1]
      return toRemove
}

func main(){
    mq:=Queue{}
    fmt.Println(mq)
    mq.Enqueu(111)
    mq.Enqueu(200)
    mq.Enqueu(300)
    mq.Enqueu(350)
    mq.Enqueu(3770)
    mq.Enqueu(700)
    mq.Enqueu(800)


    fmt.Println(mq)
    mq.Dequeu()
    fmt.Println(mq)
}

package main
import(
	"fmt"
	"sync"
)
var (
	worker= []string{"worker1","worker2","worker3","worker4","worker5"}
	cycle=3
	wg sync.WaitGroup
)
func assemble(part string) {
	fmt.Println("woking on",part)
	fmt.Println("completed",part)
	wg.Done()
	// fmt.Println("assembling",part)
}

func main() {
	for i:=1;i<=cycle;i++ {
		fmt.Println("assembly cycle",i)
		wg.Add(len(worker))
		for _,part := range worker {
			go assemble(part)
		}
		wg.Wait()
		fmt.Println("assembly cycle",i,"completed")
	}
}
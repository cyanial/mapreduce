package mr

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

const (
	WORKER_IDLE = 0
	WORKER_INPROGRESS
	WORKER_COMPLETED
)

type WorkerHandler struct {
	id    int
	name  string
	state int
}

type Coordinator struct {
	// Your definitions here.
	nSplits int

	mu      sync.Mutex
	workers []*WorkerHandler
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

func (c *Coordinator) Register(args *RegisterArgs, reply *RegisterReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	worker := &WorkerHandler{}
	worker.id = len(c.workers)
	worker.name = args.WorkerName
	worker.state = WORKER_IDLE

	c.workers = append(c.workers, worker)
	reply.Id = worker.id

	fmt.Printf("Register: %s, %d\n", worker.name, worker.id)
	return nil
}

func (c *Coordinator) Unregister(args *ExampleArgs, reply *ExampleReply) error {
	return nil
}

func (c *Coordinator) FetchJob(args *FetchJobArgs, reply *FetchJobReply) error {
	//
	fmt.Println("someone is asking for a job (blocking)")
	wg := sync.WaitGroup{}
	wg.Add(1)

	wg.Wait()
	return nil
}

//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	// l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.

	return ret
}

//
// create a Coordinator.
// mian/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Read files and splits into small-chunks
	// just read-in and save
	// idx := 0
	// for _, filename := range files {
	// 	file, err := os.Open(filename)
	// 	if err != nil {
	// 		log.Fatalf("cannot open %v", filename)
	// 	}

	// 	oname := fmt.Sprintf("splits/mr-split-%d", idx)
	// 	ofile, _ := os.Create(oname)
	// 	io.Copy(ofile, file)
	// 	idx++
	// 	file.Close()
	// }

	// c.nSplits = idx

	fmt.Println("Start coordinator...")
	c.server()
	return &c
}

package mapreduce

import (
	"container/list"
	"fmt"
)

//
// schedule() starts and waits for all tasks in the given phase (Map
// or Reduce). the mapFiles argument holds the names of the files that
// are the inputs to the map phase, one per map task. nReduce is the
// number of reduce tasks. the registerChan argument yields a stream
// of registered workers; each item is the worker's RPC address,
// suitable for passing to call(). registerChan will yield all
// existing registered workers (if any) and new ones as they register.
//
func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var ntasks int
	var n_other int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, n_other)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	//
	fmt.Printf("Schedule: %v phase done\n", phase)
	workCh := make(chan taskResult)
	// 当前可用的worker
	worker := list.New()
	// 需要完成的任务
	pending := list.New()
	for i := 0; i < ntasks; i++ {
		pending.PushBack(i)
	}
LOOP:
	for {
		select {
		case srv := <-registerChan:
			if pending.Len() > 0 {
				front := pending.Front()
				pending.Remove(front)
				taskID := front.Value.(int)
				args := &DoTaskArgs{
					JobName:       jobName,
					File:          "",
					Phase:         phase,
					TaskNumber:    taskID,
					NumOtherPhase: n_other,
				}
				if phase == mapPhase {
					args.File = mapFiles[taskID]
				}
				go runTask(workCh, taskID, srv, args)
			} else {
				worker.PushBack(srv)
			}
		case res := <-workCh:
			if !res.Result {
				// 失败需要转移重试
				pending.PushBack(res.TaskID)
			} else {
				// 加入到worker中
				worker.PushBack(res.Service)
			}
			// 查看是否还有能执行的
			if pending.Len() == 0 {
				debug("stop schedule,%+v", phase)
				break LOOP
			}
			if worker.Len() > 0 {
				taskFront := pending.Front()
				pending.Remove(taskFront)
				workFront := worker.Front()
				worker.Remove(workFront)
				taskID := taskFront.Value.(int)
				srv := workFront.Value.(string)
				args := &DoTaskArgs{
					JobName:       jobName,
					File:          "",
					Phase:         phase,
					TaskNumber:    taskID,
					NumOtherPhase: n_other,
				}
				if phase == mapPhase {
					args.File = mapFiles[taskID]
				}
				go runTask(workCh, taskID, srv, args)
			}
		}
	}
}

type taskResult struct {
	Result  bool
	TaskID  int
	Service string
}

func runTask(workCh chan<- taskResult, taskID int, srv string, req *DoTaskArgs) {
	res := call(srv, "Worker.DoTask", req, nil)
	workCh <- taskResult{Result: res, TaskID: taskID, Service: srv}
}

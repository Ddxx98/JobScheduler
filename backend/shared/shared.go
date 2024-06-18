package shared

import (
    "container/heap"
    "sync"
    "github.com/gorilla/websocket"
    "backend/models"
)

// JobQueue is a priority queue implemented with a min-heap.
type JobQueue []*models.Job

func (jq JobQueue) Len() int { return len(jq) }

func (jq JobQueue) Less(i, j int) bool {
    return jq[i].Duration < jq[j].Duration // Shortest job first
}

func (jq JobQueue) Swap(i, j int) {
    jq[i], jq[j] = jq[j], jq[i]
}

func (jq *JobQueue) Push(x interface{}) {
    job := x.(*models.Job)
    *jq = append(*jq, job)
}

func (jq *JobQueue) Pop() interface{} {
    old := *jq
    n := len(old)
    job := old[n-1]
    *jq = old[0 : n-1]
    return job
}

// A thread-safe instance of the priority queue
var (
    JobQueueInstance = &JobQueue{}
    JobQueueMutex    = &sync.Mutex{}
    Clients          = make(map[*websocket.Conn]bool)
    Broadcast        = make(chan *models.Job)
)

func InitJobQueue() {
    heap.Init(JobQueueInstance)
}
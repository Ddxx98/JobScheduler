package scheduler

import (
    "container/heap"
    "time"
    "backend/models"
    "backend/shared"
)

func ScheduleJobs() {
    for {
        shared.JobQueueMutex.Lock()
        if shared.JobQueueInstance.Len() > 0 {
            job := heap.Pop(shared.JobQueueInstance).(*models.Job)
            job.Status = "running"
            shared.JobQueueMutex.Unlock()
            shared.Broadcast <- job  // Broadcast job status update
            time.Sleep(job.Duration * time.Second )
            job.Status = "completed"
            shared.Broadcast <- job  // Broadcast job completion update
        } else {
            shared.JobQueueMutex.Unlock()
        }
        time.Sleep(1 * time.Second)
    }
}

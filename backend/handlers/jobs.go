package handlers

import (
	"container/heap"
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/shared"
	"backend/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}

func SubmitJob(w http.ResponseWriter, r *http.Request)  {
    var job models.Job
    if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    //verify the name and duration
    if job.Name == "" || job.Duration == 0 {
        http.Error(w, "Name and duration are required", http.StatusBadRequest)
        return
    }
    job.ID = utils.GenerateID()
    job.Status = "pending" 
    shared.JobQueueMutex.Lock()
    heap.Push(shared.JobQueueInstance, &job)
    shared.JobQueueMutex.Unlock()
    // Broadcast new job to all clients
    shared.Broadcast <- &job  
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(job) 
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
    shared.JobQueueMutex.Lock()
    defer shared.JobQueueMutex.Unlock()
    jobs := make([]*models.Job, shared.JobQueueInstance.Len())
    copy(jobs, *shared.JobQueueInstance)
    json.NewEncoder(w).Encode(jobs)
}
package agent

import "sync"

var agentsPoolObject []WorkerAgent
var agentsPoolOnce sync.Once
var poolMutex sync.Mutex // Mutex to ensure thread-safe operations

// GetAgentsPool returns the singleton slice of WorkerAgents
func GetAgentsPool() *[]WorkerAgent {
	agentsPoolOnce.Do(func() {
		// Initialize the AgentsPool only once
		agentsPoolObject = make([]WorkerAgent, 0)
	})

	return &agentsPoolObject
}

// AppendToAgentsPool adds a WorkerAgent to the pool
func AppendToAgentsPool(agent WorkerAgent) {
	poolMutex.Lock() // Ensure thread-safety
	defer poolMutex.Unlock()

	agentsPoolPtr := GetAgentsPool()
	*agentsPoolPtr = append(*agentsPoolPtr, agent)
}

// DeleteFromAgentsPool removes a WorkerAgent from the pool by IP address
func DeleteFromAgentsPool(agentIP string) {
	poolMutex.Lock() // Ensure thread-safety
	defer poolMutex.Unlock()

	agentsPoolPtr := GetAgentsPool()
	for i, agent := range *agentsPoolPtr {
		if agent.IP == agentIP {
			// Remove the agent from the slice
			*agentsPoolPtr = append((*agentsPoolPtr)[:i], (*agentsPoolPtr)[i+1:]...)
			break
		}
	}
}

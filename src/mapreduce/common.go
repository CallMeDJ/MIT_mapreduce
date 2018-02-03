package mapreduce

import (
	"fmt"
	"strconv"
)

// Debugging enabled?
const debugEnabled = false

// debug() will only print if debugEnabled is true
func debug(format string, a ...interface{}) (n int, err error) {
	if debugEnabled {
		n, err = fmt.Printf(format, a...)
	}
	return
}

// jobPhase indicates whether a task is scheduled as a map or reduce task.
type jobPhase string

const projectUrl string = "/home/big/MIT_map_reduce_labor/6.824"
const (
	mapPhase    jobPhase = "Map"
	reducePhase          = "Reduce"
)

// KeyValue is a type used to hold the key/value pairs passed to the map and
// reduce functions.
type KeyValue struct {
	Key   string
	Value string
}

type KeyValueSlice []KeyValue
func (s KeyValueSlice) Len() int { return len(s) }
func (s KeyValueSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s KeyValueSlice) Less(i, j int) bool { return s[i].Key < s[j].Key }




// reduceName constructs the name of the intermediate file which map task
// <mapTask> produces for reduce task <reduceTask>.
func reduceName(jobName string, mapTask int, reduceTask int) string {
	return "mrtmp." + jobName + "-" + strconv.Itoa(mapTask) + "-" + strconv.Itoa(reduceTask)
}

// mergeName constructs the name of the output file of reduce task <reduceTask>
func mergeName(jobName string, reduceTask int) string {
	return "mrtmp." + jobName + "-res-" + strconv.Itoa(reduceTask)
}

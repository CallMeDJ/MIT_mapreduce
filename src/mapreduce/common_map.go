package mapreduce

import (
	"hash/fnv"
	"io/ioutil"
	"os"
	"strconv"
	"fmt"
	"encoding/json"
)

// doMap manages one map task: it reads one of the input files
// (inFile), calls the user-defined map function (mapF) for that file's
// contents, and partitions the output into nReduce intermediate files.
func doMap(
	jobName string, // the name of the MapReduce job
	mapTaskNumber int, // which map task this is
	inFile string,
	nReduce int, // the number of reduce task that will be run ("R" in the paper)
	mapF func(file string, contents string) []KeyValue,
) {
	//
	// You will need to write this function.
	//
	// The intermediate output of a map task is stored as multiple
	// files, one per destination reduce task. The file name includes
	// both the map task number and the reduce task number. Use the
	// filename generated by reduceName(jobName, mapTaskNumber, r) as
	// the intermediate file for reduce task r. Call ihash() (see below)
	// on each key, mod nReduce, to pick r for a key/value pair.
	//
	// mapF() is the map function provided by the application. The first
	// argument should be the input file name, though the map function
	// typically ignores it. The second argument should be the entire
	// input file contents. mapF() returns a slice containing the
	// key/value pairs for reduce; see common.go for the definition of
	// KeyValue.
	//
	// Look at Go's ioutil and os packages for functions to read
	// and write files.
	//
	// Coming up with a scheme for how to format the key/value pairs on
	// disk can be tricky, especially when taking into account that both
	// keys and values could contain newlines, quotes, and any other
	// character you can think of.
	//
	// One format often used for serializing data to a byte stream that the
	// other end can correctly reconstruct is JSON. You are not required to
	// use JSON, but as the output of the reduce tasks *must* be JSON,
	// familiarizing yourself with it here may prove useful. You can write
	// out a data structure as a JSON string to a file using the commented
	// code below. The corresponding decoding functions can be found in
	// common_reduce.go.
	//
	//   enc := json.NewEncoder(file)
	//   for _, kv := ... {
	//     err := enc.Encode(&kv)
	//
	// Remember to close the file after you have written all the values!
	//

	valuesInFile,err := ioutil.ReadFile(projectUrl+"/src/main/"+inFile)
	if err != nil{
		fmt.Println("open file error "+inFile+"\n"+err.Error());

	}
	var keyValues= mapF(inFile,string(valuesInFile[:]))

	var commonFilePrefix string = "mrtmp." + jobName + "-" + strconv.Itoa(mapTaskNumber)
	for index:=0;index<nReduce;index++ {
		var fileName string = commonFilePrefix + "-" + strconv.Itoa(index);
		fmt.Println("creating file "+fileName)
		os.Create(projectUrl+"/src/main/"+fileName);
	}


	for _ , value := range keyValues{
		var whichReduce = ihash(value.Key) % nReduce;
		var currentFileName = projectUrl+"/src/main/"+commonFilePrefix+"-"+strconv.Itoa(whichReduce);
		//fmt.Println("current file is "+currentFileName);
		currentFile , error := os.OpenFile(currentFileName,os.O_WRONLY|os.O_APPEND, 0666)
		if error != nil{
			fmt.Println("open file :"+currentFileName +"fail."+error.Error() )
		}

		enc := json.NewEncoder(currentFile)
		enc.Encode(&value)
		//currentFile.WriteString(string(jsonResult));
		currentFile.Close()
	}

}

func ihash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() & 0x7fffffff)
}

package mapreduce

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// doReduce manages one reduce task: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	outFile string, // write the output here
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	//
	// You will need to write this function.
	//
	// You'll need to read one intermediate file from each map task;
	// reduceName(jobName, m, reduceTaskNumber) yields the file
	// name from map task m.
	//
	// Your doMap() encoded the key/value pairs in the intermediate
	// files, so you will need to decode them. If you used JSON, you can
	// read and decode by creating a decoder and repeatedly calling
	// .Decode(&kv) on it until it returns an error.
	//
	// You may find the first example in the golang sort package
	// documentation useful.
	//
	// reduceF() is the application's reduce function. You should
	// call it once per distinct key, with a slice of all the values
	// for that key. reduceF() returns the reduced value for that key.
	//
	// You should write the reduce output as JSON encoded KeyValue
	// objects to the file named outFile. We require you to use JSON
	// because that is what the merger than combines the output
	// from all the reduce tasks expects. There is nothing special about
	// JSON -- it is just the marshalling format we chose to use. Your
	// output code will look something like this:
	//
	// enc := json.NewEncoder(file)
	// for key := ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//

	// 这其实是shuffle的过程
	// shuffle包括几个步骤:
	// 1:数据迁移,从map节点将数据迁移到reduce节点
	// 2:数据合并,将相同key的数据合并成一个数组
	// 3:数据排序,按照key进行排序,这个貌似不是必须的
	kvMap := make(map[string][]string)
	for i := 0; i < nMap; i++ {
		fileName := reduceName(jobName, i, reduceTaskNumber)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return
		}
		kvList := make([]KeyValue, 0)
		if err := json.Unmarshal(data, &kvList); err != nil {
			log.Printf("Unmarshal fail,%+v", err)
			return
		}

		for _, kv := range kvList {
			kvMap[kv.Key] = append(kvMap[kv.Key], kv.Value)
		}
	}

	// 执行reduce
	file, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Printf("open output file fail,%+v", err)
		return
	}
	defer file.Close()
	debug("write output file,%+v", outFile)

	enc := json.NewEncoder(file)
	//
	for key, value := range kvMap {
		if err := enc.Encode(KeyValue{key, reduceF(key, value)}); err != nil {
			log.Printf("encode fail,%+v", err)
			return
		}
	}
}

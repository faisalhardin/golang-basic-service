package src

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"task1/entity"
)

func ListFiles(fileDir string) (fileNames []string, err error) {
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		return fileNames, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}

	}

	return fileNames, nil
}

func ReadFiles(filePath string) (transactionLogs []entity.Transaction, err error) {
	filePath = "./subsetdata/2022-11-10-1668042139551695521.ndjson"
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		trxLine := sc.Bytes()
		transaction, e := ConvertToStruct(trxLine)
		if e != nil {
			return
		}

		transactionLogs = append(transactionLogs, transaction)
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	return transactionLogs, err
}

func ConvertToStruct(line []byte) (transaction entity.Transaction, err error) {

	err = json.Unmarshal(line, &transaction)
	log.Default().Print(transaction)
	return transaction, err

}

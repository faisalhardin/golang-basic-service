package filereader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"task1/entity"
	mockrepo "task1/entity/mock"
)

var (
	ioutilReadDir   = ioutil.ReadDir
	osOpenFile      = os.OpenFile
	bufioNewScanner = BufioNewScanner
)

func ListFiles(fileDir string) (fileNames []string, err error) {
	files, err := ioutilReadDir(fileDir)
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

func ReadFilesWithChannel(prefix string, filePaths []string) <-chan entity.StockCodeToTransactionLogKeyValue {
	record := make(chan entity.StockCodeToTransactionLogKeyValue)
	go func() {
		for _, filePath := range filePaths {
			f, err := osOpenFile(fmt.Sprintf("%s/%s", prefix, filePath), os.O_RDONLY, os.ModePerm)
			if err != nil {
				log.Fatalf("open file error: %v", err)
			}
			defer f.Close()
			sc := bufioNewScanner(f)
			for sc.Scan() {
				trxLine := sc.Bytes()
				transaction, e := ConvertToStruct(trxLine)
				if e != nil {
					log.Fatalf("open file error: %v", e)
				}

				stockCodeToTransactionLogKeyValue := entity.StockCodeToTransactionLogKeyValue{
					StockCode:      transaction.Stock,
					TransactionLog: string(trxLine),
				}
				record <- stockCodeToTransactionLogKeyValue
			}
			if err := sc.Err(); err != nil {
				log.Fatalf("scan file error: %v", err)
			}
		}
		close(record)
	}()
	return record
}

func ConvertToStruct(line []byte) (transaction entity.Transaction, err error) {
	err = json.Unmarshal(line, &transaction)
	return transaction, err
}

func BufioNewScanner(r io.Reader) mockrepo.NewScanner {
	return bufio.NewScanner(r)
}

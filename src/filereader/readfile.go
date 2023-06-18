package filereader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"task1/entity"
	"task1/src/repo"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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

func ReadFilesAndSendMessage(prefix string, filePath string, msgKafka *repo.KafkaOption) (transactionLogs []entity.Transaction, err error) {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", prefix, filePath), os.O_RDONLY, os.ModePerm)
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

		msgKafka.SendMessage(&kafka.Message{
			Key:   []byte(transaction.Stock),
			Value: trxLine,
		})
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	return transactionLogs, err
}

func ReadFilesWithChannel(prefix string, filePaths []string) <-chan entity.StockCodeToTransactionLogKeyValue {
	record := make(chan entity.StockCodeToTransactionLogKeyValue)
	go func() {
		for _, filePath := range filePaths {
			f, err := os.OpenFile(fmt.Sprintf("%s/%s", prefix, filePath), os.O_RDONLY, os.ModePerm)
			if err != nil {
				log.Fatalf("open file error: %v", err)
			}
			defer f.Close()

			sc := bufio.NewScanner(f)
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

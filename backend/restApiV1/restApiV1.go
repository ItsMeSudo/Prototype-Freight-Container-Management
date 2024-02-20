package restApiV1

import (
	"backend/crossFunc"
	"backend/initFlag"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetBlockInfoPost(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(initFlag.JsonServerFullPath + "/containers")
	if err != nil {
		http.Error(w, "Error fetching containers from API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&crossFunc.Containers); err != nil {
		http.Error(w, "Error decoding container data", http.StatusInternalServerError)
		return
	}

	blockContainers := make(map[int][]crossFunc.Container)
	for _, container := range crossFunc.Containers {
		blockContainers[container.BlockID] = append(blockContainers[container.BlockID], container)
	}

	var result []map[string]interface{}
	for blockID, block := range blockContainers {
		capacity := float64(len(block)) / float64(5*5*5) * 100
		averageAge, oldestID, newestID := crossFunc.CalculateStatistics(block)
		emptyPositions, emptyBays, emptyStacks := crossFunc.CalculateEmptyPositions(block)

		blockData := map[string]interface{}{
			"blockId":           blockID,
			"capacity":          fmt.Sprintf("%.2f", capacity),
			"averageAge":        fmt.Sprintf("%.2f", averageAge),
			"oldestContainerId": oldestID,
			"newestContainerId": newestID,
			"emptyPositions":    emptyPositions,
			"emptyBays":         emptyBays,
			"emptyStacks":       emptyStacks,
		}
		result = append(result, blockData)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func InsertJsonData(w http.ResponseWriter, r *http.Request) {
	crossFunc.RefreshContainers()
	var newContainers []crossFunc.Container
	if err := json.NewDecoder(r.Body).Decode(&newContainers); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	successCount := 0
	incorrectPositions := make([]string, 0)

	for _, newContainer := range newContainers {
		if crossFunc.IsDuplicate(newContainer) || crossFunc.HasConflict(newContainer) {
			incorrectPositions = append(incorrectPositions, newContainer.ID)
		} else {
			if err := crossFunc.InsertContainer(newContainer); err != nil {
				log.Println("Error inserting container:", err)
				continue
			}
			successCount++
		}
	}

	response := struct {
		Success            int      `json:"success"`
		IncorrectPositions []string `json:"incorrectPositions"`
	}{
		Success:            successCount,
		IncorrectPositions: incorrectPositions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleCSVBinaryUpload(w http.ResponseWriter, r *http.Request) {
	crossFunc.RefreshContainers()
	csvData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	reader := csv.NewReader(bytes.NewReader(csvData))
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Failed to parse CSV data", http.StatusBadRequest)
		return
	}

	successCount := 0
	incorrectPositions := make([]string, 0)

	for _, record := range records {
		container := crossFunc.Container{
			ID:        record[0],
			BlockID:   crossFunc.RecordInt(record[1]),
			BayNum:    crossFunc.RecordInt(record[2]),
			StackNum:  crossFunc.RecordInt(record[3]),
			TierNum:   crossFunc.RecordInt(record[4]),
			ArrivedAt: time.Unix(0, crossFunc.RecordInt64(record[5])*int64(time.Millisecond)),
		}

		if container.BlockID == 0 {
			continue
		}
		if crossFunc.IsDuplicate(container) || crossFunc.HasConflict(container) {
			incorrectPositions = append(incorrectPositions, container.ID)
		} else {
			if err := crossFunc.InsertContainer(container); err != nil {
				log.Println("Error inserting container:", err)
				continue
			}
			successCount++
		}
	}

	response := struct {
		Success            int      `json:"success"`
		IncorrectPositions []string `json:"incorrectPositions"`
	}{
		Success:            successCount,
		IncorrectPositions: incorrectPositions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

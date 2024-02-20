package restApiV2

import (
	"backend/crossFunc"
	"backend/initFlag"
	"encoding/csv"
	"encoding/json"
	"fmt"
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

func GetAllContainers(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(crossFunc.Containers)
}

func InsertJsonFileData(w http.ResponseWriter, r *http.Request) {
	crossFunc.RefreshContainers()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var newContainers []crossFunc.Container
	if err := json.NewDecoder(file).Decode(&newContainers); err != nil {
		http.Error(w, "Error decoding JSON data from file", http.StatusBadRequest)
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

func HandleCSVFormUpload(w http.ResponseWriter, r *http.Request) {
	crossFunc.RefreshContainers()
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
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

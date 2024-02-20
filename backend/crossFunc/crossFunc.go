package crossFunc

import (
	"backend/initFlag"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"
)

type Container struct {
	ID        string    `json:"id"`
	BlockID   int       `json:"blockId"`
	BayNum    int       `json:"bayNum"`
	StackNum  int       `json:"stackNum"`
	TierNum   int       `json:"tierNum"`
	ArrivedAt time.Time `json:"arrivedAt"`
}

var Containers []Container

func CalculateStatistics(containers []Container) (float64, string, string) {
	var sumAge float64
	var oldestDate, newestDate time.Time
	var oldestID, newestID string

	if len(containers) == 0 {
		return 0, "", ""
	}

	oldestDate = containers[0].ArrivedAt
	newestDate = containers[0].ArrivedAt
	oldestID = containers[0].ID
	newestID = containers[0].ID

	for _, container := range containers {
		sumAge += time.Since(container.ArrivedAt).Hours() / 24
		if container.ArrivedAt.Before(oldestDate) {
			oldestDate = container.ArrivedAt
			oldestID = container.ID
		}
		if container.ArrivedAt.After(newestDate) {
			newestDate = container.ArrivedAt
			newestID = container.ID
		}
	}

	averageAge := sumAge / float64(len(containers))
	return averageAge, oldestID, newestID
}

func CalculateEmptyPositions(containers []Container) (int, int, int) {
	var emptyPositions, emptyBays, emptyStacks int

	blockSize := 5 * 5 * 5

	sort.SliceStable(containers, func(i, j int) bool {
		if containers[i].BayNum == containers[j].BayNum {
			return containers[i].StackNum < containers[j].StackNum
		}
		return containers[i].BayNum < containers[j].BayNum
	})

	currentBay := containers[0].BayNum
	currentStack := containers[0].StackNum

	for _, container := range containers {
		if container.BayNum != currentBay {
			emptyBays++
			currentBay = container.BayNum
		}
		if container.StackNum != currentStack {
			emptyStacks++
			currentStack = container.StackNum
		}
	}

	emptyPositions = blockSize - len(containers)
	return emptyPositions, emptyBays, emptyStacks
}

func RefreshContainers() {
	resp, err := http.Get(initFlag.JsonServerFullPath + "/containers")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&Containers); err != nil {
		log.Println(err)
		return
	}
}

func IsDuplicate(newContainer Container) bool {
	for _, c := range Containers {
		if c.ID == newContainer.ID {
			return true
		}
	}
	return false
}

func HasConflict(newContainer Container) bool {
	for _, c := range Containers {
		if c.BlockID == newContainer.BlockID &&
			c.BayNum == newContainer.BayNum &&
			c.StackNum == newContainer.StackNum &&
			c.TierNum == newContainer.TierNum {
			return true
		}
	}
	return false
}

func InsertContainer(newContainer Container) error {
	data, err := json.Marshal(newContainer)
	if err != nil {
		return err
	}

	resp, err := http.Post(initFlag.JsonServerFullPath+"/containers", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API returned non-OK status code: %d", resp.StatusCode)
	}

	return nil
}

func RecordInt(s string) int {
	var result int
	if _, err := fmt.Sscanf(s, "%d", &result); err != nil {
		result = 0
	}
	return result
}

func RecordInt64(s string) int64 {
	var result int64
	if _, err := fmt.Sscanf(s, "%d", &result); err != nil {
		result = 0
	}
	return result
}

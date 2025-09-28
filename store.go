package main

// store.go berisi fungsi untuk memuat dan menyimpan tugas ke file JSON
import (
	"encoding/json"
	"errors"
	"os"
	"sort"
)

const dataFile = "tasks.json"
// ensureDataFile memastikan bahwa file data ada, jika tidak ada, buat file baru
func ensureDataFile() error {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return saveTasks([]Task{})
	}
	return nil
}

// loadTasks memuat tugas dari file JSON
func loadTasks() ([]Task, error) {
	if err := ensureDataFile(); err != nil {
		return nil, err
	} // ini memastikan file ada sebelum mencoba membacanya
	b, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	if len(b) == 0 {
		return tasks, nil // kembalikan slice kosong jika file kosong
	}
	if err := json.Unmarshal(b, &tasks); err != nil {
		return []Task{}, nil // kembalikan slice kosong jika ada kesalahan unmarshalling
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})
	return tasks, nil
}

// saveTasks menyimpan tugas ke file JSON
func saveTasks(tasks []Task) error {
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	tmp := dataFile + ".tmp"
	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, dataFile)
}

// nextID mengembalikan ID berikutnya untuk tugas baru
func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1 // return max + 1 sebagai ID berikutnya
}

// findTaskIndex mencari indeks tugas berdasarkan ID
func findTaskIndex(tasks []Task, id int) (int, error) {
	for i, t := range tasks {
		if t.ID == id {
			return i, nil // return index task jika ditemukan, dan nil untuk error
		}
	}
	return -1, errors.New("task not found") // return -1 dan error jika task tidak ditemukan
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const taskFile = "tasks.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: todo add <タスク> | todo list | todo remove <番号>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("タスクを指定してください")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		addTask(task)
	case "list":
		listTasks()
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("削除するタスクの番号を指定してください")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil || index < 1 {
			fmt.Println("正しい番号を入力してください")
			return
		}
		removeTask(index)
	default:
		fmt.Println("不明なコマンド:", command)
	}
}

func addTask(task string) {
	file, err := os.OpenFile(taskFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ファイルを開けません:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(task + "\n")
	if err != nil {
		fmt.Println("タスクを保存できません:", err)
		return
	}

	fmt.Println("タスク追加:", task)
}

func listTasks() {
	file, err := os.Open(taskFile)
	if err != nil {
		fmt.Println("タスクファイルが存在しません")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("タスクリスト:")
	i := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", i, scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("タスクを読み込めません:", err)
	}
}

func removeTask(index int) {
	file, err := os.Open(taskFile)
	if err != nil {
		fmt.Println("タスクファイルが存在しません")
		return
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if index > len(tasks) {
		fmt.Println("指定された番号のタスクは存在しません")
		return
	}

	tasks = append(tasks[:index-1], tasks[index:]...)

	file, err = os.Create(taskFile)
	if err != nil {
		fmt.Println("タスクファイルを更新できません:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		_, err := file.WriteString(task + "\n")
		if err != nil {
			fmt.Println("タスクを保存できません:", err)
			return
		}
	}

	fmt.Printf("タスク %d を削除しました\n", index)
}
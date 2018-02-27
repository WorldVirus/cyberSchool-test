package main

import (
	"os"
	"encoding/csv"
	"strconv"
	"sort"
	"fmt"
)

type CsvLine struct {
	Column2 []int
	Column3 []int
	AnswerUser []string
	AnswerIp []string
}

var data CsvLine
var mpUser = make(map[string]int)
var mpIp = make(map[string]int)

func main() {
	f, err := os.Open("shkib.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("You have to wait compilation programm, becuause there is not the best algorithms. It is killing golang")

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	
	defer f.Close()

	i:=0
	j:=0
	for _, line := range lines {

		i1, _ := strconv.Atoi(line[6])
		i2, _ := strconv.Atoi(line[7])

			data.Column2 = append(data.Column2, i1)
			data.Column3 = append(data.Column3, i2)

			searcherIpUser(mpUser,line[1],&i,&data.AnswerUser)
			searcherIpUser(mpIp,line[2],&j,&data.AnswerIp)

		}

	sort.Ints(data.Column3)
	sort.Ints(data.Column2)

	Creation()
	firstTask()
	fmt.Println("Yeah ! You are waited for )")
}


func Creation()  {
	var file, err = os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

}


func firstTask()  {

	var file, error = os.OpenFile("result.txt", os.O_RDWR, 0644)
	if error != nil {
		panic(error)
	}
	defer file.Close()
	file.WriteString("#Поиск 5ти пользователей, сгенерировавших наибольшее количество запросов\nРешение1\n")

	for i:=1;i<6;i++{
		input:=strconv.Itoa(data.Column3[len(data.Column3)-i])
		file.WriteString(input+" ")
	}
	file.WriteString("\n")
	file.WriteString("#Поиск 5ти пользователей, отправивших наибольшее количество данных\nРешение2\n")

	for i:=1;i<6;i++{
		input:=strconv.Itoa(data.Column2[len(data.Column2)-i])
		file.WriteString(input+" ")
	}
	file.WriteString("\n")

	file.WriteString("\n")
	file.WriteString("#Поиск регулярных запросов (запросов выполняющихся периодически) по полю src_user\nРешение3\n")
	for i:=0;i<len(data.AnswerUser)-1;i++{
		file.WriteString(data.AnswerUser[i]+" ")
	}
	file.WriteString("\n"+"\n")

	file.WriteString("#Поиск регулярных запросов (запросов выполняющихся периодически) по полю src_ip\nРешение4\n")
	for i:=0;i<len(data.AnswerIp)-1;i++{
		file.WriteString(data.AnswerIp[i]+" ")
	}
	file.WriteString("\n")

	file.Sync()
}

func searcherIpUser(mpUser map[string]int,hash string,i *int, answer *[] string)  {
	if mpUser[hash] == 0 {
		*i+=1
		mpUser[hash] = *i
	} else {
		if len(*answer) == 0 {
			*answer = append(*answer, hash)
		} else {
			flag := true
			for _,j:=range *answer{
				if j == hash {
						flag = false
						break
					}
				}
			if flag {
				*answer = append(*answer,hash)
			}
			*i+=1
		}
	}
}
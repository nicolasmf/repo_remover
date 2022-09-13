package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetRepoNames(str string, start string, end1 string, end2 string, end3 string) (result string) {
	name := strings.Split(str, "\n")[1]

	s := strings.Index(name, start)
	if s == -1 {
		return
	}

	s += len(start)

	e := strings.Index(name[s:], end1)

	if e == -1 {
		end := strings.Index(name[s:], end2)
		if end == -1 {
			return name[s:]
		} else {
			e = end
		}
	} else {
		end := strings.Index(name[s:], end2)
		if end != -1 && end < e {
			e = end
		}
	}
	return name[s : s+e]
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func DeleteRepo(repoName string) {
	files, err := os.ReadDir("/etc/yum.repos.d/")
	Check(err)

	for _, f := range files {
		data, err := os.ReadFile("/etc/yum.repos.d/" + f.Name())
		Check(err)
		data_string := string(data)
		result := GetRepoNames(data_string, "=", " -", " $", "\n")
		if result == repoName {
			err := os.Remove("/etc/yum.repos.d/" + f.Name())
			Check(err)
			fmt.Println("Deleted " + f.Name())
		}
	}
}

func main() {
	files, err := os.ReadDir("/etc/yum.repos.d/")
	Check(err)

	var names []string

	for _, f := range files {
		data, err := os.ReadFile("/etc/yum.repos.d/" + f.Name())
		Check(err)
		data_string := string(data)
		result := GetRepoNames(data_string, "=", " -", " $", "\n")
		names = append(names, result)
	}

	names = removeDuplicateStr(names)

	var question = []*survey.Question{
		{
			Name: "Repository",
			Prompt: &survey.Select{
				Message: "Select a Repository to delete :",
				Options: names,
			},
		},
	}

	answers := struct {
		RepoToDelete string `survey:"Repository"`
	}{}

	ans_error := survey.Ask(question, &answers)
	Check(ans_error)

	DeleteRepo(answers.RepoToDelete)

}

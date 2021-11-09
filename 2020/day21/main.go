package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

func parseIngredientList(in string) map[string][][]string {
	r := strings.NewReader(in)
	s := bufio.NewScanner(r)

	out := map[string][][]string{}

	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, "(")

		encryptedItemsStr := strings.TrimSpace(parts[0])
		encryptedItems := strings.Split(encryptedItemsStr, " ")

		ingredientListStr := parts[1]
		ingredientListStr = strings.ReplaceAll(ingredientListStr, "contains", "")
		ingredientListStr = strings.ReplaceAll(ingredientListStr, ")", "")
		ingredientListStr = strings.TrimSpace(ingredientListStr)
		ingredientListStr = strings.ReplaceAll(ingredientListStr, " ", "")
		ingredientList := strings.Split(ingredientListStr, ",")
		for i := range ingredientList {
			ingredientList[i] = strings.TrimSpace(ingredientList[i])
		}

		for _, ingredient := range ingredientList {
			if _, ok := out[ingredient]; !ok {
				out[ingredient] = [][]string{}
			}
			out[ingredient] = append(out[ingredient], encryptedItems)
		}
	}

	return out
}

func ingredientListToCandidates(ingredients map[string][][]string) map[string][]string {
	out := map[string][]string{}
	for ingredient, encryptedLists := range ingredients {
		seen := map[string]int{}
		for i := range encryptedLists {
			for j := range encryptedLists[i] {
				seen[encryptedLists[i][j]]++
			}
		}
		out[ingredient] = []string{}
		for k, v := range seen {
			if v == len(ingredients[ingredient]) {
				out[ingredient] = append(out[ingredient], k)
			}
		}
	}
	return out
}

func getItemsLines(in string) [][]string {
	r := strings.NewReader(in)
	s := bufio.NewScanner(r)

	out := [][]string{}

	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, "(")
		encryptedItemsStr := strings.TrimSpace(parts[0])
		encryptedItems := strings.Split(encryptedItemsStr, " ")
		out = append(out, encryptedItems)
	}

	return out
}

func removeElement(in []string, toRemove string) []string {
	var out []string
	for _, s := range in {
		if s != toRemove {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	partOneFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		ingredientList := parseIngredientList(string(in))

		candidates := ingredientListToCandidates(ingredientList)
		itemLines := getItemsLines(string(in))

		allEncrypted := map[string]int{}
		for _, itemLine := range itemLines {
			for i := range itemLine {
				allEncrypted[itemLine[i]]++
			}
		}

		var sum int
	outer:
		for encrypted := range allEncrypted {
			for _, items := range candidates {
				for i := range items {
					if encrypted == items[i] {
						continue outer
					}
				}
			}
			sum += allEncrypted[encrypted]
		}

		fmt.Println(sum)

		return nil
	}

	partTwoFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		ingredientList := parseIngredientList(string(in))

		candidates := ingredientListToCandidates(ingredientList)
		itemLines := getItemsLines(string(in))

		allEncrypted := map[string]int{}
		for _, itemLine := range itemLines {
			for i := range itemLine {
				allEncrypted[itemLine[i]]++
			}
		}

		inert := []string{}

	outer:
		for encrypted := range allEncrypted {
			for _, items := range candidates {
				for i := range items {
					if encrypted == items[i] {
						continue outer
					}
				}
			}
			inert = append(inert, encrypted)
		}

	allFound:
		for {
			for i := range candidates {
				if len(candidates[i]) > 1 {
					break
				}
				break allFound
			}

			for i := range candidates {
				if len(candidates[i]) == 1 {
					item := candidates[i][0]
					for j := range candidates {
						if i == j {
							continue
						}
						candidates[j] = removeElement(candidates[j], item)
					}
				}
			}
		}

		keys := []string{}
		for i := range candidates {
			keys = append(keys, i)
		}

		sort.Strings(keys)
		sb := &strings.Builder{}
		for i := range keys {
			if i != 0 {
				sb.WriteRune(',')
			}
			sb.WriteString(candidates[keys[i]][0])
		}

		result := sb.String()
		fmt.Println(result)

		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}

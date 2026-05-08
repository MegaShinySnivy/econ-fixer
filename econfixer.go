package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	inputFile := "Sandbox.sbc"
	outputFile := "Sandbox_noNPCs.sbc"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	content := string(data)

	factionBlockRe := regexp.MustCompile(`(?s)<MyObjectBuilder_Faction>(.*?)</MyObjectBuilder_Faction>`)
	tagValueRe      := regexp.MustCompile(`<Tag>(.*?)</Tag>`)

	var removalStartIndices []int 
	for _, m := range factionBlockRe.FindAllStringIndex(content, -1) {
		blockBytes := data[m[0]:m[1]]
		tagMatch := tagValueRe.FindStringSubmatch(string(blockBytes))

		var tagValue string
		if len(tagMatch) >= 2 {
			rawTag := tagMatch[1]
			tagValue = regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(rawTag, "")
		}

		if len(tagValue) == 4 {
			removalStartIndices = append(removalStartIndices, m[0])
			fmt.Printf("Removing faction: Tag='%s' (length=%d, meaning it must be an NPC)\n", tagValue, len(tagValue))
		} else if len(tagValue) > 0 {
			fmt.Printf("Keeping faction: Tag='%s' (length is not 4, meaning it must be a player faction)\n", tagValue)
		} else {
			fmt.Println("Keeping faction: no Tag field. This should never happen.")
		}
	}

	result := factionBlockRe.ReplaceAllFunc(data, func(match []byte) []byte {
		tagMatch := tagValueRe.FindStringSubmatch(string(match))
		var keep bool
		if len(tagMatch) >= 2 {
			rawTag := tagMatch[1]
			tagValue := regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(rawTag, "")
			keep = len(tagValue) != 4
		} else {
			keep = true
		}

		if keep {
			return match                  
		}
		return nil                       
	})

	// Change GenerateFactionsOnStart to true in output (for debugging purposes - temporarily set to false to test faction removal)
	result = regexp.MustCompile(`(?i)<GenerateFactionsOnStart>([^<]*)</GenerateFactionsOnStart>`).ReplaceAllFunc(result, func(match []byte) []byte {
		return []byte(`<GenerateFactionsOnStart>true</GenerateFactionsOnStart>`)
	})

	err = os.WriteFile(outputFile, result, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Printf("Processing complete. Output written to %s (%d bytes)\n", outputFile, len(result))
}

package ourcode

import (
	"fmt"
	"os"
	"strings"
)

func GenerateASCIIArt(text, banner string) (string, error) {
	


	// Read banner file
	bannerFile := fmt.Sprintf("banners/%s.txt", banner)
	// newslice:=[]string{}
	// newslice = append(newslice, banner)
	// fmt.Print(newslice)
	content, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("We can read the banner file")
		return "", fmt.Errorf("could not read banner file: %v", err)
	}
	// lines := strings.Split(string(content), "\n")
	
	// if bannerFile=="shadow"||bannerFile=="thinkertoy"{
	// 	lines = strings.Split(string(content), "\r\n")
	// }


		lines := strings.Split(string(content), "\n")
	if banner=="thinkertoy"||banner=="shadow"{
		lines = strings.Split(string(content), "\r\n")
	}
	
	// Split text by newlines
	textLines := strings.Split(text, "\n")
	

	result := ""
	for _, line := range textLines {
		if line == "" {
			// result.WriteString("\n")
			continue
		}

		// Generate 8 lines for each text line
		for i := 0; i < 9; i++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					// result="You  don't put a caracter from ascii 😆"
					// fmt.Println("You  don't put a caracter from ascii 😆")!
					continue
				}

				startLine := int(char-32)*9 + i
				if startLine < len(lines)&&startLine>=0{
					result += lines[startLine]
				}
			}
			result += "\n"
		}
	}
// fmt.Println("ASCII Result:\n", result,banner)
	return result, nil
}

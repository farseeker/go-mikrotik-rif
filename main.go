package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		printHelp()
	}
	filename := os.Args[1]
	if filename == "" {
		printHelp()
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening RIF file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var section strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if line == "--BEGIN ROUTEROS SUPOUT SECTION" {
			section.Reset()
			continue
		}
		if line == "--END ROUTEROS SUPOUT SECTION" {
			err := mikrotikDecode(section.String())
			if err != nil {
				fmt.Println("Error decoding RIF section:", err)
			}
		}
		section.WriteString(line)
	}

	return
}

func printHelp() {
	fullPath := strings.Replace(os.Args[0], "\\", "/", -1) // Windows uses \ for dir separator, which doesn't work on path.Split
	_, exe := path.Split(fullPath)
	fmt.Printf("Usage: %s /path/to/supout.rif\n", exe)
	os.Exit(1)
}

func mikrotikDecode(section string) error {
	if len(section) == 0 {
		return fmt.Errorf("Empty section data")
	}

	const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=" //Terminating "=" is so that index% 64 == 0 for pad char
	var out []byte
	for i := 0; i < len(section); i += 4 {
		packet := section[i : i+4]
		o := strings.Index(b64, string(packet[3]))%64<<18 |
			strings.Index(b64, string(packet[2]))%64<<12 |
			strings.Index(b64, string(packet[1]))%64<<6 |
			strings.Index(b64, string(packet[0]))%64

		out = append(out, byte(o%256), byte((o>>8)%256), byte((o>>16)%256))
	}

	sectionSplit := bytes.Index(out, []byte{0x0})
	sectionName := string(out[0:sectionSplit])
	sectionDataZ := out[sectionSplit+1:]

	zR, err := zlib.NewReader(bytes.NewReader(sectionDataZ))

	if err != nil {
		return err
	}
	defer zR.Close()

	fmt.Println("== SECTION", sectionName)
	sectionData, err := ioutil.ReadAll(zR)
	if err != nil {
		return err
	}
	fmt.Println(string(sectionData))
	fmt.Println()

	return nil
}

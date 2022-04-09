package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var spfRecord, dmarcRecord string
	var hasMX, hasSPF, hasDMARC bool

	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Err: %v", err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)

	if err != nil {
		hasSPF = false
		log.Printf("Err: %v", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		hasDMARC = false
		log.Printf("Err: %v", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%t\n%t\n%s\n%t\n%s\n", hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

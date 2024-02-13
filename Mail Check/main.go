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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain,hasMX,hasSPF,SPFrecord,hasDMARC,dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read from input:%v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasDMARC, hasSPF bool
	var SPFrecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFrecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v", domain, hasMX, hasSPF, SPFrecord, hasDMARC, dmarcRecord)
}

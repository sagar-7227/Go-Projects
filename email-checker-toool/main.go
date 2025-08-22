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
	fmt.Printf("domain,hasMX,hasSPF,sprfRecord,hasDMARC,dmarcRecord\n")

	for scanner.Scan() {
		
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkDomain(domain string){

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	
	fmt.Printf("%s,%t,%t,\"%s\",%t,\"%s\"\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error looking up MX records for domain %s: %v", domain, err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error looking up TXT records for domain %s: %v", domain, err)
	}

	for _, record := range txtRecords {

		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error looking up DMARC records for domain %s: %v", domain, err)
	}

	for _, record := range dmarcRecords {

		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%s,%v,%v,%v,%v,%v",domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	fmt.Println()

}
package main
import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"strings"
)

// experiment with certificate management related code libs in Go:
func main() {

	certPool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	rawSubjects := certPool.Subjects()


	fmt.Println(len(rawSubjects), " certs found.")

	// 171 certs with cgo enabled
	// 197 certs with cgo disabled!



	// return list of byte slices containing DER encoding raw subjects

	for _, rawSubject := range rawSubjects {

		var rdnSubject pkix.RDNSequence

		_, err := asn1.Unmarshal(rawSubject, &rdnSubject)
		if err != nil {
			panic(err)
		}
		name := pkix.Name{}
		name.FillFromRDNSequence(&rdnSubject)

		//fmt.Println(name.SerialNumber)
		//fmt.Println(name.CommonName)

		commonName := strings.ToLower(name.CommonName)
		if strings.Contains(commonName, "kube") {
			fmt.Println("woot")
			fmt.Println(name)
		}
	}
	
	/*
	for _, rawSubject:= range rawSubjects {

		var rdnSubject pkix.RDNSequence
		_, err := asn1.Unmarshal(rawSubject, &rdnSubject)
		if err != nil {
			panic(err)
		}

		for _, foo := range rdnSubject {
			for _, bar := range foo {
				lbar := strings.ToLower(bar.Value.(string))
				if strings.Contains(lbar, "kube") {
					fmt.Println("woot")
					fmt.Println(lbar)
				}
			}
		}


	}*/
}


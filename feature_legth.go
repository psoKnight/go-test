package main

import (
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(len("DhYKB/Xw5wEEGuYWLdDY9Nge4/f4HNwBsxoh+APUHioP3csf2e/4FhAR9NfrJiQI9uoJBjUYHgQM4hfmDSALHvkKEaoBzt3P/RX/Eh/zDTs05A3m3hEJ5wgVARL5vPzjFwopGQUfH+sU/Pnm7xb7C8L04v883Czk2CAKCf3X+OMOI93cEQzxJA3r/u5FGTLk/Cbz19jlKRoMywcC8hjm7/XR+wsFCe0d8hQJMQQPG/np4/7+vb8AIBv08+rtBQwLHzkzAQYJxvwwCezk7SAZK/oEADP35/LX/hPm1/cFB0/V/Ake1fL//Mf8784cAxbt/N7n2hLkPv3sINPPDfi/Iw=="))

	decodeString, err := base64.StdEncoding.DecodeString("DhYKB/Xw5wEEGuYWLdDY9Nge4/f4HNwBsxoh+APUHioP3csf2e/4FhAR9NfrJiQI9uoJBjUYHgQM4hfmDSALHvkKEaoBzt3P/RX/Eh/zDTs05A3m3hEJ5wgVARL5vPzjFwopGQUfH+sU/Pnm7xb7C8L04v883Czk2CAKCf3X+OMOI93cEQzxJA3r/u5FGTLk/Cbz19jlKRoMywcC8hjm7/XR+wsFCe0d8hQJMQQPG/np4/7+vb8AIBv08+rtBQwLHzkzAQYJxvwwCezk7SAZK/oEADP35/LX/hPm1/cFB0/V/Ake1fL//Mf8784cAxbt/N7n2hLkPv3sINPPDfi/Iw==")
	if err != nil {
		logrus.Errorf("decode str err: %v", err)
		return
	}
	fmt.Println(len(decodeString))
	fmt.Println(decodeString)
}

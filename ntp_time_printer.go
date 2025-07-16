package ntp_time_printer

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func PrintNtpTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(time)
}

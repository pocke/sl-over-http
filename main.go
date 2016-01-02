package main

import (
	"bufio"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	f := func(w http.ResponseWriter, r *http.Request) {
		c := exec.Command("sl")
		out, err := c.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		sc := bufio.NewScanner(out)
		sc.Split(bufio.ScanBytes)

		err = c.Start()
		if err != nil {
			log.Fatal(err)
		}

		for sc.Scan() {
			w.Write(sc.Bytes())
			w.(http.Flusher).Flush()
		}
	}
	http.HandleFunc("/", f)
	http.ListenAndServe(":9999", nil)
}

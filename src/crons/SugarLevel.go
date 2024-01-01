// crons/SugarLevel.go
package crons

import (
    "time"
	"log"
	"io/ioutil"
    "net/http"
	"strconv"
	"os"
)

func StartCron() {
    ticker := time.NewTicker(5 * time.Second)

    func() {
        for {
            select {
            case <-ticker.C:
                client := &http.Client{}
				location, err := time.LoadLocation("Asia/Calcutta")
				if err != nil {
					log.Fatal(err)
				}

				now := time.Now().In(location)

				midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

				// Convert to Unix timestamp
				timestamp := midnight.Unix()

				req, err := http.NewRequest("GET", os.Getenv("API_URL") + "?timestamp="+strconv.FormatInt(timestamp, 10) + "&uh_user=false", nil)
				if err != nil {
					log.Printf("Error creating request:", err)
					return
				}

				req.Header.Add("authorization", os.Getenv("AUTHORIZATION"))
				req.Header.Add("timezone", os.Getenv("TIMEZONE"))
				req.Header.Add("apiversion", os.Getenv("API_VERSION"))


                resp, err := client.Do(req)
                if err != nil {
                    log.Printf("Error making request:", err)
                    return
                }
                defer resp.Body.Close()

                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    log.Printf("Error reading response:", err)
                    return
                }

                log.Printf("Response:", string(body))
            }
        }
    }()
}
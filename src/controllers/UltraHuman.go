// controllers/root.go
package controllers

import (
    "net/http"
    "os"
    "log"
    "io/ioutil"
    "encoding/json"
    "strings"
)

type BIOSResponse struct {
    Data struct {
        MetabaseURL string `json:"metabase_url"`
    } `json:"data"`
}

type MetabaseURLResponse struct {
    Data struct {
        Rows [][]interface{} `json:"rows"`
    } `json:"data"`
}

type StandardResponse struct {
    Payload [][]interface{} `json:"payload"`
    Error   string      `json:"error"`
    Status  int         `json:"status"`
}

func GetBiosData(response http.ResponseWriter, request *http.Request) {
    switch request.Method {
    case http.MethodGet:
        log.Printf("Received request from client IP: %s, Method: %s, URL: %s\n", request.RemoteAddr, request.Method, request.URL)

        passcode := request.URL.Query().Get("passcode")
        if passcode == "" {
            resp := StandardResponse{
                Payload: nil,
                Error:   "Passcode is required",
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        log.Printf("Passcode: ", passcode)

        req, err := http.NewRequest("GET", os.Getenv("BIOS_API_URL") + "?passcode="+ passcode, nil)
        if err != nil {
            resp := StandardResponse{
                Payload: nil,
                Error:   err.Error(),
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            resp := StandardResponse{
                Payload: nil,
                Error:   err.Error(),
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            resp := StandardResponse{
                Payload: nil,
                Error:   err.Error(),
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        // Convert body to string and print it
        bodyString := string(body)
        log.Printf("Response Body: %s", bodyString)

        var biosResponse BIOSResponse
        err = json.Unmarshal(body, &biosResponse)
        if err != nil {
            resp := StandardResponse{
                Payload: nil,
                Error:   err.Error(),
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        // Now you can access the Metabase URL
        metabaseURL := biosResponse.Data.MetabaseURL
        log.Printf("Metabase URL: %s", metabaseURL)

        // Extract JWT token
        parts := strings.Split(metabaseURL, "dashboard/")
        if len(parts) < 2 {
            resp := StandardResponse{
                Payload: nil,
                Error:   "Error: JWT token not found in URL",
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        jwtParts := strings.Split(parts[1], "#theme")
        if len(jwtParts) < 1 {
            resp := StandardResponse{
                Payload: nil,
                Error:   "Error: JWT token not found in URL",
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        jwtToken := jwtParts[0]
        log.Printf("JWT Token: %s", jwtToken)


        req, err = http.NewRequest("GET", os.Getenv("BIOS_METABASE_API_URL")+"/"+jwtToken+"/"+os.Getenv("BIOS_AVERAGE_SUGAR_CARD"), nil)
        if err != nil {
            resp := StandardResponse{
                Payload: nil,
                Error:   err.Error(),
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        client = &http.Client{}
        resp, err = client.Do(req)
        if err != nil {
            resp := StandardResponse{
                Payload: nil,
                Error:   err.Error(),
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }
        defer resp.Body.Close()

        body, err = ioutil.ReadAll(resp.Body)
        if err != nil {
            resp := StandardResponse{
                Payload: nil,
                Error:   err.Error(),
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        // Convert body to string and print it
        bodyString = string(body)
        log.Printf("Response Body: %s", bodyString)

        var metabaseResponse MetabaseURLResponse
        err = json.Unmarshal(body, &metabaseResponse)
        if err != nil {
            resp := StandardResponse{
                Payload: nil,
                Error:   err.Error(),
                Status:  http.StatusBadRequest,
            }
            json.NewEncoder(response).Encode(resp)
            return
        }

        // Now you can access the Metabase URL
        averageGlucose := metabaseResponse.Data.Rows
        log.Printf("Average Glucose url response: %s", averageGlucose)

        responseAverageGlucose := make([][]interface{}, len(averageGlucose))

        for i, v := range averageGlucose {
            responseAverageGlucose[i] = []interface{}{v[0], v[2]}
        }
        log.Printf("Average Glucose response: %s", responseAverageGlucose)

        // Create a StandardResponse
        returnResp := StandardResponse{
            Payload: responseAverageGlucose,
            Error:   "",
            Status:  http.StatusOK,
        }

        // Set the content type to application/json
        response.Header().Set("Content-Type", "application/json")

        // Write the StandardResponse to the response
        json.NewEncoder(response).Encode(returnResp)

    default:
        http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
// controllers/root.go
package controllers

import (
    "net/http"
    "html/template"
)

func RenderPrivacyPage(response http.ResponseWriter, request *http.Request) {
    switch request.Method {
    case http.MethodGet:
        data := map[string]string{
            "EffectiveDate": "January 6, 2024",
            "Passcode":      "We collect the passcode you provide as part of the API request for authentication purposes.",
            // Add more data as needed
        }

        // Define an HTML template
        tmplStr := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Privacy Policy for CGM Analyser API</title>
        </head>
        <body>
            <h1>Privacy Policy for CGM Analyser API</h1>
            <p>Effective Date: {{.EffectiveDate}}</p>

            <h2>Introduction</h2>
            <p>Welcome to the CGM Analyser API! This Privacy Policy is designed to help you understand how we collect, use, and safeguard your data when you access and use our API.</p>

            <!-- Add more sections and data using {{.Passcode}} -->

            <h2>Data Sharing and Disclosure</h2>
            <p>We do not share or disclose your data with third parties, except in the following circumstances:</p>
            <ul>
                <li>When required by law: We may share your data with law enforcement authorities or other governmental agencies when required by applicable laws and regulations.</li>
            </ul>

            <!-- Add more sections and data using {{.Passcode}} -->

            <h2>Contact Us</h2>
            <p>If you have any questions, concerns, or requests related to this Privacy Policy or the CGM Analyser API, please contact us.</p>

            <p>Thank you for using the CGM Analyser API. Your privacy is important to us, and we are committed to protecting your data.</p>
        </body>
        </html>
        `

        // Create a template and parse it
        tmpl, err := template.New("privacy-policy").Parse(tmplStr)
        if err != nil {
            http.Error(response, err.Error(), http.StatusInternalServerError)
            return
        }

        // Execute the template with the data and write it to the response
        err = tmpl.Execute(response, data)
        if err != nil {
            http.Error(response, err.Error(), http.StatusInternalServerError)
            return
        }
    default:
        http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

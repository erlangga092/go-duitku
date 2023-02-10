package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type CreateInvoiceResponse struct {
	MerchantCode  string `json:"merchantCode"`
	Reference     string `json:"reference"`
	PaymentURL    string `json:"paymentUrl"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

type CreateInvoiceRequest struct {
	PaymentAmount   int         `json:"paymentAmount"`
	MerchantOrderID string      `json:"merchantOrderId"`
	ProductDetails  string      `json:"productDetails"`
	Email           string      `json:"email"`
	ItemDetails     interface{} `json:"itemDetails"`
	CallbackURL     string      `json:"callbackUrl"`
	ReturnURL       string      `json:"returnUrl"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error when load .env file!")
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	merchantCode := os.Getenv("DUITKU_MERCHANT_CODE")
	apiKey := os.Getenv("DUITKU_API_KEY")
	URL := os.Getenv("DUITKU_BASE_URL")

	r.Get("/test-payment-duitku", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		jsonDataReq := CreateInvoiceRequest{
			PaymentAmount:   50000,
			MerchantOrderID: "1648542419",
			ProductDetails:  "Test Pay with duitku",
			Email:           "test@test.com",
			CallbackURL:     "https://example.com/api-pop/backend/callback.php",
			ReturnURL:       "https://example.com/api-pop/backend/redirect.php",
			ItemDetails: []map[string]interface{}{
				{"name": "Apel", "quantity": 2, "price": 50000},
			},
		}

		jsonDataReqBytes, err := json.Marshal(jsonDataReq)
		if err != nil {
			fmt.Println(err.Error())
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/createInvoice", URL), bytes.NewBuffer(jsonDataReqBytes))
		if err != nil {
			panic(err.Error())
		}

		currTimestamp := time.Now().UnixNano() / int64(time.Millisecond)
		currTimestampForHeader := strconv.Itoa(int(currTimestamp))

		sigString := fmt.Sprintf("%s%v%s", merchantCode, currTimestampForHeader, apiKey)
		hash := sha256.Sum256([]byte(sigString))
		sigHexHash := hex.EncodeToString(hash[:])

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-duitku-signature", sigHexHash)
		req.Header.Set("x-duitku-timestamp", currTimestampForHeader)
		req.Header.Set("x-duitku-merchantcode", merchantCode)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
		}

		defer resp.Body.Close()

		fmt.Println("response status:", resp.Status)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	fmt.Println("[ Server ] running on port 9000")

	http.ListenAndServe(":9000", r)
}

package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"simpleBankingApplication/models"

	"github.com/gin-gonic/gin"
)

func MakePayment(context *gin.Context) {
	var transaction models.Transaction
	err := context.BindJSON(&transaction)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid request",
		})
		return
	}

	// debit the account sending the money
	err = models.DebitAccount(transaction.UserID,transaction.Amount)

	

	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"error":"An error occurred while debiting the account"})
		return
	}

	// transaction.Transaction_reference = models.GenerateRandomNumbers(30);

	// Call the third party service
    //    request:= models.Request{
	// 	Amount: transaction.Amount,
	// 	Account_id: transaction.AccountNumber,
	// 	Reference: transaction.Transaction_reference,
	//    }
	// _,err = CallPaymentService(request)

	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError,gin.H{"error":"An error occurred while making the payment"})
	// 	return 
	// }

	// credit the account receiving the money
	err = models.CreditAccount(transaction.AccountNumber,transaction.Amount)

	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"error":"An error occurred while crediting the account"})
		return
	}

	err = transaction.SaveTransaction()

	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"error":"An error occurred while saving the transaction"})
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message":"Transaction saved successfully"})
	
}

func GetPaymentByReference(context *gin.Context) {
	reference := context.Param("reference")
	payment,err := models.GetPaymentByReference(reference)

	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"error":"An error occurred while fetching the payment"})
		return
	}

	context.JSON(http.StatusOK,payment)
}

func CallPaymentService(request models.Request) (models.Response,error) {
	url := "http://thirdparty.com/make-payment"
	data,err := json.Marshal(request)

	if err != nil {
		return models.Response{},err
	}
	requestBody := bytes.NewBuffer(data)
	respone, err := http.Post(url,"application/json",requestBody)

	if err != nil {
		return models.Response{},err
	}
	defer respone.Body.Close()

	if respone.StatusCode != http.StatusOK {
		return models.Response{},fmt.Errorf("azn error occurred while making the payment")
	}

	var resp models.Response
	err = json.NewDecoder(respone.Body).Decode(&resp)
	if err != nil {
		return models.Response{},err
	}

	return resp,nil
}

func GetPaymentsByReference(reference string) (models.Response,error) {
	url := fmt.Sprintf("http://thirdparty.com/get-payment/%s",reference)
	response,err := http.Get(url)

	if err != nil {
		return models.Response{},err
	}
	defer response.Body.Close()

	var rep models.Response
	err = json.NewDecoder(response.Body).Decode(&rep)

	if err != nil {
		return models.Response{},err
	}

	return rep,nil
}
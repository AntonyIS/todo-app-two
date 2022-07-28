package app

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserModel struct {
	PublicID    string `json:"publicID" dynamodbav:"publicID"`
	FirstName   string `json:"firstname" dynamodbav:"firstname"`
	LastName    string `json:"lastname" dynamodbav:"lastname"`
	Email       string `json:"email" dynamodbav:"email"`
	Phone       string `json:"phone" dynamodbav:"phone"`
	CompanyName string `json:"companyName" dynamodbav:"companyName"`
	Password    string `json:"Password" dynamodbav:"password"`
	Role        string `json:"role" dynamodbav:"role"`
	IsClient    bool   `json:"isClient" dynamodbav:"isClient"`
	IsAdmin     bool   `json:"isAdmin" dynamodbav:"isAdmin"`
	IsManager   bool   `json:"isManager" dynamodbav:"isManager"`
	IsAPIUser   bool   `json:"isAPIUser" dynamodbav:"isAPIUser"`
	CreatedAt   int64  `json:"createdAt" dynamodbav:"createdAt"`
}

func (user UserModel) GetKey() map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(user.PublicID)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": id}
}

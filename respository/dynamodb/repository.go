package dynamodb

import (
	"context"

	"bitbucket.com/AntonyIS/vision/app"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

type dynamodbRepostory struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewDynamoDBRepository(TableName string) (app.UserRepositoryInterface, error) {
	repo := dynamodbRepostory{}
	repo.TableName = TableName
	return &repo, nil
}

func (d *dynamodbRepostory) CreateUser(u *app.UserModel) error {
	user, err := attributevalue.MarshalMap(u)
	if err != nil {
		return errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.CreateUser")
	}
	_, err = d.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(d.TableName),
		Item:      user,
	})

	if err != nil {
		return errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.CreateUser")
	}

	return nil
}

func (d *dynamodbRepostory) User(id string) (*app.UserModel, error) {
	user := app.UserModel{PublicID: id}
	response, err := d.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:       user.GetKey(),
		TableName: aws.String(d.TableName),
	})

	if err != nil {
		return nil, errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.User")
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &user)
		if err != nil {
			return nil, errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.User")
		}
	}
	return &user, nil
}

func (d *dynamodbRepostory) Users() (*[]app.UserModel, error) {
	var users []app.UserModel
	var response *dynamodb.QueryOutput
	keyEx := expression.Key("").Equal(expression.Value(""))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.Users")
	} else {
		response, err = d.DynamoDbClient.Query(context.TODO(), &dynamodb.QueryInput{
			TableName:                 aws.String(d.TableName),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		})
		if err != nil {
			return nil, errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.User")
		} else {
			err = attributevalue.UnmarshalListOfMaps(response.Items, &users)
			if err != nil {
				return nil, errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.User")
			}
		}
	}
	return &users, nil

}

func (d *dynamodbRepostory) UpdateUser(u *app.UserModel) error {
	var response *dynamodb.UpdateItemOutput
	var attributeMap map[string]map[string]interface{}
	update := expression.Set(expression.Name("firstname"), expression.Value(u.FirstName))
	update.Set(expression.Name("lastname"), expression.Value(u.LastName))
	update.Set(expression.Name("email"), expression.Value(u.Email))
	update.Set(expression.Name("phone"), expression.Value(u.Phone))
	update.Set(expression.Name("companyName"), expression.Value(u.CompanyName))
	update.Set(expression.Name("password"), expression.Value(u.Password))
	update.Set(expression.Name("role"), expression.Value(u.Role))
	update.Set(expression.Name("isClient"), expression.Value(u.IsClient))
	update.Set(expression.Name("isAdmin"), expression.Value(u.IsAdmin))
	update.Set(expression.Name("isManager"), expression.Value(u.IsManager))
	update.Set(expression.Name("isAPIUser"), expression.Value(u.IsAPIUser))
	update.Set(expression.Name("createdAt"), expression.Value(u.CreatedAt))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.UpdateUser")
	} else {
		response, err = d.DynamoDbClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName:                 aws.String(d.TableName),
			Key:                       u.GetKey(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			return errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.UpdateUser")
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &attributeMap)
			if err != nil {
				return errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.UpdateUser")
			}
		}
	}
	return nil
}

func (d *dynamodbRepostory) DeleteUser(id string) error {
	user := app.UserModel{PublicID: id}

	_, err := d.DynamoDbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(d.TableName), Key: user.GetKey(),
	})
	if err != nil {
		return errors.Wrap(app.ErrorInternalServerError, "dynamodb.repositry.User")
	}
	return err
}

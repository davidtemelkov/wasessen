package data

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type RecipeQueueItem struct {
	ID         string `json:"id"`
	RecipeName string `json:"recipe_name"`
	Cook       string `json:"cook"`
	Position   int    `json:"position"`
}

func GetRecipeQueueItems(ctx context.Context) ([]RecipeQueueItem, error) {
	expressionAttributeNames := map[string]string{
		"#pk": PK,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{
			Value: RECIPE_QUEUE,
		},
	}

	scanInput := &dynamodb.ScanInput{
		TableName:                 aws.String(TABLE_NAME),
		FilterExpression:          aws.String("#pk = :pk"),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
	}

	result, err := Db.Scan(ctx, scanInput)
	if err != nil {
		return nil, err
	}

	recipeQueueItems := make([]RecipeQueueItem, 0)
	for _, item := range result.Items {
		recipeQueueItem := RecipeQueueItem{}
		if err := attributevalue.UnmarshalMap(item, &recipeQueueItem); err != nil {
			return nil, err
		}

		recipeQueueItems = append(recipeQueueItems, recipeQueueItem)
	}

	return recipeQueueItems, nil
}

func InsertRecipeQueueItem(ctx context.Context, recipeName string, cook string) error {
	id := uuid.New().String()

	recipeQueueCount, err := getRecipeQueueCount(ctx)
	if err != nil {
		return err
	}

	item := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{
			Value: RECIPE_QUEUE,
		},
		SK: &types.AttributeValueMemberS{
			Value: id,
		},
		ID: &types.AttributeValueMemberS{
			Value: id,
		},
		RECIPE_NAME: &types.AttributeValueMemberS{
			Value: recipeName,
		},
		COOK: &types.AttributeValueMemberS{
			Value: cook,
		},
		POSITION: &types.AttributeValueMemberN{
			Value: strconv.Itoa(recipeQueueCount + 1),
		},
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      item,
	}

	_, err = Db.PutItem(ctx, putItemInput)
	if err != nil {
		return err
	}

	err = updateRecipeQueueCount(ctx, true)
	if err != nil {
		return err
	}

	return nil
}

func RemoveRecipeQueueItem(ctx context.Context, id string) error {
	key := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{Value: RECIPE_QUEUE},
		SK: &types.AttributeValueMemberS{Value: id},
	}

	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(TABLE_NAME),
		Key:       key,
	}

	_, err := Db.DeleteItem(ctx, deleteItemInput)
	if err != nil {
		return err
	}

	err = updateRecipeQueueCount(ctx, false)
	if err != nil {
		return err
	}

	return nil
}

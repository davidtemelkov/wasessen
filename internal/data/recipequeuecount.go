package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type RecipeQueueCount struct {
	Count int `json:"count"`
}

func getRecipeQueueCount(ctx context.Context) (int, error) {
	key := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{Value: RECIPE_QUEUE_COUNT},
		SK: &types.AttributeValueMemberS{Value: RECIPE_QUEUE_COUNT},
	}

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key:       key,
	}

	result, err := Db.GetItem(ctx, getItemInput)
	if err != nil {
		return 0, err
	}

	if result.Item == nil {
		return 0, nil
	}

	var recipeQueueCount RecipeQueueCount
	if err := attributevalue.UnmarshalMap(result.Item, &recipeQueueCount); err != nil {
		return 0, err
	}

	return recipeQueueCount.Count, nil
}

func updateRecipeQueueCount(ctx context.Context, increment bool) error {
	var updateExpression string
	if increment {
		updateExpression = "SET #count = #count + :inc"
	} else {
		updateExpression = "SET #count = #count - :inc"
	}

	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: RECIPE_QUEUE_COUNT},
			"SK": &types.AttributeValueMemberS{Value: RECIPE_QUEUE_COUNT},
		},
		UpdateExpression: aws.String(updateExpression),
		ExpressionAttributeNames: map[string]string{
			"#count": "Count",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":inc": &types.AttributeValueMemberN{Value: "1"},
		},
	}

	_, err := Db.UpdateItem(ctx, updateItemInput)
	if err != nil {
		return err
	}

	return nil
}

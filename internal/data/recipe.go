package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type Recipe struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Preparation string `json:"preparation"`
	Difficulty  string `json:"difficulty"`
	ImageURL    string `json:"image_url"`
}

func InsertRecipe(ctx context.Context, recipe Recipe) error {
	id := uuid.New().String()

	item := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{
			Value: RECIPE,
		},
		SK: &types.AttributeValueMemberS{
			Value: id,
		},
		ID: &types.AttributeValueMemberS{
			Value: id,
		},
		NAME: &types.AttributeValueMemberS{
			Value: recipe.Name,
		},
		INGREDIENTS: &types.AttributeValueMemberS{
			Value: recipe.Ingredients,
		},
		PREPARATION: &types.AttributeValueMemberS{
			Value: recipe.Preparation,
		},
		DIFFICULTY: &types.AttributeValueMemberS{
			Value: recipe.Difficulty,
		},
		IMAGE_URL: &types.AttributeValueMemberS{
			Value: recipe.ImageURL,
		},
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      item,
	}

	_, err := Db.PutItem(ctx, putInput)
	if err != nil {
		return err
	}

	return nil
}

func GetRecipes(ctx context.Context) ([]Recipe, error) {
	expressionAttributeNames := map[string]string{
		"#pk": PK,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{
			Value: RECIPE,
		},
	}

	scanInput := &dynamodb.ScanInput{
		TableName:                 aws.String(TABLE_NAME),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
	}

	result, err := Db.Scan(ctx, scanInput)
	if err != nil {
		return nil, err
	}

	recipes := make([]Recipe, 0)
	for _, item := range result.Items {
		recipe := Recipe{}
		if err := attributevalue.UnmarshalMap(item, &recipe); err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func RemoveRecipe(ctx context.Context, id string) error {
	key := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{Value: RECIPE},
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

	return nil
}

// TODO: Add update recipe

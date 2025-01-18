package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Recipe struct {
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Preparation string `json:"preparation"`
	Difficulty  string `json:"difficulty"`
	ImageURL    string `json:"image_url"`
}

func InsertRecipe(ctx context.Context, recipe Recipe) error {
	item := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{
			Value: RECIPE,
		},
		SK: &types.AttributeValueMemberS{
			Value: RECIPE_PREFIX + recipe.Name,
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
	keyConditionExpression := "#pk = :pk AND begins_with(#sk, :sk)"
	expressionAttributeNames := map[string]string{
		"#pk": PK,
		"#sk": SK,
	}
	expressionAttributeValues := map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{
			Value: RECIPE,
		},
		":sk": &types.AttributeValueMemberS{
			Value: RECIPE_PREFIX,
		},
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(TABLE_NAME),
		KeyConditionExpression:    aws.String(keyConditionExpression),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
	}

	result, err := Db.Query(ctx, queryInput)
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

// TODO: Add update and remove recipe

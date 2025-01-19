package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(TABLE_NAME),
		KeyConditionExpression:    aws.String("#pk = :pk"),
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

func GetRecipeByID(ctx context.Context, id string) (Recipe, error) {
	key := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{Value: RECIPE},
		SK: &types.AttributeValueMemberS{Value: id},
	}

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key:       key,
	}

	result, err := Db.GetItem(ctx, getItemInput)
	if err != nil {
		return Recipe{}, err
	}

	// TODO: Err record not found
	if result.Item == nil {
		return Recipe{}, nil
	}

	var recipe Recipe
	if err := attributevalue.UnmarshalMap(result.Item, &recipe); err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

func UpdateRecipe(ctx context.Context, oldRecipe Recipe, newRecipe Recipe) error {
	update := expression.UpdateBuilder{}
	hasChanges := false

	if oldRecipe.Name != newRecipe.Name {
		update = update.Set(expression.Name(NAME), expression.Value(newRecipe.Name))
		hasChanges = true
	}
	if oldRecipe.Ingredients != newRecipe.Ingredients {
		update = update.Set(expression.Name(INGREDIENTS), expression.Value(newRecipe.Ingredients))
		hasChanges = true
	}
	if oldRecipe.Preparation != newRecipe.Preparation {
		update = update.Set(expression.Name(PREPARATION), expression.Value(newRecipe.Preparation))
		hasChanges = true
	}
	if oldRecipe.Difficulty != newRecipe.Difficulty {
		update = update.Set(expression.Name(DIFFICULTY), expression.Value(newRecipe.Difficulty))
		hasChanges = true
	}
	if oldRecipe.ImageURL != newRecipe.ImageURL && newRecipe.ImageURL != "" {
		update = update.Set(expression.Name(IMAGE_URL), expression.Value(newRecipe.ImageURL))
		hasChanges = true
	}

	// TODO: Return error or nil?
	if !hasChanges {
		return nil
	}

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]types.AttributeValue{
			PK: &types.AttributeValueMemberS{Value: RECIPE},
			SK: &types.AttributeValueMemberS{Value: oldRecipe.ID},
		},
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	_, err = Db.UpdateItem(ctx, updateInput)
	if err != nil {
		return err
	}

	return nil
}

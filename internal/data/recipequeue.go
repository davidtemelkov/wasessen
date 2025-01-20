package data

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type RecipeQueue struct {
	Queue []RecipeQueueItem `json:"queue"`
	Count int               `json:"count"`
}

type RecipeQueueItem struct {
	ID         string `json:"id"`
	RecipeName string `json:"recipe_name"`
	Cook       string `json:"cook"`
	Position   int    `json:"position"`
}

func GetRecipeQueue(ctx context.Context) (RecipeQueue, error) {
	key := map[string]types.AttributeValue{
		PK: &types.AttributeValueMemberS{Value: RECIPE_QUEUE},
		SK: &types.AttributeValueMemberS{Value: RECIPE_QUEUE},
	}

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key:       key,
	}

	result, err := Db.GetItem(ctx, getItemInput)
	if err != nil {
		return RecipeQueue{}, err
	}

	if result.Item == nil {
		return RecipeQueue{}, errors.New("not found") // Return better error if needed
	}

	recipeQueue := RecipeQueue{}
	if err := attributevalue.UnmarshalMap(result.Item, &recipeQueue); err != nil {
		return RecipeQueue{}, err
	}

	return recipeQueue, nil
}

func AddRecipeToQueue(ctx context.Context, recipeName, cook string) (RecipeQueue, error) {
	recipeQueue, err := GetRecipeQueue(ctx)
	if err != nil {
		return RecipeQueue{}, err
	}

	id := uuid.New().String()
	position := recipeQueue.Count + 1

	newItem := RecipeQueueItem{
		ID:         id,
		RecipeName: recipeName,
		Cook:       cook,
		Position:   position,
	}

	recipeQueue.Queue = append(recipeQueue.Queue, newItem)
	recipeQueue.Count++

	item, err := attributevalue.MarshalMap(recipeQueue)
	if err != nil {
		return RecipeQueue{}, err
	}

	item[PK] = &types.AttributeValueMemberS{
		Value: RECIPE_QUEUE,
	}
	item[SK] = &types.AttributeValueMemberS{
		Value: RECIPE_QUEUE,
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      item,
	}

	_, err = Db.PutItem(ctx, putItemInput)
	if err != nil {
		return RecipeQueue{}, err
	}

	return recipeQueue, nil
}

func RemoveRecipeFromQueue(ctx context.Context, id string) (RecipeQueue, error) {
	recipeQueue, err := GetRecipeQueue(ctx)
	if err != nil {
		return RecipeQueue{}, err
	}

	var indexToRemove int
	var found bool
	for i, item := range recipeQueue.Queue {
		if item.ID == id {
			indexToRemove = i
			found = true
			break
		}
	}

	if !found {
		return RecipeQueue{}, errors.New("item not found") // TODO: better error
	}

	recipeQueue.Queue = append(recipeQueue.Queue[:indexToRemove], recipeQueue.Queue[indexToRemove+1:]...)

	for i := indexToRemove; i < len(recipeQueue.Queue); i++ {
		recipeQueue.Queue[i].Position -= 1
	}

	recipeQueue.Count--

	item, err := attributevalue.MarshalMap(recipeQueue)
	if err != nil {
		return RecipeQueue{}, err
	}

	item[PK] = &types.AttributeValueMemberS{
		Value: RECIPE_QUEUE,
	}
	item[SK] = &types.AttributeValueMemberS{
		Value: RECIPE_QUEUE,
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      item,
	}

	_, err = Db.PutItem(ctx, putItemInput)
	if err != nil {
		return RecipeQueue{}, err
	}

	return recipeQueue, nil
}

func MoveRecipeInQueue(ctx context.Context, id string, up bool) (RecipeQueue, error) {
	recipeQueue, err := GetRecipeQueue(ctx)
	if err != nil {
		return RecipeQueue{}, err
	}

	var indexToMove int
	var found bool
	for i, item := range recipeQueue.Queue {
		if item.ID == id {
			indexToMove = i
			found = true
			break
		}
	}

	// If the item is not found or is at the boundary
	if !found || (up && indexToMove == 0) || (!up && indexToMove == len(recipeQueue.Queue)-1) {
		return recipeQueue, nil
	}

	// Determine the swap index based on direction
	swapIndex := indexToMove - 1
	if !up {
		swapIndex = indexToMove + 1
	}

	// Swap positions
	recipeQueue.Queue[indexToMove].Position, recipeQueue.Queue[swapIndex].Position =
		recipeQueue.Queue[swapIndex].Position, recipeQueue.Queue[indexToMove].Position

	// Reorder the slice
	recipeQueue.Queue[indexToMove], recipeQueue.Queue[swapIndex] =
		recipeQueue.Queue[swapIndex], recipeQueue.Queue[indexToMove]

	item, err := attributevalue.MarshalMap(recipeQueue)
	if err != nil {
		return RecipeQueue{}, err
	}

	item[PK] = &types.AttributeValueMemberS{
		Value: RECIPE_QUEUE,
	}
	item[SK] = &types.AttributeValueMemberS{
		Value: RECIPE_QUEUE,
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      item,
	}

	_, err = Db.PutItem(ctx, putItemInput)
	if err != nil {
		return RecipeQueue{}, err
	}

	return recipeQueue, nil
}

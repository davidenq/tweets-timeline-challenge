package dynamo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/davidenq/tweets-timeline-challenge/app/domain/usecases"
	"github.com/rs/zerolog/log"
)

type DynamoService struct {
	client dynamodb.Client
}

func (d DynamoService) GetLastRecord(tableName string, out interface{}) error {
	var limit int32 = 1
	result, err := d.client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(string(tableName)),
		Limit:     aws.Int32(limit),
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	data := map[string]interface{}{}
	for _, item := range result.Items {
		for key, value := range item {
			var valMap interface{}
			attributevalue.Unmarshal(value, &valMap)
			data[key] = valMap
		}
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	json.Unmarshal(bytes, &out)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}

func (d DynamoService) GetEntityByID(tableName, id string, out interface{}) error {
	return d.queryByTag(tableName, "id", id, out)
}

func (d DynamoService) GetEntityByUsername(tableName string, username string, out interface{}) error {
	return d.queryByTag(tableName, "username", username, out)
}

func (d DynamoService) SaveEntity(tableName string, entity interface{}) error {
	attributes := mapFromEntityToMapAttributes(entity)
	_, err := d.client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: aws.String(string(tableName)),
			Item:      attributes,
		},
	)

	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	return nil
}

func (d DynamoService) EditEntity(tableName string, entity map[string]interface{}) error {
	panic("not implemented yet")
}

func (d DynamoService) DeleteEntity(tableName, id string) error {
	panic("not implemented yet")
}

func (d DynamoService) CreateTables(schema interface{}) error {
	schematt := schema.(usecases.SchemaToTable)
	var units int64 = 1
	attributes := make([]types.AttributeDefinition, 0)
	for _, el := range schematt.Attributes {
		attributeDefinition := types.AttributeDefinition{
			AttributeName: aws.String(el.Name),
			AttributeType: d.checkAttributeType(el.Type),
		}
		attributes = append(attributes, attributeDefinition)
	}

	params := &dynamodb.CreateTableInput{
		TableName: aws.String(schematt.TableName),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String(schematt.Key),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: attributes,
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  &units,
			WriteCapacityUnits: &units,
		},
	}
	_, err := d.client.CreateTable(context.Background(), params)
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	log.Info().Msg("table " + schematt.TableName + " has been created")
	time.Sleep(5 * time.Second)
	//if schematt.TableName == string(domain.OAuth) {
	//	ttlInput := &dynamodb.UpdateTimeToLiveInput{
	//		TableName: aws.String(schematt.TableName),
	//		TimeToLiveSpecification: &types.TimeToLiveSpecification{
	//			AttributeName: aws.String("expires_on"),
	//			Enabled:       aws.Bool(true),
	//		},
	//	}
	//	_, err = d.client.UpdateTimeToLive(context.Background(), ttlInput)
	//	if err != nil {
	//		log.Error().Msg(err.Error())
	//		panic(err)
	//	}
	//	log.Info().Msg("table " + schematt.TableName + " has been updated")
	//}

	return nil
}

func (d DynamoService) RemoveTables() error {
	resp, err := d.client.ListTables(context.Background(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(10),
	})
	if err != nil {
		log.Error().Msg(err.Error())
	}
	for _, tableName := range resp.TableNames {
		_, err = d.client.DeleteTable(context.Background(), &dynamodb.DeleteTableInput{
			TableName: aws.String(tableName),
		})
		if err != nil {
			log.Error().Msg(err.Error())
		}
		log.Info().Msg("table " + tableName + " has been removed")
	}
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	return nil
}

func (d DynamoService) checkAttributeType(value string) types.ScalarAttributeType {
	switch value {
	case "bool":
		return types.ScalarAttributeTypeB
	default:
		return types.ScalarAttributeTypeS
	}
}

func mapFromEntityToMapAttributes(entity interface{}) map[string]types.AttributeValue {
	item := map[string]types.AttributeValue{}
	switch mapValues := entity.(type) {
	case map[string]string:
		for index, value := range mapValues {
			item[index] = &types.AttributeValueMemberS{Value: value}
		}
	case map[string]interface{}:
		for index, value := range mapValues {
			switch val := value.(type) {
			case bool:
				item[index] = &types.AttributeValueMemberBOOL{Value: val}
			default:
				item[index] = &types.AttributeValueMemberS{Value: val.(string)}
			}
		}
	}

	return item
}

func (d DynamoService) queryByTag(tableName, key, value string, out interface{}) error {
	result, err := d.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			key: &types.AttributeValueMemberS{Value: value},
		},
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	data := map[string]interface{}{}
	for key, value := range result.Item {
		var valMap interface{}
		attributevalue.Unmarshal(value, &valMap)
		data[key] = valMap
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	json.Unmarshal(bytes, &out)
	return nil
}

func NewRepository(client dynamodb.Client) *DynamoService {
	return &DynamoService{
		client: client,
	}
}

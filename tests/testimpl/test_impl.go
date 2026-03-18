package testimpl

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	eventbridgetypes "github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("VerifyTerraformOutputs", func(t *testing.T) {
		eventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
		configuredEventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "event_bus_name")
		assert.Equal(t, configuredEventBusName, eventBusName, "id output should match configured event bus name")
	})

	t.Run("VerifyPolicyViaAWSAPI", func(t *testing.T) {
		eventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
		require.NotEmpty(t, eventBusName, "event bus name required for API verification")

		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(getAWSRegion(t, ctx)))
		require.NoError(t, err)

		client := eventbridge.NewFromConfig(cfg)
		resp, err := client.DescribeEventBus(context.Background(), &eventbridge.DescribeEventBusInput{
			Name: aws.String(eventBusName),
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Policy, "Policy must be present on event bus — policy may not be configured")
		require.NotEmpty(t, *resp.Policy, "Policy document should not be empty")

		var policyDoc map[string]interface{}
		err = json.Unmarshal([]byte(*resp.Policy), &policyDoc)
		require.NoError(t, err)
		statements, ok := policyDoc["Statement"].([]interface{})
		require.True(t, ok, "Policy should have Statement array")
		assert.GreaterOrEqual(t, len(statements), 1, "Policy should have at least one statement")
	})

	t.Run("VerifyPutEventsSucceeds", func(t *testing.T) {
		eventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
		require.NotEmpty(t, eventBusName, "event bus name required for PutEvents")

		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(getAWSRegion(t, ctx)))
		require.NoError(t, err)

		client := eventbridge.NewFromConfig(cfg)
		resp, err := client.PutEvents(context.Background(), &eventbridge.PutEventsInput{
			Entries: []eventbridgetypes.PutEventsRequestEntry{
				{
					Source:       aws.String("terratest.cloudwatch_event_bus_policy"),
					DetailType:   aws.String("FunctionalTest"),
					Detail:       aws.String(`{"test":"put_events"}`),
					EventBusName: aws.String(eventBusName),
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Len(t, resp.Entries, 1)
		assert.NotEmpty(t, resp.Entries[0].EventId, "PutEvents should return event ID on success")
		assert.Empty(t, resp.Entries[0].ErrorCode, "PutEvents should not have error code")
	})
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	t.Run("VerifyTerraformOutputs", func(t *testing.T) {
		eventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
		configuredEventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "event_bus_name")
		assert.Equal(t, configuredEventBusName, eventBusName, "id output should match configured event bus name")
	})

	t.Run("VerifyPolicyViaAWSAPI", func(t *testing.T) {
		eventBusName := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
		require.NotEmpty(t, eventBusName, "event bus name required for API verification")

		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(getAWSRegion(t, ctx)))
		require.NoError(t, err)

		client := eventbridge.NewFromConfig(cfg)
		resp, err := client.DescribeEventBus(context.Background(), &eventbridge.DescribeEventBusInput{
			Name: aws.String(eventBusName),
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Policy, "Policy must be present on event bus — policy may not be configured")
		require.NotEmpty(t, *resp.Policy, "Policy document should not be empty")

		var policyDoc map[string]interface{}
		err = json.Unmarshal([]byte(*resp.Policy), &policyDoc)
		require.NoError(t, err)
		statements, ok := policyDoc["Statement"].([]interface{})
		require.True(t, ok, "Policy should have Statement array")
		assert.GreaterOrEqual(t, len(statements), 1, "Policy should have at least one statement")

		firstStmt, ok := statements[0].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "AllowSameAccountPutEvents", firstStmt["Sid"], "Policy statement Sid should match")
		assert.Equal(t, "Allow", firstStmt["Effect"], "Policy statement Effect should be Allow")
	})
}

func getAWSRegion(t *testing.T, ctx types.TestContext) string {
	// Use region from Terraform output to match where resources were deployed
	region, err := terraform.OutputE(t, ctx.TerratestTerraformOptions(), "region")
	if err == nil && region != "" {
		return region
	}
	// Fallback to env var or default
	if r := ctx.TerratestTerraformOptions().EnvVars["AWS_DEFAULT_REGION"]; r != "" {
		return r
	}
	return "us-east-1"
}

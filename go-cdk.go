package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type GoCdkStackProps struct {
	awscdk.StackProps
}

func NewGoCdkStack(scope constructs.Construct, id string, props *GoCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	function := awslambda.NewFunction(stack, jsii.String("gofunction"), &awslambda.FunctionProps{
		FunctionName: jsii.String("hellogo"),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./handler"), nil),
		Handler:      jsii.String("handler"),
		Runtime:      awslambda.Runtime_GO_1_X(),
	})

	integration := awsapigatewayv2integrations.NewLambdaProxyIntegration(&awsapigatewayv2integrations.LambdaProxyIntegrationProps{
		Handler:              function,
		PayloadFormatVersion: awsapigatewayv2.PayloadFormatVersion_VERSION_1_0(),
	})

	api := awsapigatewayv2.NewHttpApi(stack, jsii.String("goapi"), &awsapigatewayv2.HttpApiProps{
		ApiName: jsii.String("hellogoapi"),
	})

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/"),
		Integration: integration,
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewGoCdkStack(app, "GoCdkStack", &GoCdkStackProps{
		awscdk.StackProps{
			StackName: jsii.String("GoCDKStack"),
			Env:       env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}

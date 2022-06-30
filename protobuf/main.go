package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "google.golang.org/protobuf/encoding/protojson"
	aws "protobuf/aws"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

// const aws_default_region = "eu-central-1"

func getMemberAccounts() []types.Account {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	org := organizations.NewFromConfig(cfg)

	// describeOrgOutput, _ := org.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{})
	// masterAccountId := describeOrgOutput.Organization.MasterAccountId

	listAccountsOutput, _ := org.ListAccounts(ctx, &organizations.ListAccountsInput{})

	return listAccountsOutput.Accounts
}

func main() {
	var accounts []aws.Account
	for _, a := range getMemberAccounts() {
		account := aws.Account{
			Id:    *a.Id,
			Name:  *a.Name,
			Arn:   *a.Arn,
			Email: *a.Email,
		}
		accounts = append(accounts, account)
	}

	env, err := cel.NewEnv(
		cel.Types(&aws.Account{}),
		cel.Declarations(
			decls.NewVar("account",
				decls.NewObjectType("aws.Account"),
			),
		),
	)
	if err != nil {
		os.Exit(1)
	}

	// ast, issues := env.Compile(`name.startsWith("/groups/" + group) && 'UK' in ['US', 'UK']`)
	// ast, issues := env.Compile(`account.Id == '705083396685'`)
	// ast, issues := env.Compile(`! (account.Id in ['142883165113'])`)
	// ast, issues := env.Compile(`account.Name.startsWith("security")`)
	// ast, issues := env.Compile(`account.Name.endsWith("hub")`)
	// ast, issues := env.Compile(`["a", "b", "345"].exists_one(r, r=="346")`)
	ast, issues := env.Compile(`["234284372160", "705083396685"].exists_one(a, a==account.Id) // member accounts`)
	if issues != nil && issues.Err() != nil {
		log.Fatalf("compile type-check error: %s", issues.Err())
	}

	program, _ := env.Program(ast)
	if err != nil {
		log.Fatalf("program construction error: %s", err)
	}

	for _, account := range accounts {

		message := &account
		out, _, err := program.Eval(map[string]interface{}{"account": message})
		if err != nil {
			log.Fatalf("eval error: %v\n", err)
		}

		fmt.Printf("Eval results for account=%s: %v\n", account.Name, out)
	}
}

// p := pb.Person{
// 	Id:   1234,
// 	Name: "Jane Doe",
// }

// jsonBytes, _ := protojson.Marshal(&p)
// fmt.Println(string(jsonBytes))

// jsonInput := []byte("{\"name\":\"Billy Doe\", \"id\":999}")
// p2 := &pb.Person{}
// protojson.Unmarshal(jsonInput, p2)
// fmt.Printf("Person2: %v\n", p2)

// fmt.Printf("Person: %v\n", p)
// msg := p.ProtoMessage()
// fmt.Println(protojson.Format(msg))
// m := p.ProtoReflectM()
// fmt.Println(protojson.Format(m))

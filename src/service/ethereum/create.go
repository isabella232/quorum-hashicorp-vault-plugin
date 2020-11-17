package ethereum

import (
	"context"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/service/formatters"
	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (c *controller) NewCreateOperation() *framework.PathOperation {
	exampleAccount := utils.ExampleETHAccount()
	successExample := utils.Example200Response()

	return &framework.PathOperation{
		Callback:    c.createHandler(),
		Summary:     "Creates a new Ethereum account",
		Description: "Creates a new Ethereum account by generating a private key, storing it in the Vault and computing its public key and address",
		Examples: []framework.RequestExample{
			{
				Description: "Creates a new account on the tenant0 namespace",
				Data: map[string]interface{}{
					namespaceLabel: exampleAccount.Namespace,
				},
				Response: successExample,
			},
		},
		Responses: map[int][]framework.Response{
			200: {*successExample},
			400: {utils.Example400Response()},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) createHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		namespace := data.Get("namespace").(string)

		ctx = utils.WithLogger(ctx, c.logger)
		account, err := c.useCases.CreateAccount().WithStorage(req.Storage).Execute(ctx, namespace, "")
		if err != nil {
			return nil, err
		}

		return formatters.FormatAccountResponse(account), nil
	}
}
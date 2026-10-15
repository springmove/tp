package irita

import (
	"fmt"
	"testing"
	"time"

	"github.com/bianjieai/avata-sdk-go/v3/models"
	"github.com/springmove/sptty"
	"github.com/springmove/tp/src/base"
)

func initService() error {
	if base.IIrita == nil {
		srv := &Service{
			cfg: Config{
				BaseConfig: sptty.BaseConfig{
					Enable: true,
				},
				Configs: []IritaConfig{
					{
						URL:    "https://stage.apis.avata.bianjie.ai",
						Key:    "V2R63140G1C0G0Z2e1e7r4s6S0A243or",
						Secret: "G2D6K1h0v1l0H0G2g1J7y4u6b06233kT",
					},
				},
			},
		}

		if err := srv.initClients(); err != nil {
			return err
		}

		base.IIrita = srv
	}
	return nil
}

func TestCreateAccount(t *testing.T) {
	if err := initService(); err != nil {
		fmt.Println(err.Error())
		return
	}

	params := &models.CreateAccountReq{
		Name:        sptty.GenerateUID(),
		OperationID: fmt.Sprintf("%s-%d", "account", time.Now().UnixMicro()),
	}

	rt, err := base.IIrita.Clients().Account.CreateAccount(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", rt)
}

func TestQueryAccount(t *testing.T) {
	if err := initService(); err != nil {
		fmt.Println(err.Error())
		return
	}

	params := &models.QueryAccountsReq{
		PageKey: "",
	}

	rt, err := base.IIrita.Clients().Account.QueryAccounts(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", rt)
}

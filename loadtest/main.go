package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/codchen/new-chain-poc/app"
	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typestx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"google.golang.org/grpc"
)

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	// NOTE: this field will be renamed to Codec
	Marshaler codec.Codec
	TxConfig  client.TxConfig
	Amino     *codec.LegacyAmino
}

func MakeTestEncodingConfig() EncodingConfig {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	res := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             cdc,
	}
	std.RegisterLegacyAminoCodec(res.Amino)
	std.RegisterInterfaces(res.InterfaceRegistry)
	app.ModuleBasics.RegisterLegacyAminoCodec(res.Amino)
	app.ModuleBasics.RegisterInterfaces(res.InterfaceRegistry)
	return res
}

func main() {
	userHomeDir, _ := os.UserHomeDir()
	mnemonic := "guess kite refuse confirm salt cigar relax chase shiver august medal tiger dentist people armed quit hour flash sign column mountain gap tool accident"
	kr, _ := keyring.New(sdk.KeyringServiceName(), "os", filepath.Join(userHomeDir, ".new-chain-poc"), os.Stdin)
	keyringAlgos, _ := kr.SupportedAlgorithms()
	algoStr := string(hd.Secp256k1Type)
	algo, _ := keyring.NewSigningAlgoFromString(algoStr, keyringAlgos)
	hdpath := hd.CreateHDPath(sdk.GetConfig().GetCoinType(), 0, 0).String()
	derivedPriv, _ := algo.Derive()(mnemonic, "", hdpath)
	privKey := algo.Generate()(derivedPriv)
	// ko, _ := keyring.NewKeyOutput("tony", keyring.TypeLocal, sdk.AccAddress(privKey.PubKey().Address()), privKey.PubKey())
	// kos := []keyring.KeyOutput{ko}
	// out, err := yaml.Marshal(&kos)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintln(os.Stdout, string(out))

	encCfg := MakeTestEncodingConfig()
	contractAddr := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"

	txBytesArr := [][]byte{}
	for i := 50000; i < 55000; i++ {
		body := map[string]interface{}{
			"limit_buy": map[string]interface{}{
				"price":    100 + i,
				"quantity": 10,
			},
		}
		serialized_body, _ := json.Marshal(body)
		msg := wasmtypes.MsgExecuteContract{
			Sender:   sdk.AccAddress(privKey.PubKey().Address()).String(),
			Contract: contractAddr,
			Msg:      serialized_body,
		}
		txBuilder := encCfg.TxConfig.NewTxBuilder()
		_ = txBuilder.SetMsgs(&msg)
		txBuilder.SetGasLimit(1500000)
		txBuilder.SetFeeAmount([]sdk.Coin{
			sdk.NewCoin("stake", sdk.NewInt(1000)),
		})

		var sigsV2 []signing.SignatureV2
		sigV2 := signing.SignatureV2{
			PubKey: privKey.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: 0,
		}
		sigsV2 = append(sigsV2, sigV2)
		_ = txBuilder.SetSignatures(sigsV2...)
		sigsV2 = []signing.SignatureV2{}
		signerData := xauthsigning.SignerData{
			ChainID:       "new-chain-poc",
			AccountNumber: 0,
			Sequence:      0,
		}
		sigV2, _ = clienttx.SignWithPrivKey(
			encCfg.TxConfig.SignModeHandler().DefaultMode(),
			signerData,
			txBuilder,
			privKey,
			encCfg.TxConfig,
			0,
		)
		sigsV2 = append(sigsV2, sigV2)
		_ = txBuilder.SetSignatures(sigsV2...)

		txBytes, _ := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())

		txBytesArr = append(txBytesArr, txBytes)
	}

	grpcConn, _ := grpc.Dial(
		"127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	txClient := typestx.NewServiceClient(grpcConn)

	f, _ := os.OpenFile("/home/tony_chen/outputs/test_new_chain_5", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	var wg sync.WaitGroup
	var mu sync.Mutex

	fmt.Printf(time.Now().String())
	// overallStart := time.Now()
	for i, txBytes := range txBytesArr {
		wg.Add(1)

		i := i
		txBytes := txBytes

		go func() {
			defer wg.Done()

			grpcRes, err := txClient.BroadcastTx(
				context.Background(),
				&typestx.BroadcastTxRequest{
					Mode:    typestx.BroadcastMode_BROADCAST_MODE_SYNC,
					TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
				},
			)
			if err != nil {
				panic(err)
			}
			if grpcRes.TxResponse.Code != 0 {
				fmt.Printf("Error on %d: %d\n", i, grpcRes.TxResponse.Code)
			} else {
				mu.Lock()
				defer mu.Unlock()
				// duration := time.Since(start)
				f.WriteString(fmt.Sprintf("%s\n", grpcRes.TxResponse.TxHash))
			}

		}()
	}

	wg.Wait()
	// fmt.Printf("Overall: %s", time.Since(overallStart))
}

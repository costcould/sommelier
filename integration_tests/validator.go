package integration_tests

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"

	sdkcrypto "github.com/cosmos/cosmos-sdk/crypto"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	gravitytypes "github.com/peggyjv/gravity-bridge/module/x/gravity/types"
	"github.com/peggyjv/sommelier/v3/app"
	tmcfg "github.com/tendermint/tendermint/config"
	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
)

type validator struct {
	chain            *chain
	index            int
	moniker          string
	mnemonic         string
	keyInfo          keyring.Info
	privateKey       cryptotypes.PrivKey
	consensusKey     privval.FilePVKey
	consensusPrivKey cryptotypes.PrivKey
	nodeKey          p2p.NodeKey
	ethereumKey      ethereumKey
}

type ethereumKey struct {
	publicKey  string
	privateKey string
	address    string
}

func (v *validator) instanceName() string {
	return fmt.Sprintf("%s%d", v.moniker, v.index)
}

func (v *validator) configDir() string {
	return fmt.Sprintf("%s/%s", v.chain.configDir(), v.instanceName())
}

func (v *validator) createConfig() error {
	p := path.Join(v.configDir(), "config")
	return os.MkdirAll(p, 0755)
}

func (v *validator) init() error {
	if err := v.createConfig(); err != nil {
		return err
	}

	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	genDoc, err := getGenDoc(v.configDir())
	if err != nil {
		return err
	}

	appState, err := json.MarshalIndent(app.ModuleBasics.DefaultGenesis(cdc), "", " ")
	if err != nil {
		return fmt.Errorf("failed to JSON encode app genesis state: %w", err)
	}

	genDoc.ChainID = v.chain.id
	genDoc.Validators = nil
	genDoc.AppState = appState

	if err = genutil.ExportGenesisFile(genDoc, config.GenesisFile()); err != nil {
		return fmt.Errorf("failed to export app genesis state: %w", err)
	}

	tmcfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
	return nil
}

func (v *validator) createNodeKey() error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return err
	}

	v.nodeKey = *nodeKey
	return nil
}

func (v *validator) createConsensusKey() error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	pvKeyFile := config.PrivValidatorKeyFile()
	if err := tmos.EnsureDir(filepath.Dir(pvKeyFile), 0777); err != nil {
		return err
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err := tmos.EnsureDir(filepath.Dir(pvStateFile), 0777); err != nil {
		return err
	}

	filePV := privval.LoadOrGenFilePV(pvKeyFile, pvStateFile)
	v.consensusKey = filePV.Key

	return nil
}

func (v *validator) createKeyFromMnemonic(name, mnemonic string, passphrase string) error {
	kb, err := keyring.New(keyringAppName, keyring.BackendTest, v.configDir(), nil)
	if err != nil {
		return err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return err
	}

	info, err := kb.NewAccount(name, mnemonic, passphrase, sdk.FullFundraiserPath, algo)
	if err != nil {
		return err
	}

	privKeyArmor, err := kb.ExportPrivKeyArmor(name, keyringPassphrase)
	if err != nil {
		return err
	}

	privKey, _, err := sdkcrypto.UnarmorDecryptPrivKey(privKeyArmor, keyringPassphrase)
	if err != nil {
		return err
	}

	v.keyInfo = info
	v.mnemonic = mnemonic
	v.privateKey = privKey

	return nil
}

func (v *validator) createKey(name string) error {
	mnemonic, err := createMnemonic()
	if err != nil {
		return err
	}

	return v.createKeyFromMnemonic(name, mnemonic, "")
}

func (v *validator) generateEthereumKey() error {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("unexpected public key type; expected: %T, got: %T", &ecdsa.PublicKey{}, publicKey)
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	v.ethereumKey = ethereumKey{
		privateKey: hexutil.Encode(privateKeyBytes),
		publicKey:  hexutil.Encode(publicKeyBytes),
		address:    crypto.PubkeyToAddress(*publicKeyECDSA).Hex(),
	}

	return nil
}

func (v *validator) generateEthereumKeyFromMnemonic(mnemonic string) error {
	ethKey, err := ethereumKeyFromMnemonic(mnemonic)
	if err != nil {
		return err
	}
	v.ethereumKey = *ethKey
	return nil
}

func (v *validator) buildCreateValidatorMsg(amount sdk.Coin) (sdk.Msg, error) {
	description := stakingtypes.NewDescription(v.moniker, "", "", "", "")
	commissionRates := stakingtypes.CommissionRates{
		Rate:          sdk.MustNewDecFromStr("0.1"),
		MaxRate:       sdk.MustNewDecFromStr("0.2"),
		MaxChangeRate: sdk.MustNewDecFromStr("0.01"),
	}

	// get the initial validator min self delegation
	minSelfDelegation, _ := sdk.NewIntFromString("1")

	valPubKey, err := cryptocodec.FromTmPubKeyInterface(v.consensusKey.PubKey)
	if err != nil {
		return nil, err
	}

	return stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(v.keyInfo.GetAddress()),
		valPubKey,
		amount,
		description,
		commissionRates,
		minSelfDelegation,
	)
}

func (v *validator) buildDelegateKeysMsg() sdk.Msg {
	privKeyBz, err := hexutil.Decode(v.ethereumKey.privateKey)
	if err != nil {
		panic(fmt.Sprintf("failed to HEX decode private key: %s", err))
	}

	privKey, err := crypto.ToECDSA(privKeyBz)
	if err != nil {
		panic(fmt.Sprintf("failed to convert private key: %s", err))
	}

	signMsg := gravitytypes.DelegateKeysSignMsg{
		ValidatorAddress: sdk.ValAddress(v.keyInfo.GetAddress()).String(),
		Nonce:            0,
	}

	signMsgBz := cdc.MustMarshal(&signMsg)
	hash := crypto.Keccak256Hash(signMsgBz).Bytes()
	ethSig, err := gravitytypes.NewEthereumSignature(hash, privKey)
	if err != nil {
		panic(fmt.Sprintf("failed to create Ethereum signature: %s", err))
	}

	return gravitytypes.NewMsgDelegateKeys(
		sdk.ValAddress(v.keyInfo.GetAddress()),
		v.chain.orchestrators[v.index].keyInfo.GetAddress(),
		v.ethereumKey.address,
		ethSig,
	)
}

func (v *validator) signMsg(msgs ...sdk.Msg) (*sdktx.Tx, error) {
	txBuilder := encodingConfig.TxConfig.NewTxBuilder()

	if err := txBuilder.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	txBuilder.SetMemo(fmt.Sprintf("%s@%s:26656", v.nodeKey.ID(), v.instanceName()))
	txBuilder.SetFeeAmount(sdk.NewCoins())
	txBuilder.SetGasLimit(200000)

	signerData := authsigning.SignerData{
		ChainID:       v.chain.id,
		AccountNumber: 0,
		Sequence:      0,
	}

	// For SIGN_MODE_DIRECT, calling SetSignatures calls setSignerInfos on
	// TxBuilder under the hood, and SignerInfos is needed to generate the sign
	// bytes. This is the reason for setting SetSignatures here, with a nil
	// signature.
	//
	// Note: This line is not needed for SIGN_MODE_LEGACY_AMINO, but putting it
	// also doesn't affect its generated sign bytes, so for code's simplicity
	// sake, we put it here.
	sig := txsigning.SignatureV2{
		PubKey: v.keyInfo.GetPubKey(),
		Data: &txsigning.SingleSignatureData{
			SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
			Signature: nil,
		},
		Sequence: 0,
	}

	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	bytesToSign, err := encodingConfig.TxConfig.SignModeHandler().GetSignBytes(
		txsigning.SignMode_SIGN_MODE_DIRECT,
		signerData,
		txBuilder.GetTx(),
	)
	if err != nil {
		return nil, err
	}

	sigBytes, err := v.privateKey.Sign(bytesToSign)
	if err != nil {
		return nil, err
	}

	sig = txsigning.SignatureV2{
		PubKey: v.keyInfo.GetPubKey(),
		Data: &txsigning.SingleSignatureData{
			SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
			Signature: sigBytes,
		},
		Sequence: 0,
	}
	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	signedTx := txBuilder.GetTx()
	bz, err := encodingConfig.TxConfig.TxEncoder()(signedTx)
	if err != nil {
		return nil, err
	}

	return decodeTx(bz)
}

func (v *validator) keyring() (keyring.Keyring, error) {
	return keyring.New(keyringAppName, keyring.BackendTest, v.configDir(), nil)
}

func (v *validator) clientContext(nodeURI string) (*client.Context, error) {
	kb, err := v.keyring()
	if err != nil {
		return nil, err
	}
	return v.chain.clientContext(nodeURI, &kb, "val", v.keyInfo.GetAddress())
}

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	wutils "github.com/hyperledger-labs/weaver-dlt-interoperability/core/network/fabric-interop-cc/libs/utils"
)

type TokenAssetType struct {
	Issuer          string              `json:"issuer"`
	Value           int                 `json:"value"`
}
type TokenWallet struct {
	WalletMap       map[string]uint64   `json:"walletlist"`
}
type TokenAsset struct {
	Type          string      `json:"type"`
	NumUnits      uint64      `json:"numunits"`
	Owner         string      `json:"owner"`
}

func matchClaimWithTokenAssetPledge(pledgeAssetDetails, claimAssetDetails []byte) bool {
	var pledgeAsset, claimAsset TokenAsset
	err := json.Unmarshal(pledgeAssetDetails, &pledgeAsset)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = json.Unmarshal(claimAssetDetails, &claimAsset)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return (pledgeAsset.Type == claimAsset.Type &&
		pledgeAsset.Owner == claimAsset.Owner &&
		pledgeAsset.NumUnits == claimAsset.NumUnits)
}


// InitTokenAssetLedger adds a base set of assets to the ledger
func (s *SmartContract) InitTokenAssetLedger(ctx contractapi.TransactionContextInterface) error {
	_, err := s.CreateTokenAssetType(ctx, "token1", "CentralBank", 1)
	if err != nil {
		return err
	}
	return err
}

// CreateTokenAssetType issues a new token asset type to the world state with given details.
func (s *SmartContract) CreateTokenAssetType(ctx contractapi.TransactionContextInterface, tokenAssetType string, issuer string, value int) (bool, error) {
	if tokenAssetType == "" {
		return false, fmt.Errorf("Token asset type cannot be blank")
	}
	exists, err := s.TokenAssetTypeExists(ctx, tokenAssetType)
	if err != nil {
		return false, err
	}
	if exists {
		return false, fmt.Errorf("the token asset type %s already exists.", tokenAssetType)
	}

	asset := TokenAssetType{
		Issuer: issuer,
		Value: value,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return false, err
	}
	id := getTokenAssetTypeId(tokenAssetType)
	err = ctx.GetStub().PutState(id, assetJSON)

	if err != nil {
		return false, fmt.Errorf("failed to create token asset type %s. %v", tokenAssetType, err)
	}
	return true, nil
}

// ReadTokenAssetType returns the token asset type stored in the world state with given type.
func (s *SmartContract) ReadTokenAssetType(ctx contractapi.TransactionContextInterface, tokenAssetType string) (*TokenAssetType, error) {
	id := getTokenAssetTypeId(tokenAssetType)
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read token asset type %s: %v", tokenAssetType, err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the token asset type %s does not exist.", tokenAssetType)
	}

	var fat TokenAssetType
	err = json.Unmarshal(assetJSON, &fat)
	if err != nil {
		return nil, err
	}

	return &fat, nil
}

// DeleteTokenAssetType deletes an given token asset type from the world state.
func (s *SmartContract) DeleteTokenAssetType(ctx contractapi.TransactionContextInterface, tokenAssetType string) error {
	exists, err := s.TokenAssetTypeExists(ctx, tokenAssetType)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the token asset type %s does not exist.", tokenAssetType)
	}

	id := getTokenAssetTypeId(tokenAssetType)
	err = ctx.GetStub().DelState(id)
	if err != nil {
		return fmt.Errorf("failed to delete token asset type %s: %v", tokenAssetType, err)
	}
	return nil
}

// TokenAssetTypeExists returns true when token asset type with given ID exists in world state
func (s *SmartContract) TokenAssetTypeExists(ctx contractapi.TransactionContextInterface, tokenAssetType string) (bool, error) {
	id := getTokenAssetTypeId(tokenAssetType)
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// IssueTokenAssets issues new token assets to an owner.
func (s *SmartContract) IssueTokenAssets(ctx contractapi.TransactionContextInterface, tokenAssetType string, numUnits uint64, owner string) error {
	if owner == "" {
		return fmt.Errorf("Owner cannot be blank")
	}

	exists, err := s.TokenAssetTypeExists(ctx, tokenAssetType)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("cannot issue: the token asset type %s does not exist", tokenAssetType)
	}

	id := getWalletId(owner)
	return addTokenAssetsHelper(ctx, tokenAssetType, numUnits, id)
}

// DeleteTokenAssets burns the token assets from an owner.
func (s *SmartContract) DeleteTokenAssets(ctx contractapi.TransactionContextInterface, tokenAssetType string, numUnits uint64) error {
	owner, err := getECertOfTxCreatorBase64(ctx)
	if err != nil {
		return err
	}
	exists, err := s.TokenAssetTypeExists(ctx, tokenAssetType)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the token asset type %s does not exist", tokenAssetType)
	}

	id := getWalletId(owner)
	return subTokenAssetsHelper(ctx, tokenAssetType, numUnits, id)
}

// TransferTokenAssets transfers the token assets from client's account to newOwner
func (s *SmartContract) TransferTokenAssets(ctx contractapi.TransactionContextInterface, tokenAssetType string, numUnits uint64, newOwner string) error {
	exists, err := s.TokenAssetTypeExists(ctx, tokenAssetType)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the token asset type %s does not exist", tokenAssetType)
	}

	if newOwner == "" {
		return fmt.Errorf("New owner cannot be blank")
	}

	owner, err := getECertOfTxCreatorBase64(ctx)
	if err != nil {
		return err
	}

	ownerId := getWalletId(owner)
	newOwnerId := getWalletId(newOwner)

	err = subTokenAssetsHelper(ctx, tokenAssetType, numUnits, ownerId)
	if err != nil {
		return err
	}
	return addTokenAssetsHelper(ctx, tokenAssetType, numUnits, newOwnerId)
}

// PledgeTokenAsset locks an asset for transfer to a different ledger/network.
func (s *SmartContract) PledgeTokenAsset(ctx contractapi.TransactionContextInterface, assetType string, numUnits uint64, remoteNetworkId, recipientCert string, expiryTimeSecs uint64) (string, error) {
	// Verify asset balance for this transaction's client using app-specific-logic
	lockerHasEnoughTokens, err := s.TokenAssetsExist(ctx, assetType, numUnits)
	if err != nil {
		return "", err
	}
	if !lockerHasEnoughTokens {
		return "", fmt.Errorf("cannot pledge token asset of type %s as there are not enough tokens", assetType)
	}

	// Get asset owner (this transaction's client) using app-specific-logic
	owner, err := getECertOfTxCreatorBase64(ctx)
	if err != nil {
		return "", err
	}

	asset := TokenAsset{
		Type: assetType,
		Owner: owner,
        NumUnits: numUnits,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	// Pledge the asset using common (library) logic
    if contractId, err := wutils.PledgeAsset(ctx, assetJSON, assetType, "", numUnits, owner, remoteNetworkId, recipientCert, expiryTimeSecs); err == nil {
		// Deduce asset balance using app-specific logic
		return contractId, s.DeleteTokenAssets(ctx, assetType, numUnits)
	} else {
		return "", err
	}
}

// ClaimRemoteTokenAsset gets ownership of an asset transferred from a different ledger/network.
func (s *SmartContract) ClaimRemoteTokenAsset(ctx contractapi.TransactionContextInterface, assetType, id string, numUnits uint64, owner, remoteNetworkId, pledgeJSON string) error {
	// (Optional) Ensure that this function is being called by the Fabric Interop CC

	// Claim the asset using common (library) logic
	claimer, err := getECertOfTxCreatorBase64(ctx)
	if err != nil {
		return err
	}
	pledgeAssetDetails, err := wutils.ClaimRemoteAsset(ctx, assetType, id, owner, remoteNetworkId, pledgeJSON, claimer)
	if err != nil {
		return err
	}

	// Validate pledged asset details using app-specific-logic
	var asset TokenAsset
	err = json.Unmarshal(pledgeAssetDetails, &asset)
	if err != nil {
		return err
	}
	if asset.NumUnits == 0 {
		return fmt.Errorf("cannot claim asset %s as it has not been pledged in %s", id, remoteNetworkId)
	}
	if asset.Type != assetType {
		return fmt.Errorf("cannot claim asset %s as its type doesn't match the pledge", id)
	}
	if asset.NumUnits != numUnits {
		return fmt.Errorf("cannot claim asset %s as its ID doesn't match the pledge", id)
	}
	if asset.Owner != owner {
		return fmt.Errorf("cannot claim asset %s as it has not been pledged by the given owner", id)
	}

	// Recreate the asset in this network and chaincode using app-specific logic: make the recipient the owner of the asset
	return s.IssueTokenAssets(ctx, assetType, asset.NumUnits, claimer)
}

// ReclaimTokenAsset gets back the ownership of an asset pledged for transfer to a different ledger/network.
func (s *SmartContract) ReclaimTokenAsset(ctx contractapi.TransactionContextInterface, assetType, id, recipientCert, remoteNetworkId, claimStatusJSON string) error {
	// (Optional) Ensure that this function is being called by the Fabric Interop CC

	// Reclaim the asset using common (library) logic
	claimAssetDetails, pledgeAssetDetails, err := wutils.ReclaimAsset(ctx, assetType, id, recipientCert, remoteNetworkId, claimStatusJSON)
	if err != nil {
		return err
	}

	// Validate reclaimed asset details using app-specific-logic
	var claimAsset TokenAsset
	err = json.Unmarshal(claimAssetDetails, &claimAsset)
	if err != nil {
		return err
	}
	if (claimAsset.Type != "" &&
		claimAsset.NumUnits != 0 &&
		claimAsset.Owner != "") {
		// Run checks on the claim parameter to see if it is what we expect and to ensure it has not already been made in the other network
		if !matchClaimWithAssetPledge(pledgeAssetDetails, claimAssetDetails) {
			return fmt.Errorf("claim info for asset %s does not match pledged asset details on ledger: %s", id, pledgeAssetDetails)
		}
	}

	// Recreate the asset in this network and chaincode using app-specific logic
	return s.IssueTokenAssets(ctx, assetType, claimAsset.NumUnits, claimAsset.Owner)
}

// GetTokenAssetPledgeStatus returns the asset pledge status.
func (s *SmartContract) GetTokenAssetPledgeStatus(ctx contractapi.TransactionContextInterface, assetType, id, owner, recipientNetworkId, recipientCert string) (string, error) {
	// (Optional) Ensure that this function is being called by the relay via the Fabric Interop CC

	// Create blank asset details using app-specific-logic
	blankAsset := TokenAsset{
		Type: "",
		Owner: "",
		NumUnits: 0,
	}
	blankAssetJSON, err := json.Marshal(blankAsset)
	if err != nil {
		return "", err
	}

	// Fetch asset pledge details using common (library) logic
	pledgeAssetDetails, pledgeJSON, blankPledgeJSON, err := wutils.GetAssetPledgeStatus(ctx, assetType, id, recipientNetworkId, recipientCert, blankAssetJSON)
	if err != nil {
		return blankPledgeJSON, err
	}
	if pledgeAssetDetails == nil {
		return blankPledgeJSON, err
	}

	// Validate returned asset details using app-specific-logic
	var lookupPledgeAsset TokenAsset
	err = json.Unmarshal(pledgeAssetDetails, &lookupPledgeAsset)
	if err != nil {
		return blankPledgeJSON, err
	}
	if lookupPledgeAsset.Owner != owner {
		return blankPledgeJSON, nil      // Return blank
	}

	return pledgeJSON, nil
}

// TODO: GetTokenAssetClaimStatus returns the asset claim status and present time (of invocation).
/*func (s *SmartContract) GetTokenAssetClaimStatus(ctx contractapi.TransactionContextInterface, assetType, id, recipientCert, pledger, pledgerNetworkId string, pledgeExpiryTimeSecs uint64) (string, error) {
	// (Optional) Ensure that this function is being called by the relay via the Fabric Interop CC

	// Create blank asset details using app-specific-logic
	blankAsset := BondAsset{
		Type: "",
		Owner: "",
		NumUnits: 0,
	}
	blankAssetJSON, err := json.Marshal(blankAsset)
	if err != nil {
		return "", err
	}

	// Fetch asset claim details using common (library) logic
	claimAssetDetails, claimJSON, blankClaimJSON, err := wutils.GetAssetClaimStatus(ctx, assetType, id, recipientCert, pledger, pledgerNetworkId, pledgeExpiryTimeSecs, blankAssetJSON)
	if err != nil {
		return blankClaimJSON, err
	}
	if claimAssetDetails == nil {
		return blankClaimJSON, err
	}

	// Validate returned asset details using app-specific-logic

	// The asset should be recorded if it has been claimed, so we should look that up first
	assetJSON, err := ctx.GetStub().GetState(getBondAssetKey(assetType, id))
	if err != nil {
		return blankClaimJSON, fmt.Errorf("failed to read asset record from world state: %v", err)
	}
	if assetJSON == nil {
		return blankClaimJSON, nil      // Return blank
	}

	// Check if this asset is owned by the given recipient
	var asset BondAsset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return blankClaimJSON, err
	}
	if asset.Owner != recipientCert {
		return blankClaimJSON, nil      // Return blank
	}

	// Match pledger identity in claim with request parameters
	var lookupClaimAsset BondAsset
	err = json.Unmarshal(claimAssetDetails, &lookupClaimAsset)
	if err != nil {
		return blankClaimJSON, err
	}
	if lookupClaimAsset.Owner != pledger {
		return blankClaimJSON, nil      // Return blank
	}

	return claimJSON, nil
}*/

// GetBalance returns the amount of given token asset type owned by an owner.
func (s *SmartContract) GetBalance(ctx contractapi.TransactionContextInterface, tokenAssetType string, owner string) (uint64, error) {
	exists, err := s.TokenAssetTypeExists(ctx, tokenAssetType)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, fmt.Errorf("the token asset type %s does not exist", tokenAssetType)
	}

	id := getWalletId(owner)
	walletJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return 0, fmt.Errorf("failed to read owner's wallet from world state: %v", err)
	}
	if walletJSON == nil {
		return 0, fmt.Errorf("owner does not have a wallet")
	}

	var wallet TokenWallet
	err = json.Unmarshal(walletJSON, &wallet)
	if err != nil {
		return 0, err
	}
	balance := wallet.WalletMap[tokenAssetType]
	return balance, nil
}

// GetMyWallet returns the available amount for each token asset type owned by an owner.
func (s *SmartContract) GetMyWallet(ctx contractapi.TransactionContextInterface) (string, error) {
	owner, err := getECertOfTxCreatorBase64(ctx)
	if err != nil {
		return "", err
	}

	id := getWalletId(owner)
	walletJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", fmt.Errorf("failed to read owner's wallet from world state: %v", err)
	}
	if walletJSON == nil {
		return "", fmt.Errorf("owner does not have a wallet")
	}

	var wallet TokenWallet
	err = json.Unmarshal(walletJSON, &wallet)
	if err != nil {
		return "", err
	}
	return createKeyValuePairs(wallet.WalletMap), nil
}

// Checks if owner has some given amount of token asset
func (s *SmartContract) TokenAssetsExist(ctx contractapi.TransactionContextInterface, tokenAssetType string, numUnits uint64) (bool, error) {
	owner, err := getECertOfTxCreatorBase64(ctx)
	if err != nil {
		return false, err
	}
	balance, err := s.GetBalance(ctx, tokenAssetType, owner)
	if err != nil {
		return false, err
	}
	return balance >= numUnits, nil
}

// Helper Functions for token asset
func addTokenAssetsHelper(ctx contractapi.TransactionContextInterface, tokenAssetType string, numUnits uint64, id string) error {
	walletJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return logThenErrorf("failed to retrieve entry from ledger: %+v", err)
	}
	var wallet TokenWallet
	if walletJSON != nil {
		err = json.Unmarshal(walletJSON, &wallet)
		if err != nil {
			return err
		}
		balance := wallet.WalletMap[tokenAssetType]
		wallet.WalletMap[tokenAssetType] = balance + numUnits
	} else {
		walletMap := make(map[string]uint64)
		walletMap[tokenAssetType] = numUnits
		wallet = TokenWallet{
			WalletMap: walletMap,
		}
	}

	walletNewJSON, err := json.Marshal(wallet)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, walletNewJSON)
}

func subTokenAssetsHelper(ctx contractapi.TransactionContextInterface, tokenAssetType string, numUnits uint64, id string) error {
	walletJSON, err := ctx.GetStub().GetState(id)
	var wallet TokenWallet
	if err != nil {
		return err
	}
	if walletJSON == nil {
		return fmt.Errorf("owner does not have a wallet")
	}

	err = json.Unmarshal(walletJSON, &wallet)
	if err != nil {
		return err
	}

	// Check if owner has sufficient amount of given type to delete
	_, exists := wallet.WalletMap[tokenAssetType]
	if !exists {
		return fmt.Errorf("the owner does not possess any units of the token asset type %s", tokenAssetType)
	}
	if wallet.WalletMap[tokenAssetType] < numUnits {
		return fmt.Errorf("the owner does not possess enough units of the token asset type %s", tokenAssetType)
	}

	// Subtract after all checks
	wallet.WalletMap[tokenAssetType] -= numUnits

	// Delete token asset type from map if num of units becomes zero
	if wallet.WalletMap[tokenAssetType] == 0 {
		delete(wallet.WalletMap, tokenAssetType)
	}

	if len(wallet.WalletMap) == 0 {
		// Delete the entry from State if wallet becomes empty
		return ctx.GetStub().DelState(id)
	} else {
		// Update the new wallet object otherwise
		walletNewJSON, err := json.Marshal(wallet)
		if err != nil {
			return err
		}
		return ctx.GetStub().PutState(id, walletNewJSON)
	}
}

func getTokenAssetTypeId(tokenAssetType string) string {
	return "FAT_" + tokenAssetType
}
func getWalletId(owner string) string {
	return "W_" + owner
}

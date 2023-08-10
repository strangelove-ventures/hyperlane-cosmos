package types

const (
	EventTypeGasDataSet     = "remotegasdataset"
	EventTypeSetGasOverhead = "destinationgasoverheadset"
	EventTypeCreateIgp      = "createigp"
	EventTypeUpdateOracle   = "updateoracle"
	EventTypeCreateOracle   = "createoracle"
	EventTypeSetBeneficiary = "setbeneficiary"

	AttributeIgpId             = "igpid"
	AttributeOracleAddress     = "oracle"
	AttributeBeneficiary       = "beneficiary"
	AttributeRemoteDomain      = "remotedomain"
	AttributeTokenExchangeRate = "tokenexchangerate"
	AttributeGasPrice          = "gasprice"
	AttributeDestination       = "destination"
	AttributeOverheadAmount    = "amount"
	AttributeKeySender         = "sender"
)

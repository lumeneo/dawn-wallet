package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in LMO
	DefaultTransferFee float64 = 0.1
	// DefaultTransferMixin is the default mixin
	DefaultTransferMixin = 3

	logWalletdCurrentSessionFilename     = "walletd-session.log"
	logWalletdAllSessionsFilename        = "walletd.log"
	logLumeneodCurrentSessionFilename    = "lumeneod-session.log"
	logLumeneodAllSessionsFilename       = "Lumeneod.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "walletd"
	lumeneodCommandName                  = "Lumeneod"
)

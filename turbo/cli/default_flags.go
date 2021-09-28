package cli

import (
	"github.com/ledgerwatch/erigon/cmd/utils"

	"github.com/urfave/cli"
)

// DefaultFlags contains all flags that are used and supported by Erigon binary.
var DefaultFlags = []cli.Flag{
	utils.DataDirFlag,
	utils.MdbxAugmentLimitFlag,
	utils.EthashDatasetDirFlag,
	utils.TxPoolV2Flag,
	utils.TxPoolDisableFlag,
	utils.TxPoolLocalsFlag,
	utils.TxPoolNoLocalsFlag,
	utils.TxPoolJournalFlag,
	utils.TxPoolRejournalFlag,
	utils.TxPoolPriceLimitFlag,
	utils.TxPoolPriceBumpFlag,
	utils.TxPoolAccountSlotsFlag,
	utils.TxPoolGlobalSlotsFlag,
	utils.TxPoolAccountQueueFlag,
	utils.TxPoolGlobalQueueFlag,
	utils.TxPoolLifetimeFlag,
	PruneFlag,
	PruneHistoryFlag,
	PruneReceiptFlag,
	PruneTxIndexFlag,
	PruneCallTracesFlag,
	PruneHistoryBeforeFlag,
	PruneReceiptBeforeFlag,
	PruneTxIndexBeforeFlag,
	PruneCallTracesBeforeFlag,
	SnapshotModeFlag,
	SeedSnapshotsFlag,
	SnapshotDatabaseLayoutFlag,
	ExternalSnapshotDownloaderAddrFlag,
	BatchSizeFlag,
	BlockDownloaderWindowFlag,
	DatabaseVerbosityFlag,
	PrivateApiAddr,
	PrivateApiRateLimit,
	EtlBufferSizeFlag,
	TLSFlag,
	TLSCertFlag,
	TLSKeyFlag,
	TLSCACertFlag,
	StateStreamFlag,
	SyncLoopThrottleFlag,
	BadBlockFlag,
	utils.ListenPortFlag,
	utils.ListenPort65Flag,
	utils.NATFlag,
	utils.NoDiscoverFlag,
	utils.DiscoveryV5Flag,
	utils.NetrestrictFlag,
	utils.NodeKeyFileFlag,
	utils.NodeKeyHexFlag,
	utils.DNSDiscoveryFlag,
	utils.StaticPeersFlag,
	utils.MaxPeersFlag,
	utils.ChainFlag,
	utils.VMEnableDebugFlag,
	utils.NetworkIdFlag,
	utils.FakePoWFlag,
	utils.GpoBlocksFlag,
	utils.GpoPercentileFlag,
	utils.InsecureUnlockAllowedFlag,
	utils.MetricsEnabledFlag,
	utils.MetricsEnabledExpensiveFlag,
	utils.MetricsHTTPFlag,
	utils.MetricsPortFlag,
	utils.IdentityFlag,
	utils.CliqueSnapshotCheckpointIntervalFlag,
	utils.CliqueSnapshotInmemorySnapshotsFlag,
	utils.CliqueSnapshotInmemorySignaturesFlag,
	utils.CliqueDataDirFlag,
	utils.MiningEnabledFlag,
	utils.MinerNotifyFlag,
	utils.MinerGasTargetFlag,
	utils.MinerGasLimitFlag,
	utils.MinerEtherbaseFlag,
	utils.MinerExtraDataFlag,
	utils.MinerNoVerfiyFlag,
	utils.MinerSigningKeyFileFlag,
	utils.SentryAddrFlag,
}

// Generated by eos-abigen version master

package eosio

import (
	"encoding/json"

	base "github.com/maxifom/eos-abigen/pkg/base"
)

type AbiHash struct {
	Owner string `json:"owner"`
	Hash  string `json:"hash"`
}
type Activate struct {
	FeatureDigest string `json:"feature_digest"`
}
type Authority struct {
	Threshold uint32                  `json:"threshold"`
	Keys      []KeyWeight             `json:"keys"`
	Accounts  []PermissionLevelWeight `json:"accounts"`
	Waits     []WaitWeight            `json:"waits"`
}
type BidRefund struct {
	Bidder string     `json:"bidder"`
	Amount base.Asset `json:"amount"`
}
type Bidname struct {
	Bidder  string     `json:"bidder"`
	Newname string     `json:"newname"`
	Bid     base.Asset `json:"bid"`
}
type Bidrefund struct {
	Bidder  string `json:"bidder"`
	Newname string `json:"newname"`
}
type BlockHeader struct {
	Timestamp        uint32          `json:"timestamp"`
	Producer         string          `json:"producer"`
	Confirmed        uint16          `json:"confirmed"`
	Previous         string          `json:"previous"`
	TransactionMroot string          `json:"transaction_mroot"`
	ActionMroot      string          `json:"action_mroot"`
	ScheduleVersion  uint32          `json:"schedule_version"`
	NewProducers     json.RawMessage `json:"new_producers"`
}
type BlockSigningAuthorityV0 struct {
	Threshold uint32      `json:"threshold"`
	Keys      []KeyWeight `json:"keys"`
}
type BlockchainParameters struct {
	MaxBlockNetUsage               base.UInt64 `json:"max_block_net_usage"`
	TargetBlockNetUsagePct         uint32      `json:"target_block_net_usage_pct"`
	MaxTransactionNetUsage         uint32      `json:"max_transaction_net_usage"`
	BasePerTransactionNetUsage     uint32      `json:"base_per_transaction_net_usage"`
	NetUsageLeeway                 uint32      `json:"net_usage_leeway"`
	ContextFreeDiscountNetUsageNum uint32      `json:"context_free_discount_net_usage_num"`
	ContextFreeDiscountNetUsageDen uint32      `json:"context_free_discount_net_usage_den"`
	MaxBlockCpuUsage               uint32      `json:"max_block_cpu_usage"`
	TargetBlockCpuUsagePct         uint32      `json:"target_block_cpu_usage_pct"`
	MaxTransactionCpuUsage         uint32      `json:"max_transaction_cpu_usage"`
	MinTransactionCpuUsage         uint32      `json:"min_transaction_cpu_usage"`
	MaxTransactionLifetime         uint32      `json:"max_transaction_lifetime"`
	DeferredTrxExpirationWindow    uint32      `json:"deferred_trx_expiration_window"`
	MaxTransactionDelay            uint32      `json:"max_transaction_delay"`
	MaxInlineActionSize            uint32      `json:"max_inline_action_size"`
	MaxInlineActionDepth           uint16      `json:"max_inline_action_depth"`
	MaxAuthorityDepth              uint16      `json:"max_authority_depth"`
}
type Buyram struct {
	Payer    string     `json:"payer"`
	Receiver string     `json:"receiver"`
	Quant    base.Asset `json:"quant"`
}
type Buyrambytes struct {
	Payer    string `json:"payer"`
	Receiver string `json:"receiver"`
	Bytes    uint32 `json:"bytes"`
}
type Buyrex struct {
	From   string     `json:"from"`
	Amount base.Asset `json:"amount"`
}
type Canceldelay struct {
	CancelingAuth PermissionLevel `json:"canceling_auth"`
	TrxId         string          `json:"trx_id"`
}
type Cfgpowerup struct {
	Args PowerupConfig `json:"args"`
}
type Claimrewards struct {
	Owner string `json:"owner"`
}
type Closerex struct {
	Owner string `json:"owner"`
}
type Cnclrexorder struct {
	Owner string `json:"owner"`
}
type Connector struct {
	Balance base.Asset   `json:"balance"`
	Weight  base.Float64 `json:"weight"`
}
type Consolidate struct {
	Owner string `json:"owner"`
}
type Defcpuloan struct {
	From    string      `json:"from"`
	LoanNum base.UInt64 `json:"loan_num"`
	Amount  base.Asset  `json:"amount"`
}
type Defnetloan struct {
	From    string      `json:"from"`
	LoanNum base.UInt64 `json:"loan_num"`
	Amount  base.Asset  `json:"amount"`
}
type Delegatebw struct {
	From             string     `json:"from"`
	Receiver         string     `json:"receiver"`
	StakeNetQuantity base.Asset `json:"stake_net_quantity"`
	StakeCpuQuantity base.Asset `json:"stake_cpu_quantity"`
	Transfer         base.Bool  `json:"transfer"`
}
type DelegatedBandwidth struct {
	From      string     `json:"from"`
	To        string     `json:"to"`
	NetWeight base.Asset `json:"net_weight"`
	CpuWeight base.Asset `json:"cpu_weight"`
}
type Deleteauth struct {
	Account    string `json:"account"`
	Permission string `json:"permission"`
}
type Deposit struct {
	Owner  string     `json:"owner"`
	Amount base.Asset `json:"amount"`
}
type EosioGlobalState struct {
	MaxRamSize                 base.UInt64             `json:"max_ram_size"`
	TotalRamBytesReserved      base.UInt64             `json:"total_ram_bytes_reserved"`
	TotalRamStake              int64                   `json:"total_ram_stake"`
	LastProducerScheduleUpdate base.BlockTimestampType `json:"last_producer_schedule_update"`
	LastPervoteBucketFill      base.TimePoint          `json:"last_pervote_bucket_fill"`
	PervoteBucket              int64                   `json:"pervote_bucket"`
	PerblockBucket             int64                   `json:"perblock_bucket"`
	TotalUnpaidBlocks          uint32                  `json:"total_unpaid_blocks"`
	TotalActivatedStake        int64                   `json:"total_activated_stake"`
	ThreshActivatedStakeTime   base.TimePoint          `json:"thresh_activated_stake_time"`
	LastProducerScheduleSize   uint16                  `json:"last_producer_schedule_size"`
	TotalProducerVoteWeight    base.Float64            `json:"total_producer_vote_weight"`
	LastNameClose              base.BlockTimestampType `json:"last_name_close"`
}
type EosioGlobalState2 struct {
	NewRamPerBlock            uint16                  `json:"new_ram_per_block"`
	LastRamIncrease           base.BlockTimestampType `json:"last_ram_increase"`
	LastBlockNum              base.BlockTimestampType `json:"last_block_num"`
	TotalProducerVotepayShare base.Float64            `json:"total_producer_votepay_share"`
	Revision                  uint8                   `json:"revision"`
}
type EosioGlobalState3 struct {
	LastVpayStateUpdate      base.TimePoint `json:"last_vpay_state_update"`
	TotalVpayShareChangeRate base.Float64   `json:"total_vpay_share_change_rate"`
}
type EosioGlobalState4 struct {
	ContinuousRate     base.Float64 `json:"continuous_rate"`
	InflationPayFactor int64        `json:"inflation_pay_factor"`
	VotepayFactor      int64        `json:"votepay_factor"`
}
type ExchangeState struct {
	Supply base.Asset `json:"supply"`
	Base   Connector  `json:"base"`
	Quote  Connector  `json:"quote"`
}
type Fundcpuloan struct {
	From    string      `json:"from"`
	LoanNum base.UInt64 `json:"loan_num"`
	Payment base.Asset  `json:"payment"`
}
type Fundnetloan struct {
	From    string      `json:"from"`
	LoanNum base.UInt64 `json:"loan_num"`
	Payment base.Asset  `json:"payment"`
}
type Init struct {
	Version uint32      `json:"version"`
	Core    base.Symbol `json:"core"`
}
type KeyWeight struct {
	Key    string `json:"key"`
	Weight uint16 `json:"weight"`
}
type Linkauth struct {
	Account     string `json:"account"`
	Code        string `json:"code"`
	Type        string `json:"type"`
	Requirement string `json:"requirement"`
}
type Mvfrsavings struct {
	Owner string     `json:"owner"`
	Rex   base.Asset `json:"rex"`
}
type Mvtosavings struct {
	Owner string     `json:"owner"`
	Rex   base.Asset `json:"rex"`
}
type NameBid struct {
	Newname     string         `json:"newname"`
	HighBidder  string         `json:"high_bidder"`
	HighBid     int64          `json:"high_bid"`
	LastBidTime base.TimePoint `json:"last_bid_time"`
}
type Newaccount struct {
	Creator string    `json:"creator"`
	Name    string    `json:"name"`
	Owner   Authority `json:"owner"`
	Active  Authority `json:"active"`
}
type Onblock struct {
	Header BlockHeader `json:"header"`
}
type Onerror struct {
	SenderId base.BigInt `json:"sender_id"`
	SentTrx  string      `json:"sent_trx"`
}
type PairTimePointSecInt64 struct {
	Key   base.TimePointSec `json:"key"`
	Value int64             `json:"value"`
}
type PermissionLevel struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}
type PermissionLevelWeight struct {
	Permission PermissionLevel `json:"permission"`
	Weight     uint16          `json:"weight"`
}
type Powerup struct {
	Payer      string     `json:"payer"`
	Receiver   string     `json:"receiver"`
	Days       uint32     `json:"days"`
	NetFrac    int64      `json:"net_frac"`
	CpuFrac    int64      `json:"cpu_frac"`
	MaxPayment base.Asset `json:"max_payment"`
}
type PowerupConfig struct {
	Net           PowerupConfigResource `json:"net"`
	Cpu           PowerupConfigResource `json:"cpu"`
	PowerupDays   json.RawMessage       `json:"powerup_days"`
	MinPowerupFee json.RawMessage       `json:"min_powerup_fee"`
}
type PowerupConfigResource struct {
	CurrentWeightRatio json.RawMessage `json:"current_weight_ratio"`
	TargetWeightRatio  json.RawMessage `json:"target_weight_ratio"`
	AssumedStakeWeight json.RawMessage `json:"assumed_stake_weight"`
	TargetTimestamp    json.RawMessage `json:"target_timestamp"`
	Exponent           json.RawMessage `json:"exponent"`
	DecaySecs          json.RawMessage `json:"decay_secs"`
	MinPrice           json.RawMessage `json:"min_price"`
	MaxPrice           json.RawMessage `json:"max_price"`
}
type PowerupOrder struct {
	Version   uint8             `json:"version"`
	Id        base.UInt64       `json:"id"`
	Owner     string            `json:"owner"`
	NetWeight int64             `json:"net_weight"`
	CpuWeight int64             `json:"cpu_weight"`
	Expires   base.TimePointSec `json:"expires"`
}
type PowerupState struct {
	Version       uint8                `json:"version"`
	Net           PowerupStateResource `json:"net"`
	Cpu           PowerupStateResource `json:"cpu"`
	PowerupDays   uint32               `json:"powerup_days"`
	MinPowerupFee base.Asset           `json:"min_powerup_fee"`
}
type PowerupStateResource struct {
	Version              uint8             `json:"version"`
	Weight               int64             `json:"weight"`
	WeightRatio          int64             `json:"weight_ratio"`
	AssumedStakeWeight   int64             `json:"assumed_stake_weight"`
	InitialWeightRatio   int64             `json:"initial_weight_ratio"`
	TargetWeightRatio    int64             `json:"target_weight_ratio"`
	InitialTimestamp     base.TimePointSec `json:"initial_timestamp"`
	TargetTimestamp      base.TimePointSec `json:"target_timestamp"`
	Exponent             base.Float64      `json:"exponent"`
	DecaySecs            uint32            `json:"decay_secs"`
	MinPrice             base.Asset        `json:"min_price"`
	MaxPrice             base.Asset        `json:"max_price"`
	Utilization          int64             `json:"utilization"`
	AdjustedUtilization  int64             `json:"adjusted_utilization"`
	UtilizationTimestamp base.TimePointSec `json:"utilization_timestamp"`
}
type Powerupexec struct {
	User string `json:"user"`
	Max  uint16 `json:"max"`
}
type ProducerInfo struct {
	Owner             string          `json:"owner"`
	TotalVotes        base.Float64    `json:"total_votes"`
	ProducerKey       string          `json:"producer_key"`
	IsActive          base.Bool       `json:"is_active"`
	Url               string          `json:"url"`
	UnpaidBlocks      uint32          `json:"unpaid_blocks"`
	LastClaimTime     base.TimePoint  `json:"last_claim_time"`
	Location          uint16          `json:"location"`
	ProducerAuthority json.RawMessage `json:"producer_authority"`
}
type ProducerInfo2 struct {
	Owner                  string         `json:"owner"`
	VotepayShare           base.Float64   `json:"votepay_share"`
	LastVotepayShareUpdate base.TimePoint `json:"last_votepay_share_update"`
}
type ProducerKey struct {
	ProducerName    string `json:"producer_name"`
	BlockSigningKey string `json:"block_signing_key"`
}
type ProducerSchedule struct {
	Version   uint32        `json:"version"`
	Producers []ProducerKey `json:"producers"`
}
type Refund struct {
	Owner string `json:"owner"`
}
type RefundRequest struct {
	Owner       string            `json:"owner"`
	RequestTime base.TimePointSec `json:"request_time"`
	NetAmount   base.Asset        `json:"net_amount"`
	CpuAmount   base.Asset        `json:"cpu_amount"`
}
type Regproducer struct {
	Producer    string `json:"producer"`
	ProducerKey string `json:"producer_key"`
	Url         string `json:"url"`
	Location    uint16 `json:"location"`
}
type Regproducer2 struct {
	Producer          string          `json:"producer"`
	ProducerAuthority json.RawMessage `json:"producer_authority"`
	Url               string          `json:"url"`
	Location          uint16          `json:"location"`
}
type Regproxy struct {
	Proxy   string    `json:"proxy"`
	Isproxy base.Bool `json:"isproxy"`
}
type Rentcpu struct {
	From        string     `json:"from"`
	Receiver    string     `json:"receiver"`
	LoanPayment base.Asset `json:"loan_payment"`
	LoanFund    base.Asset `json:"loan_fund"`
}
type Rentnet struct {
	From        string     `json:"from"`
	Receiver    string     `json:"receiver"`
	LoanPayment base.Asset `json:"loan_payment"`
	LoanFund    base.Asset `json:"loan_fund"`
}
type RexBalance struct {
	Version       uint8                   `json:"version"`
	Owner         string                  `json:"owner"`
	VoteStake     base.Asset              `json:"vote_stake"`
	RexBalance    base.Asset              `json:"rex_balance"`
	MaturedRex    int64                   `json:"matured_rex"`
	RexMaturities []PairTimePointSecInt64 `json:"rex_maturities"`
}
type RexFund struct {
	Version uint8      `json:"version"`
	Owner   string     `json:"owner"`
	Balance base.Asset `json:"balance"`
}
type RexLoan struct {
	Version     uint8          `json:"version"`
	From        string         `json:"from"`
	Receiver    string         `json:"receiver"`
	Payment     base.Asset     `json:"payment"`
	Balance     base.Asset     `json:"balance"`
	TotalStaked base.Asset     `json:"total_staked"`
	LoanNum     base.UInt64    `json:"loan_num"`
	Expiration  base.TimePoint `json:"expiration"`
}
type RexOrder struct {
	Version      uint8          `json:"version"`
	Owner        string         `json:"owner"`
	RexRequested base.Asset     `json:"rex_requested"`
	Proceeds     base.Asset     `json:"proceeds"`
	StakeChange  base.Asset     `json:"stake_change"`
	OrderTime    base.TimePoint `json:"order_time"`
	IsOpen       base.Bool      `json:"is_open"`
}
type RexPool struct {
	Version         uint8       `json:"version"`
	TotalLent       base.Asset  `json:"total_lent"`
	TotalUnlent     base.Asset  `json:"total_unlent"`
	TotalRent       base.Asset  `json:"total_rent"`
	TotalLendable   base.Asset  `json:"total_lendable"`
	TotalRex        base.Asset  `json:"total_rex"`
	NamebidProceeds base.Asset  `json:"namebid_proceeds"`
	LoanNum         base.UInt64 `json:"loan_num"`
}
type RexReturnBuckets struct {
	Version       uint8                   `json:"version"`
	ReturnBuckets []PairTimePointSecInt64 `json:"return_buckets"`
}
type RexReturnPool struct {
	Version               uint8             `json:"version"`
	LastDistTime          base.TimePointSec `json:"last_dist_time"`
	PendingBucketTime     base.TimePointSec `json:"pending_bucket_time"`
	OldestBucketTime      base.TimePointSec `json:"oldest_bucket_time"`
	PendingBucketProceeds int64             `json:"pending_bucket_proceeds"`
	CurrentRateOfIncrease int64             `json:"current_rate_of_increase"`
	Proceeds              int64             `json:"proceeds"`
}
type Rexexec struct {
	User string `json:"user"`
	Max  uint16 `json:"max"`
}
type Rmvproducer struct {
	Producer string `json:"producer"`
}
type Sellram struct {
	Account string `json:"account"`
	Bytes   int64  `json:"bytes"`
}
type Sellrex struct {
	From string     `json:"from"`
	Rex  base.Asset `json:"rex"`
}
type Setabi struct {
	Account string `json:"account"`
	Abi     string `json:"abi"`
}
type Setacctcpu struct {
	Account   string          `json:"account"`
	CpuWeight json.RawMessage `json:"cpu_weight"`
}
type Setacctnet struct {
	Account   string          `json:"account"`
	NetWeight json.RawMessage `json:"net_weight"`
}
type Setacctram struct {
	Account  string          `json:"account"`
	RamBytes json.RawMessage `json:"ram_bytes"`
}
type Setalimits struct {
	Account   string `json:"account"`
	RamBytes  int64  `json:"ram_bytes"`
	NetWeight int64  `json:"net_weight"`
	CpuWeight int64  `json:"cpu_weight"`
}
type Setcode struct {
	Account   string `json:"account"`
	Vmtype    uint8  `json:"vmtype"`
	Vmversion uint8  `json:"vmversion"`
	Code      string `json:"code"`
}
type Setinflation struct {
	AnnualRate         int64 `json:"annual_rate"`
	InflationPayFactor int64 `json:"inflation_pay_factor"`
	VotepayFactor      int64 `json:"votepay_factor"`
}
type Setparams struct {
	Params BlockchainParameters `json:"params"`
}
type Setpriv struct {
	Account string `json:"account"`
	IsPriv  uint8  `json:"is_priv"`
}
type Setram struct {
	MaxRamSize base.UInt64 `json:"max_ram_size"`
}
type Setramrate struct {
	BytesPerBlock uint16 `json:"bytes_per_block"`
}
type Setrex struct {
	Balance base.Asset `json:"balance"`
}
type Undelegatebw struct {
	From               string     `json:"from"`
	Receiver           string     `json:"receiver"`
	UnstakeNetQuantity base.Asset `json:"unstake_net_quantity"`
	UnstakeCpuQuantity base.Asset `json:"unstake_cpu_quantity"`
}
type Unlinkauth struct {
	Account string `json:"account"`
	Code    string `json:"code"`
	Type    string `json:"type"`
}
type Unregprod struct {
	Producer string `json:"producer"`
}
type Unstaketorex struct {
	Owner    string     `json:"owner"`
	Receiver string     `json:"receiver"`
	FromNet  base.Asset `json:"from_net"`
	FromCpu  base.Asset `json:"from_cpu"`
}
type Updateauth struct {
	Account    string    `json:"account"`
	Permission string    `json:"permission"`
	Parent     string    `json:"parent"`
	Auth       Authority `json:"auth"`
}
type Updaterex struct {
	Owner string `json:"owner"`
}
type Updtrevision struct {
	Revision uint8 `json:"revision"`
}
type UserResources struct {
	Owner     string     `json:"owner"`
	NetWeight base.Asset `json:"net_weight"`
	CpuWeight base.Asset `json:"cpu_weight"`
	RamBytes  int64      `json:"ram_bytes"`
}
type Voteproducer struct {
	Voter     string   `json:"voter"`
	Proxy     string   `json:"proxy"`
	Producers []string `json:"producers"`
}
type VoterInfo struct {
	Owner             string       `json:"owner"`
	Proxy             string       `json:"proxy"`
	Producers         []string     `json:"producers"`
	Staked            int64        `json:"staked"`
	LastVoteWeight    base.Float64 `json:"last_vote_weight"`
	ProxiedVoteWeight base.Float64 `json:"proxied_vote_weight"`
	IsProxy           base.Bool    `json:"is_proxy"`
	Flags1            uint32       `json:"flags1"`
	Reserved2         uint32       `json:"reserved2"`
	Reserved3         base.Asset   `json:"reserved3"`
}
type WaitWeight struct {
	WaitSec uint32 `json:"wait_sec"`
	Weight  uint16 `json:"weight"`
}
type Withdraw struct {
	Owner  string     `json:"owner"`
	Amount base.Asset `json:"amount"`
}

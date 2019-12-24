package dto

type NodeStatistic struct {
    ReceivedBlocks int64 `json:blockRecvCnt`

    LastBlockContentSize int64 `json:lastBlockContentSize`
    LastBlockDate string `json:lastBlockDate`
    LastBlockFees int64 `json:lastBlockFees`
    LastBlockHash string `json:lastBlockHash`
    LastBlockHeight string `json:lastBlockHeight`
    LastBlockSum int64 `json:lastBlockSum`
    LastBlockTime string `json:lastBlockTime`
    LastBlockTx  int64 `json:lastBlockTx`

    State string `json:state`
    ReceivedTransactions int64 `json:txRecvCnt`
    UpTime int64 `json:uptime`
    Version string `json:version`
}

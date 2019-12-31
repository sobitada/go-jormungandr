package dto

import "math/big"

type NodeStatistic struct {
    ReceivedBlocks       big.Int `json:"blockRecvCnt"`
    LastBlockContentSize big.Int `json:"lastBlockContentSize"`
    LastBlockDate        string  `json:"lastBlockDate"`
    LastBlockFees        big.Int `json:"lastBlockFees"`
    LastBlockHash        string  `json:"lastBlockHash"`
    LastBlockHeight      string  `json:"lastBlockHeight"`
    LastBlockSum         big.Int `json:"lastBlockSum"`
    LastBlockTime        string  `json:"lastBlockTime"`
    LastBlockTx          big.Int `json:"lastBlockTx"`

    State                string  `json:"state"`
    ReceivedTransactions big.Int `json:"txRecvCnt"`
    UpTime               big.Int `json:"uptime"`
    JormungandrVersion   string  `json:"version"`
}

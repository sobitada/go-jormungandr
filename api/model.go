package api

import (
    "github.com/sobitada/go-cardano"
    "math/big"
    "time"
)

type nodeStatisticJSON struct {
    JormungandrVersion   string   `json:"version"`
    State                string   `json:"state"`
    ReceivedBlocks       *big.Int `json:"blockRecvCnt"`
    ReceivedTransactions *big.Int `json:"txRecvCnt"`
    UpTime               uint32   `json:"uptime"`

    LastBlockContentSize *big.Int               `json:"lastBlockContentSize"`
    LastBlockDate        *cardano.PlainSlotDate `json:"lastBlockDate"`
    LastBlockFees        *big.Int               `json:"lastBlockFees"`
    LastBlockHash        string                 `json:"lastBlockHash"`
    LastBlockHeight      uint64                 `json:"lastBlockHeight,string"`
    LastBlockSum         *big.Int               `json:"lastBlockSum"`
    LastBlockTime        time.Time              `json:"lastBlockTime"`
    LastBlockTx          *big.Int               `json:"lastBlockTx"`

    PeerAvailableCount   *uint64 `json:"peerAvailableCnt"`
    PeerQuarantinedCount *uint64 `json:"peerQuarantinedCnt"`
    PeerUnreachableCnt   *uint64 `json:"peerUnreachableCnt"`
}

type NodeStatistic struct {
    JormungandrVersion   string
    State                string
    ReceivedBlocks       *big.Int
    ReceivedTransactions *big.Int
    UpTime               time.Duration

    LastBlockContentSize *big.Int
    LastBlockDate        *cardano.PlainSlotDate
    LastBlockFees        *big.Int
    LastBlockHash        string
    LastBlockHeight      *big.Int
    LastBlockSum         *big.Int
    LastBlockTime        time.Time
    LastBlockTx          *big.Int

    PeerAvailableCount   *uint64
    PeerQuarantinedCount *uint64
    PeerUnreachableCnt   *uint64
}

func transformJSONToNodeStatistic(nodeStatJSON nodeStatisticJSON) *NodeStatistic {
    return &NodeStatistic{
        JormungandrVersion:   nodeStatJSON.JormungandrVersion,
        State:                nodeStatJSON.State,
        ReceivedBlocks:       nodeStatJSON.ReceivedBlocks,
        ReceivedTransactions: nodeStatJSON.ReceivedTransactions,
        UpTime:               time.Duration(int64(nodeStatJSON.UpTime)) * time.Second,
        LastBlockContentSize: nodeStatJSON.LastBlockContentSize,
        LastBlockDate:        nodeStatJSON.LastBlockDate,
        LastBlockFees:        nodeStatJSON.LastBlockFees,
        LastBlockHash:        nodeStatJSON.LastBlockHash,
        LastBlockHeight:      new(big.Int).SetUint64(nodeStatJSON.LastBlockHeight),
        LastBlockSum:         nodeStatJSON.LastBlockSum,
        LastBlockTime:        nodeStatJSON.LastBlockTime,
        LastBlockTx:          nodeStatJSON.LastBlockTx,
        PeerAvailableCount:   nodeStatJSON.PeerAvailableCount,
        PeerQuarantinedCount: nodeStatJSON.PeerQuarantinedCount,
        PeerUnreachableCnt:   nodeStatJSON.PeerUnreachableCnt,
    }
}

type LeaderAssignment struct {
    CreationTime      time.Time              `json:"created_at_time"`
    ScheduleTime      time.Time              `json:"scheduled_at_time"`
    ScheduleBlockDate *cardano.PlainSlotDate `json:"scheduled_at_date"`
    FinishingTime     time.Time              `json:"finished_at_time"`
    LeaderID          uint64                 `json:"enclave_leader_id"`
}

type LeaderCertificate struct {
    Genesis struct {
        SigKey string `yaml:"sig_key"`
        VrfKey string `yaml:"vrf_key"`
        NodeID string `yaml:"node_id"`
    } `yaml:"genesis"`
}

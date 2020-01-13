package api

import (
    "encoding/json"
    "github.com/sobitada/go-cardano"
    "gopkg.in/yaml.v2"
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
    LastBlockHeight      uint64                 `json:"lastBlockHeight,string"` //TODO: needs to be fixed eventually.
    LastBlockSum         *big.Int               `json:"lastBlockSum"`
    LastBlockTime        time.Time              `json:"lastBlockTime"`
    LastBlockTx          *big.Int               `json:"lastBlockTx"`
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
    }
}

type LeaderAssignment struct {
    CreationTime  time.Time `json:"created_at_time"`
    ScheduleTime  time.Time `json:"scheduled_at_time"`
    FinishingTime time.Time `json:"finished_at_time"`
    LeaderID      uint64    `json:"enclave_leader_id"`
}

type LeaderCertificate struct {
    Genesis struct {
        SigKey string `yaml:"sig_key"`
        VrfKey string `yaml:"vrf_key"`
        NodeID string `yaml:"node_id"`
    } `yaml:"genesis"`
}

type leaderCertificateJSON struct {
    Genesis leaderCertificateJSONGenesis `json:"genesis"`
}

type leaderCertificateJSONGenesis struct {
    SigKey string `json:"sig_key"`
    VrfKey string `json:"vrf_key"`
    NodeID string `json:"node_id"`
}

func ReadLeaderCertificate(data []byte) (LeaderCertificate, error) {
    var leaderCert LeaderCertificate
    err := yaml.Unmarshal(data, &leaderCert)
    return leaderCert, err
}

func CertToJSON(certificate LeaderCertificate) ([]byte, error) {
    leaderCertificate := leaderCertificateJSON{Genesis: leaderCertificateJSONGenesis{SigKey: certificate.Genesis.SigKey, VrfKey: certificate.Genesis.VrfKey, NodeID: certificate.Genesis.NodeID}}
    return json.Marshal(leaderCertificate)
}

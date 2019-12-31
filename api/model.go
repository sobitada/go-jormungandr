package api

import (
    "github.com/sobitada/go-jormungandr/api/dto"
    "github.com/sobitada/go-jormungandr/cardano"
    "math/big"
    "time"
)

type NodeStatistic struct {
    ReceivedBlocks       big.Int
    ReceivedTransactions big.Int
    LastBlockSlotDate    cardano.PlainSlotDate
    LastBlockTime        time.Time
    LastBlockHeight      big.Int
    LastBlockHash        string
    UpTime               time.Duration

    LastBlockContentSize big.Int
    LastBlockFees        big.Int
    LastBlockSum         big.Int

    JormungandrVersion string
}

func getNodeStatisticFromDto(dto dto.NodeStatistic) (*NodeStatistic, error) {
    lastBlockTime, err := time.Parse(time.RFC3339, dto.LastBlockTime)
    if err != nil {
        return nil, err
    }
    lastBlockSlotDate, err := cardano.ParsePlainData(dto.LastBlockDate)
    if err != nil {
        return nil, err
    }
    lastBlockHeight, success := new(big.Int).SetString(dto.LastBlockHeight, 10)
    if !success {
        return nil, err
    }
    return &NodeStatistic{
        ReceivedBlocks:       dto.ReceivedBlocks,
        ReceivedTransactions: dto.ReceivedTransactions,
        LastBlockSlotDate:    lastBlockSlotDate,
        LastBlockTime:        lastBlockTime,
        LastBlockHeight:      *lastBlockHeight,
        LastBlockHash:        dto.LastBlockHash,
        UpTime:               time.Second * time.Duration(dto.UpTime.Uint64()),

        LastBlockContentSize: dto.LastBlockContentSize,
        LastBlockFees:        dto.LastBlockFees,
        LastBlockSum:         dto.LastBlockSum,

        JormungandrVersion: dto.JormungandrVersion,
    }, nil
}

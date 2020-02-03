package api

import (
    "encoding/json"
    "gopkg.in/yaml.v2"
)

type leaderCertificateJSON struct {
    Genesis leaderCertificateJSONGenesis `json:"genesis"`
}

type leaderCertificateJSONGenesis struct {
    SigKey string `json:"sig_key"`
    VrfKey string `json:"vrf_key"`
    NodeID string `json:"node_id"`
}

// reads the genesis leader certificate from the given byte stream.
func ReadLeaderCertificate(data []byte) (LeaderCertificate, error) {
    var leaderCert LeaderCertificate
    err := yaml.Unmarshal(data, &leaderCert)
    return leaderCert, err
}

// transforms the given certificate into a JSON.
func certificateToJSON(certificate LeaderCertificate) ([]byte, error) {
    leaderCertificate := leaderCertificateJSON{Genesis: leaderCertificateJSONGenesis{SigKey: certificate.Genesis.SigKey, VrfKey: certificate.Genesis.VrfKey, NodeID: certificate.Genesis.NodeID}}
    return json.Marshal(leaderCertificate)
}

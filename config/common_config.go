// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/staking"
	"github.com/spf13/viper"
)

const TLSCertPathKey = "tls-cert-path"
const TLSKeyPathKey = "tls-key-path"

func GetTLSCertFromFile(v *viper.Viper) (*tls.Certificate, error) {
	if !v.IsSet(TLSKeyPathKey) || !v.IsSet(TLSCertPathKey) {
		return nil, fmt.Errorf("TLS key or cert path not set")
	}
	// Parse the staking key/cert paths and expand environment variables
	keyPath := getExpandedPath(v, TLSKeyPathKey)
	certPath := getExpandedPath(v, TLSCertPathKey)

	var keyMissing, certMissing bool

	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		keyMissing = true
	}
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		certMissing = true
	}
	if keyMissing != certMissing {
		// If only one of the key/cert pair is missing return an error
		// otherwise, create the staking key/cert pair
		return nil, fmt.Errorf("TLS key or cert file is missing from configured path.")
	} else if keyMissing && certMissing {
		// Create the key/cert pair if [TLSKeyPath] and [TLSCertPath] are set but the files are missing
		if err := staking.InitNodeStakingKeyPair(keyPath, certPath); err != nil {
			return nil, fmt.Errorf("couldn't generate TLS key/cert: %w", err)
		}
	}

	// Load and parse the staking key/cert
	cert, err := staking.LoadTLSCertFromFiles(keyPath, certPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't read staking certificate: %w", err)
	}
	return cert, nil
}

// getExpandedPath gets the string in viper corresponding to [key] and expands
// any variables using the OS env.
func getExpandedPath(v *viper.Viper, key string) string {
	return os.Expand(
		v.GetString(key),
		func(strVar string) string {
			return os.Getenv(strVar)
		},
	)
}

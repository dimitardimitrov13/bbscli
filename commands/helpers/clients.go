package helpers

import (
	"strings"
	"time"

	"code.cloudfoundry.org/bbs"
	"github.com/spf13/cobra"
)

const (
	clientSessionCacheSize int = -1
	maxIdleConnsPerHost    int = -1
)

type TLSConfig struct {
	BBSUrl            string
	LocketApiLocation string
	CACertFile        string
	CertFile          string
	KeyFile           string
	SkipCertVerify    bool
	Timeout           int
}

func NewBBSClient(cmd *cobra.Command, bbsClientConfig TLSConfig) (bbs.Client, error) {
	var err error
	var client bbs.Client

	if !strings.HasPrefix(bbsClientConfig.BBSUrl, "https") {
		client, err = bbs.NewClientWithConfig(bbs.ClientConfig{
			URL:            bbsClientConfig.BBSUrl,
			Retries:        1,
			RequestTimeout: time.Duration(bbsClientConfig.Timeout) * time.Second,
		})
	} else {
		client, err = bbs.NewClientWithConfig(
			bbs.ClientConfig{
				URL:                    bbsClientConfig.BBSUrl,
				IsTLS:                  true,
				InsecureSkipVerify:     bbsClientConfig.SkipCertVerify,
				CAFile:                 bbsClientConfig.CACertFile,
				CertFile:               bbsClientConfig.CertFile,
				KeyFile:                bbsClientConfig.KeyFile,
				ClientSessionCacheSize: clientSessionCacheSize,
				MaxIdleConnsPerHost:    maxIdleConnsPerHost,
				Retries:                1,
				RequestTimeout:         time.Duration(bbsClientConfig.Timeout) * time.Second,
			},
		)
	}

	return client, err
}

func (config *TLSConfig) Merge(newConfig TLSConfig) {
	if newConfig.BBSUrl != "" {
		config.BBSUrl = newConfig.BBSUrl
	}
	if newConfig.LocketApiLocation != "" {
		config.LocketApiLocation = newConfig.LocketApiLocation
	}
	if newConfig.Timeout != 0 {
		config.Timeout = newConfig.Timeout
	}
	if newConfig.KeyFile != "" {
		config.KeyFile = newConfig.KeyFile
	}
	if newConfig.CACertFile != "" {
		config.CACertFile = newConfig.CACertFile
	}
	if newConfig.CertFile != "" {
		config.CertFile = newConfig.CertFile
	}
	config.SkipCertVerify = config.SkipCertVerify || newConfig.SkipCertVerify
}

// Copyright 2020 Authors of Hubble
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conn

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cilium/hubble/pkg/defaults"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func init() {
	GRPCOptionFuncs = append(GRPCOptionFuncs, grpcOptionTLS)
}

func grpcOptionTLS(vp *viper.Viper) (grpc.DialOption, error) {
	target := vp.GetString("server")
	if !(vp.GetBool("tls") || strings.HasPrefix(target, defaults.TargetTLSPrefix)) {
		return grpc.WithInsecure(), nil
	}

	tlsConfig := tls.Config{
		InsecureSkipVerify: vp.GetBool("tls-allow-insecure"),
		ServerName:         vp.GetString("tls-server-name"),
	}

	// optional custom CAs
	caFiles := vp.GetStringSlice("tls-ca-cert-files")
	if len(caFiles) > 0 {
		ca := x509.NewCertPool()
		for _, path := range caFiles {
			certPEM, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("cannot load cert '%s': %s", path, err)
			}
			if ok := ca.AppendCertsFromPEM(certPEM); !ok {
				return nil, fmt.Errorf("cannot process cert '%s': must be a PEM encoded certificate", path)
			}
		}
		tlsConfig.RootCAs = ca
	}

	// optional mTLS
	clientCertFile := vp.GetString("tls-client-cert-file")
	clientKeyFile := vp.GetString("tls-client-key-file")
	var cert *tls.Certificate
	if clientCertFile != "" && clientKeyFile != "" {
		c, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
		if err != nil {
			return nil, err
		}
		cert = &c
	}
	tlsConfig.GetClientCertificate = func(_ *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		if cert == nil {
			return nil, errors.New("mTLS client certificate requested, but not provided")
		}
		return cert, nil
	}

	creds := credentials.NewTLS(&tlsConfig)
	return grpc.WithTransportCredentials(creds), nil
}

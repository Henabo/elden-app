package utils

import (
	"crypto/x509"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"path"
)

// NewGrpcConnection creates a gRPC connection to the Gateway server.
func NewGrpcConnection(tlsCertPath string, gatewayPeer string, peerEndpoint string) *grpc.ClientConn {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		log.Panicln(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		log.Panicf("failed to create gRPC connection: %v", err)
	}

	return connection
}

// NewIDentity creates a client IDentity for this gateway connection using x509 certificate.
func NewIDentity(certPath string, mspID string) *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		log.Panicln(err)
	}

	ID, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		log.Panicln(err)
	}

	return ID
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// NewSign generates a digital signature from a message digest using a private key
func NewSign(keyPath string) identity.Sign {
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		log.Panicf("failed to read private key directory: %v", err)
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(keyPath, files[0].Name()))
	if err != nil {
		log.Panicf("failed to read private key file: %v", err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		log.Panicln(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		log.Panicln(err)
	}

	return sign
}

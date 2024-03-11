package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/tls"
    "crypto/x509"
    "crypto/x509/pkix"
    "encoding/pem"
    "math/big"
    "time"

    "github.com/caddyserver/certmagic"
)

func tlsConfigGeneration(hostname string) (*tls.Config, error) {
    private, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, err
    }
    serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
    if err != nil {
        return nil, err
    }
    template := x509.Certificate{
        SerialNumber: serialNumber,
        Subject:      pkix.Name{
            Organization: []string{hostname},
        },
        NotBefore:    time.Now(),
        NotAfter:     time.Now().Add(10 * 365 * 24 * time.Hour),
        KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
        ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
    }
    crtDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &private.PublicKey, private)
    if err != nil {
        return nil, err
    }
    crtPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: crtDER})
    keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(private)})
    tlsConfig := &tls.Config{Certificates: []tls.Certificate{tls.X509KeyPair(crtPEM, keyPEM)}}
    return tlsConfig, nil
}

func tlsConfigApplication(hostname string) (*tls.Config, error) {
    certmagic.DefaultACME.CA = certmagic.LetsEncryptProductionCA
    certmagic.DefaultACME.Agreed = true
    certmagic.DefaultACME.Email = "cert@" + hostname
    tlsConfig, err := certmagic.TLS([]string{hostname})
    return tlsConfig, err
}
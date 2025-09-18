package tipos

import "encoding/xml"

// CryptoBinary is a base64-encoded binary value.
type CryptoBinary []byte

type Signature struct {
	XMLName        xml.Name       `xml:"ds:Signature"`
	SignedInfo     SignedInfo     `xml:"ds:SignedInfo"`
	SignatureValue SignatureValue `xml:"ds:SignatureValue"`
	KeyInfo        *KeyInfo       `xml:"ds:KeyInfo,omitempty"`
	Object         []Object       `xml:"ds:Object,omitempty"`
	Id             string         `xml:"Id,attr,omitempty"`
}

type SignatureValue struct {
	Value []byte `xml:",chardata"`
	Id    string `xml:"Id,attr,omitempty"`
}

type SignedInfo struct {
	CanonicalizationMethod CanonicalizationMethod `xml:"ds:CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"ds:SignatureMethod"`
	Reference              []Reference            `xml:"ds:Reference"`
	Id                     string                 `xml:"Id,attr,omitempty"`
}

type CanonicalizationMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
	Any       []any  `xml:",any"`
}

type SignatureMethod struct {
	Algorithm        string            `xml:"Algorithm,attr"`
	HMACOutputLength *HMACOutputLength `xml:"ds:HMACOutputLength,omitempty"`
	Any              []any             `xml:",any"`
}

type HMACOutputLength int64

type Reference struct {
	Transforms   *Transforms  `xml:"ds:Transforms,omitempty"`
	DigestMethod DigestMethod `xml:"ds:DigestMethod"`
	DigestValue  DigestValue  `xml:"ds:DigestValue"`
	Id           string       `xml:"Id,attr,omitempty"`
	URI          string       `xml:"URI,attr,omitempty"`
	Type         string       `xml:"Type,attr,omitempty"`
}

type Transforms struct {
	Transform []Transform `xml:"ds:Transform"`
}

type Transform struct {
	Algorithm string   `xml:"Algorithm,attr"`
	XPath     []string `xml:"ds:XPath,omitempty"`
	Any       []any    `xml:",any"`
}

type DigestMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
	Any       []any  `xml:",any"`
}

type DigestValue []byte

type KeyInfo struct {
	Id              string            `xml:"Id,attr,omitempty"`
	KeyName         []string          `xml:"ds:KeyName,omitempty"`
	KeyValue        []KeyValue        `xml:"ds:KeyValue,omitempty"`
	RetrievalMethod []RetrievalMethod `xml:"ds:RetrievalMethod,omitempty"`
	X509Data        []X509Data        `xml:"ds:X509Data,omitempty"`
	PGPData         []PGPData         `xml:"ds:PGPData,omitempty"`
	SPKIData        []SPKIData        `xml:"ds:SPKIData,omitempty"`
	MgmtData        []string          `xml:"ds:MgmtData,omitempty"`
	Any             []any             `xml:",any"`
}

type KeyValue struct {
	DSAKeyValue *DSAKeyValue `xml:"ds:DSAKeyValue,omitempty"`
	RSAKeyValue *RSAKeyValue `xml:"ds:RSAKeyValue,omitempty"`
	Any         []any        `xml:",any"`
}

type RetrievalMethod struct {
	Transforms *Transforms `xml:"ds:Transforms,omitempty"`
	URI        string      `xml:"URI,attr,omitempty"`
	Type       string      `xml:"Type,attr,omitempty"`
}

type X509Data struct {
	X509IssuerSerial []X509IssuerSerial `xml:"ds:X509IssuerSerial,omitempty"`
	X509SKI          [][]byte           `xml:"ds:X509SKI,omitempty"`
	X509SubjectName  []string           `xml:"ds:X509SubjectName,omitempty"`
	X509Certificate  [][]byte           `xml:"ds:X509Certificate,omitempty"`
	X509CRL          [][]byte           `xml:"ds:X509CRL,omitempty"`
	Any              []any              `xml:",any"`
}

type X509IssuerSerial struct {
	X509IssuerName   string `xml:"ds:X509IssuerName"`
	X509SerialNumber int64  `xml:"ds:X509SerialNumber"`
}

type PGPData struct {
	PGPKeyID     []byte `xml:"ds:PGPKeyID,omitempty"`
	PGPKeyPacket []byte `xml:"ds:PGPKeyPacket,omitempty"`
	Any          []any  `xml:",any"`
}

type SPKIData struct {
	SPKISexp [][]byte `xml:"ds:SPKISexp,omitempty"`
	Any      []any    `xml:",any"`
}

type Object struct {
	Id       string `xml:"Id,attr,omitempty"`
	MimeType string `xml:"MimeType,attr,omitempty"`
	Encoding string `xml:"Encoding,attr,omitempty"`
	Any      []any  `xml:",any"`
}

type Manifest struct {
	Reference []Reference `xml:"ds:Reference"`
	Id        string      `xml:"Id,attr,omitempty"`
}

type SignatureProperties struct {
	SignatureProperty []SignatureProperty `xml:"ds:SignatureProperty"`
	Id                string              `xml:"Id,attr,omitempty"`
}

type SignatureProperty struct {
	Target string `xml:"Target,attr"`
	Id     string `xml:"Id,attr,omitempty"`
	Any    []any  `xml:",any"`
}

type DSAKeyValue struct {
	P           []byte `xml:"ds:P,omitempty"`
	Q           []byte `xml:"ds:Q,omitempty"`
	G           []byte `xml:"ds:G,omitempty"`
	Y           []byte `xml:"ds:Y"`
	J           []byte `xml:"ds:J,omitempty"`
	Seed        []byte `xml:"ds:Seed,omitempty"`
	PgenCounter []byte `xml:"ds:PgenCounter,omitempty"`
}

type RSAKeyValue struct {
	Modulus  []byte `xml:"ds:Modulus"`
	Exponent []byte `xml:"ds:Exponent"`
}

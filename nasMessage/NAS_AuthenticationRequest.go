// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/huahuaLover/nas/nasType"
		// NOTE: 导包
	"github.com/sirupsen/logrus"
)

type AuthenticationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationRequestMessageIdentity
	nasType.SpareHalfOctetAndNgksi
	nasType.ABBA
	*nasType.AuthenticationParameterRAND
	*nasType.AuthenticationParameterAUTN
	*nasType.EAPMessage
	// NOTE: Add SNMAC in the authentication request
	*nasType.AuthenticationParameterSNMAC
}

func NewAuthenticationRequest(iei uint8) (authenticationRequest *AuthenticationRequest) {
	authenticationRequest = &AuthenticationRequest{}
	return authenticationRequest
}

const (
	AuthenticationRequestAuthenticationParameterRANDType  uint8 = 0x21
	AuthenticationRequestAuthenticationParameterAUTNType  uint8 = 0x20
	AuthenticationRequestEAPMessageType                   uint8 = 0x78
	// NOTE: Add IEI for SNMAC
	AuthenticationRequestAuthenticationParameterSNMACType uint8 = 0x99
)

func (a *AuthenticationRequest) EncodeAuthenticationRequest(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationRequest/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationRequest/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationRequestMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationRequest/AuthenticationRequestMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationRequest/SpareHalfOctetAndNgksi): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationRequest/ABBA): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ABBA.Buffer); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationRequest/ABBA): %w", err)
	}
	if a.AuthenticationParameterRAND != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterRAND.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/AuthenticationParameterRAND): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterRAND.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/AuthenticationParameterRAND): %w", err)
		}
	}
	if a.AuthenticationParameterAUTN != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/AuthenticationParameterAUTN): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/AuthenticationParameterAUTN): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/AuthenticationParameterAUTN): %w", err)
		}
	}
	// NOTE: Include encoding for SNMAC in the authentication request message
	if a.AuthenticationParameterSNMAC != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterSNMAC.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/AuthenticationParameterSNMAC): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterSNMAC.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/AuthenticationParameterSNMAC): %w", err)
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationRequest/EAPMessage): %w", err)
		}
	}
	return nil
}

func (a *AuthenticationRequest) DecodeAuthenticationRequest(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationRequest/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationRequest/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationRequestMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationRequest/AuthenticationRequestMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationRequest/SpareHalfOctetAndNgksi): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ABBA.Len); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationRequest/ABBA): %w", err)
	}
	if a.ABBA.Len < 2 {
		return fmt.Errorf("invalid ie length (AuthenticationRequest/ABBA): %d", a.ABBA.Len)
	}
	a.ABBA.SetLen(a.ABBA.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.ABBA.Buffer); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationRequest/ABBA): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (AuthenticationRequest/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			if ieiN == 0x99{
				tmpIeiN = ieiN
			} else {
				tmpIeiN = (ieiN & 0xf0) >> 4
			}
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case AuthenticationRequestAuthenticationParameterRANDType:
			logrus.Infof("AuthenticationParameterRAND called")
			a.AuthenticationParameterRAND = nasType.NewAuthenticationParameterRAND(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, a.AuthenticationParameterRAND.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationRequest/AuthenticationParameterRAND): %w", err)
			}
		//lihaotian:decode snmac
		case AuthenticationRequestAuthenticationParameterSNMACType:
			logrus.Infof("AuthenticationParameterSNMAC called")
			a.AuthenticationParameterSNMAC = nasType.NewAuthenticationParameterSNMAC(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, a.AuthenticationParameterSNMAC.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationRequest/AuthenticationParameterSNMAC): %w", err)
			}
		case AuthenticationRequestAuthenticationParameterAUTNType:
			logrus.Infof("AuthenticationParameterAUTN called")
			a.AuthenticationParameterAUTN = nasType.NewAuthenticationParameterAUTN(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationParameterAUTN.Len); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationRequest/AuthenticationParameterAUTN): %w", err)
			}
			//lihaotian: update from 16 to 8
			if a.AuthenticationParameterAUTN.Len != 8 {
				return fmt.Errorf("invalid ie length (AuthenticationRequest/AuthenticationParameterAUTN): %d", a.AuthenticationParameterAUTN.Len)
			}
			a.AuthenticationParameterAUTN.SetLen(a.AuthenticationParameterAUTN.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationRequest/AuthenticationParameterAUTN): %w", err)
			}
		case AuthenticationRequestEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationRequest/EAPMessage): %w", err)
			}
			if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
				return fmt.Errorf("invalid ie length (AuthenticationRequest/EAPMessage): %d", a.EAPMessage.Len)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationRequest/EAPMessage): %w", err)
			}
		default:
		}
	}
	return nil
}

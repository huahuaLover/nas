// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/machi12/nas/nasType"
)

type SecurityModeReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.SecurityModeRejectMessageIdentity
	nasType.Cause5GMM
}

func NewSecurityModeReject(iei uint8) (securityModeReject *SecurityModeReject) {
	securityModeReject = &SecurityModeReject{}
	return securityModeReject
}

func (a *SecurityModeReject) EncodeSecurityModeReject(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeReject/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SecurityModeRejectMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeReject/SecurityModeRejectMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.Cause5GMM.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeReject/Cause5GMM): %w", err)
	}
	return nil
}

func (a *SecurityModeReject) DecodeSecurityModeReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeReject/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SecurityModeRejectMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeReject/SecurityModeRejectMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeReject/Cause5GMM): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (SecurityModeReject/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		default:
		}
	}
	return nil
}
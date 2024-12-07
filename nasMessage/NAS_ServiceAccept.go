// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/machi12/nas/nasType"
)

type ServiceAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ServiceAcceptMessageIdentity
	*nasType.PDUSessionStatus
	*nasType.PDUSessionReactivationResult
	*nasType.PDUSessionReactivationResultErrorCause
	*nasType.EAPMessage
}

func NewServiceAccept(iei uint8) (serviceAccept *ServiceAccept) {
	serviceAccept = &ServiceAccept{}
	return serviceAccept
}

const (
	ServiceAcceptPDUSessionStatusType                       uint8 = 0x50
	ServiceAcceptPDUSessionReactivationResultType           uint8 = 0x26
	ServiceAcceptPDUSessionReactivationResultErrorCauseType uint8 = 0x72
	ServiceAcceptEAPMessageType                             uint8 = 0x78
)

func (a *ServiceAccept) EncodeServiceAccept(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ServiceAccept/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ServiceAccept/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ServiceAcceptMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ServiceAccept/ServiceAcceptMessageIdentity): %w", err)
	}
	if a.PDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionStatus): %w", err)
		}
	}
	if a.PDUSessionReactivationResult != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionReactivationResult): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionReactivationResult): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionReactivationResult): %w", err)
		}
	}
	if a.PDUSessionReactivationResultErrorCause != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionReactivationResultErrorCause): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionReactivationResultErrorCause): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/PDUSessionReactivationResultErrorCause): %w", err)
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ServiceAccept/EAPMessage): %w", err)
		}
	}
	return nil
}

func (a *ServiceAccept) DecodeServiceAccept(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ServiceAccept/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ServiceAccept/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ServiceAcceptMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ServiceAccept/ServiceAcceptMessageIdentity): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (ServiceAccept/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case ServiceAcceptPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (ServiceAccept/PDUSessionStatus): %w", err)
			}
			if a.PDUSessionStatus.Len < 2 || a.PDUSessionStatus.Len > 32 {
				return fmt.Errorf("invalid ie length (ServiceAccept/PDUSessionStatus): %d", a.PDUSessionStatus.Len)
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ServiceAccept/PDUSessionStatus): %w", err)
			}
		case ServiceAcceptPDUSessionReactivationResultType:
			a.PDUSessionReactivationResult = nasType.NewPDUSessionReactivationResult(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResult.Len); err != nil {
				return fmt.Errorf("NAS decode error (ServiceAccept/PDUSessionReactivationResult): %w", err)
			}
			if a.PDUSessionReactivationResult.Len < 2 || a.PDUSessionReactivationResult.Len > 32 {
				return fmt.Errorf("invalid ie length (ServiceAccept/PDUSessionReactivationResult): %d", a.PDUSessionReactivationResult.Len)
			}
			a.PDUSessionReactivationResult.SetLen(a.PDUSessionReactivationResult.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResult.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ServiceAccept/PDUSessionReactivationResult): %w", err)
			}
		case ServiceAcceptPDUSessionReactivationResultErrorCauseType:
			a.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResultErrorCause.Len); err != nil {
				return fmt.Errorf("NAS decode error (ServiceAccept/PDUSessionReactivationResultErrorCause): %w", err)
			}
			if a.PDUSessionReactivationResultErrorCause.Len < 2 || a.PDUSessionReactivationResultErrorCause.Len > 512 {
				return fmt.Errorf("invalid ie length (ServiceAccept/PDUSessionReactivationResultErrorCause): %d", a.PDUSessionReactivationResultErrorCause.Len)
			}
			a.PDUSessionReactivationResultErrorCause.SetLen(a.PDUSessionReactivationResultErrorCause.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ServiceAccept/PDUSessionReactivationResultErrorCause): %w", err)
			}
		case ServiceAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (ServiceAccept/EAPMessage): %w", err)
			}
			if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
				return fmt.Errorf("invalid ie length (ServiceAccept/EAPMessage): %d", a.EAPMessage.Len)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ServiceAccept/EAPMessage): %w", err)
			}
		default:
		}
	}
	return nil
}

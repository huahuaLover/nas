// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/machi12/nas/nasType"
)

type SecurityModeCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.SecurityModeCommandMessageIdentity
	nasType.SelectedNASSecurityAlgorithms
	nasType.SpareHalfOctetAndNgksi
	nasType.ReplayedUESecurityCapabilities
	*nasType.IMEISVRequest
	*nasType.SelectedEPSNASSecurityAlgorithms
	*nasType.Additional5GSecurityInformation
	*nasType.EAPMessage
	*nasType.ABBA
	*nasType.ReplayedS1UESecurityCapabilities
}

func NewSecurityModeCommand(iei uint8) (securityModeCommand *SecurityModeCommand) {
	securityModeCommand = &SecurityModeCommand{}
	return securityModeCommand
}

const (
	SecurityModeCommandIMEISVRequestType                    uint8 = 0x0E
	SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType uint8 = 0x57
	SecurityModeCommandAdditional5GSecurityInformationType  uint8 = 0x36
	SecurityModeCommandEAPMessageType                       uint8 = 0x78
	SecurityModeCommandABBAType                             uint8 = 0x38
	SecurityModeCommandReplayedS1UESecurityCapabilitiesType uint8 = 0x19
)

func (a *SecurityModeCommand) EncodeSecurityModeCommand(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeCommand/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeCommand/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SecurityModeCommandMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeCommand/SecurityModeCommandMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SelectedNASSecurityAlgorithms.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeCommand/SelectedNASSecurityAlgorithms): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeCommand/SpareHalfOctetAndNgksi): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ReplayedUESecurityCapabilities.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeCommand/ReplayedUESecurityCapabilities): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ReplayedUESecurityCapabilities.Buffer); err != nil {
		return fmt.Errorf("NAS encode error (SecurityModeCommand/ReplayedUESecurityCapabilities): %w", err)
	}
	if a.IMEISVRequest != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.IMEISVRequest.Octet); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/IMEISVRequest): %w", err)
		}
	}
	if a.SelectedEPSNASSecurityAlgorithms != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SelectedEPSNASSecurityAlgorithms.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/SelectedEPSNASSecurityAlgorithms): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SelectedEPSNASSecurityAlgorithms.Octet); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/SelectedEPSNASSecurityAlgorithms): %w", err)
		}
	}
	if a.Additional5GSecurityInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Additional5GSecurityInformation.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/Additional5GSecurityInformation): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Additional5GSecurityInformation.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/Additional5GSecurityInformation): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Additional5GSecurityInformation.Octet); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/Additional5GSecurityInformation): %w", err)
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/EAPMessage): %w", err)
		}
	}
	if a.ABBA != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/ABBA): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/ABBA): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/ABBA): %w", err)
		}
	}
	if a.ReplayedS1UESecurityCapabilities != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/ReplayedS1UESecurityCapabilities): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/ReplayedS1UESecurityCapabilities): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (SecurityModeCommand/ReplayedS1UESecurityCapabilities): %w", err)
		}
	}
	return nil
}

func (a *SecurityModeCommand) DecodeSecurityModeCommand(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeCommand/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeCommand/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SecurityModeCommandMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeCommand/SecurityModeCommandMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SelectedNASSecurityAlgorithms.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeCommand/SelectedNASSecurityAlgorithms): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeCommand/SpareHalfOctetAndNgksi): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ReplayedUESecurityCapabilities.Len); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeCommand/ReplayedUESecurityCapabilities): %w", err)
	}
	if a.ReplayedUESecurityCapabilities.Len < 2 || a.ReplayedUESecurityCapabilities.Len > 8 {
		return fmt.Errorf("invalid ie length (SecurityModeCommand/ReplayedUESecurityCapabilities): %d", a.ReplayedUESecurityCapabilities.Len)
	}
	a.ReplayedUESecurityCapabilities.SetLen(a.ReplayedUESecurityCapabilities.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.ReplayedUESecurityCapabilities.Buffer); err != nil {
		return fmt.Errorf("NAS decode error (SecurityModeCommand/ReplayedUESecurityCapabilities): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (SecurityModeCommand/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case SecurityModeCommandIMEISVRequestType:
			a.IMEISVRequest = nasType.NewIMEISVRequest(ieiN)
			a.IMEISVRequest.Octet = ieiN
		case SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType:
			a.SelectedEPSNASSecurityAlgorithms = nasType.NewSelectedEPSNASSecurityAlgorithms(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SelectedEPSNASSecurityAlgorithms.Octet); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/SelectedEPSNASSecurityAlgorithms): %w", err)
			}
		case SecurityModeCommandAdditional5GSecurityInformationType:
			a.Additional5GSecurityInformation = nasType.NewAdditional5GSecurityInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Additional5GSecurityInformation.Len); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/Additional5GSecurityInformation): %w", err)
			}
			if a.Additional5GSecurityInformation.Len != 1 {
				return fmt.Errorf("invalid ie length (SecurityModeCommand/Additional5GSecurityInformation): %d", a.Additional5GSecurityInformation.Len)
			}
			a.Additional5GSecurityInformation.SetLen(a.Additional5GSecurityInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.Additional5GSecurityInformation.Octet); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/Additional5GSecurityInformation): %w", err)
			}
		case SecurityModeCommandEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/EAPMessage): %w", err)
			}
			if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
				return fmt.Errorf("invalid ie length (SecurityModeCommand/EAPMessage): %d", a.EAPMessage.Len)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/EAPMessage): %w", err)
			}
		case SecurityModeCommandABBAType:
			a.ABBA = nasType.NewABBA(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ABBA.Len); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/ABBA): %w", err)
			}
			if a.ABBA.Len < 2 {
				return fmt.Errorf("invalid ie length (SecurityModeCommand/ABBA): %d", a.ABBA.Len)
			}
			a.ABBA.SetLen(a.ABBA.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ABBA.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/ABBA): %w", err)
			}
		case SecurityModeCommandReplayedS1UESecurityCapabilitiesType:
			a.ReplayedS1UESecurityCapabilities = nasType.NewReplayedS1UESecurityCapabilities(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ReplayedS1UESecurityCapabilities.Len); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/ReplayedS1UESecurityCapabilities): %w", err)
			}
			if a.ReplayedS1UESecurityCapabilities.Len < 2 || a.ReplayedS1UESecurityCapabilities.Len > 5 {
				return fmt.Errorf("invalid ie length (SecurityModeCommand/ReplayedS1UESecurityCapabilities): %d", a.ReplayedS1UESecurityCapabilities.Len)
			}
			a.ReplayedS1UESecurityCapabilities.SetLen(a.ReplayedS1UESecurityCapabilities.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (SecurityModeCommand/ReplayedS1UESecurityCapabilities): %w", err)
			}
		default:
		}
	}
	return nil
}
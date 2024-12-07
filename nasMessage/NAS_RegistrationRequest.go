// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/machi12/nas/nasType"
)

type RegistrationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationRequestMessageIdentity
	nasType.NgksiAndRegistrationType5GS
	nasType.MobileIdentity5GS
	*nasType.NoncurrentNativeNASKeySetIdentifier
	*nasType.Capability5GMM
	*nasType.UESecurityCapability
	*nasType.RequestedNSSAI
	*nasType.LastVisitedRegisteredTAI
	*nasType.S1UENetworkCapability
	*nasType.UplinkDataStatus
	*nasType.PDUSessionStatus
	*nasType.MICOIndication
	*nasType.UEStatus
	*nasType.AdditionalGUTI
	*nasType.AllowedPDUSessionStatus
	*nasType.UesUsageSetting
	*nasType.RequestedDRXParameters
	*nasType.EPSNASMessageContainer
	*nasType.LADNIndication
	*nasType.PayloadContainer
	*nasType.NetworkSlicingIndication
	*nasType.UpdateType5GS
	*nasType.NASMessageContainer
	*nasType.EPSBearerContextStatus
	// NOTE: Add random number N
	*nasType.N
}

func NewRegistrationRequest(iei uint8) (registrationRequest *RegistrationRequest) {
	registrationRequest = &RegistrationRequest{}
	return registrationRequest
}

const (
	RegistrationRequestNoncurrentNativeNASKeySetIdentifierType uint8 = 0x0C
	RegistrationRequestCapability5GMMType                      uint8 = 0x10
	RegistrationRequestUESecurityCapabilityType                uint8 = 0x2E
	RegistrationRequestRequestedNSSAIType                      uint8 = 0x2F
	RegistrationRequestLastVisitedRegisteredTAIType            uint8 = 0x52
	RegistrationRequestS1UENetworkCapabilityType               uint8 = 0x17
	RegistrationRequestUplinkDataStatusType                    uint8 = 0x40
	RegistrationRequestPDUSessionStatusType                    uint8 = 0x50
	RegistrationRequestMICOIndicationType                      uint8 = 0x0B
	RegistrationRequestUEStatusType                            uint8 = 0x2B
	RegistrationRequestAdditionalGUTIType                      uint8 = 0x77
	RegistrationRequestAllowedPDUSessionStatusType             uint8 = 0x25
	RegistrationRequestUesUsageSettingType                     uint8 = 0x18
	RegistrationRequestRequestedDRXParametersType              uint8 = 0x51
	RegistrationRequestEPSNASMessageContainerType              uint8 = 0x70
	RegistrationRequestLADNIndicationType                      uint8 = 0x74
	RegistrationRequestPayloadContainerType                    uint8 = 0x7B
	RegistrationRequestNetworkSlicingIndicationType            uint8 = 0x09
	RegistrationRequestUpdateType5GSType                       uint8 = 0x53
	RegistrationRequestNASMessageContainerType                 uint8 = 0x71
	RegistrationRequestEPSBearerContextStatusType              uint8 = 0x60
	// NOTE: Add IEI for N
	RegistrationRequestNType								   uint8 = 0x29
)

func (a *RegistrationRequest) EncodeRegistrationRequest(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationRequest/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationRequest/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.RegistrationRequestMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationRequest/RegistrationRequestMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.NgksiAndRegistrationType5GS.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationRequest/NgksiAndRegistrationType5GS): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.MobileIdentity5GS.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationRequest/MobileIdentity5GS): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.MobileIdentity5GS.Buffer); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationRequest/MobileIdentity5GS): %w", err)
	}
	if a.NoncurrentNativeNASKeySetIdentifier != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NoncurrentNativeNASKeySetIdentifier.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/NoncurrentNativeNASKeySetIdentifier): %w", err)
		}
	}
	if a.Capability5GMM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GMM.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/Capability5GMM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GMM.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/Capability5GMM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GMM.Octet[:a.Capability5GMM.GetLen()]); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/Capability5GMM): %w", err)
		}
	}
	if a.UESecurityCapability != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UESecurityCapability.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UESecurityCapability): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UESecurityCapability.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UESecurityCapability): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UESecurityCapability.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UESecurityCapability): %w", err)
		}
	}
	if a.RequestedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/RequestedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/RequestedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedNSSAI.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/RequestedNSSAI): %w", err)
		}
	}
	if a.LastVisitedRegisteredTAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LastVisitedRegisteredTAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/LastVisitedRegisteredTAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LastVisitedRegisteredTAI.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/LastVisitedRegisteredTAI): %w", err)
		}
	}
	if a.S1UENetworkCapability != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.S1UENetworkCapability.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/S1UENetworkCapability): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.S1UENetworkCapability.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/S1UENetworkCapability): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.S1UENetworkCapability.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/S1UENetworkCapability): %w", err)
		}
	}
	if a.UplinkDataStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UplinkDataStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UplinkDataStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UplinkDataStatus): %w", err)
		}
	}
	if a.PDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/PDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/PDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/PDUSessionStatus): %w", err)
		}
	}
	if a.MICOIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MICOIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/MICOIndication): %w", err)
		}
	}
	if a.UEStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UEStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UEStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UEStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UEStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UEStatus.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UEStatus): %w", err)
		}
	}
	if a.AdditionalGUTI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/AdditionalGUTI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/AdditionalGUTI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/AdditionalGUTI): %w", err)
		}
	}
	if a.AllowedPDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/AllowedPDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/AllowedPDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/AllowedPDUSessionStatus): %w", err)
		}
	}
	if a.UesUsageSetting != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UesUsageSetting.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UesUsageSetting): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UesUsageSetting.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UesUsageSetting): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UesUsageSetting.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UesUsageSetting): %w", err)
		}
	}
	if a.RequestedDRXParameters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedDRXParameters.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/RequestedDRXParameters): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedDRXParameters.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/RequestedDRXParameters): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedDRXParameters.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/RequestedDRXParameters): %w", err)
		}
	}
	if a.EPSNASMessageContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EPSNASMessageContainer.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/EPSNASMessageContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EPSNASMessageContainer.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/EPSNASMessageContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EPSNASMessageContainer.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/EPSNASMessageContainer): %w", err)
		}
	}
	if a.LADNIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LADNIndication.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/LADNIndication): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LADNIndication.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/LADNIndication): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LADNIndication.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/LADNIndication): %w", err)
		}
	}
	if a.PayloadContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/PayloadContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/PayloadContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PayloadContainer.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/PayloadContainer): %w", err)
		}
	}
	if a.NetworkSlicingIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkSlicingIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/NetworkSlicingIndication): %w", err)
		}
	}
	if a.UpdateType5GS != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UpdateType5GS.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UpdateType5GS): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UpdateType5GS.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UpdateType5GS): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UpdateType5GS.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/UpdateType5GS): %w", err)
		}
	}
	if a.NASMessageContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/NASMessageContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/NASMessageContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/NASMessageContainer): %w", err)
		}
	}
	if a.EPSBearerContextStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/EPSBearerContextStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/EPSBearerContextStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationRequest/EPSBearerContextStatus): %w", err)
		}
	}
	return nil
}

func (a *RegistrationRequest) DecodeRegistrationRequest(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationRequest/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationRequest/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationRequestMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationRequest/RegistrationRequestMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.NgksiAndRegistrationType5GS.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationRequest/NgksiAndRegistrationType5GS): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.MobileIdentity5GS.Len); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationRequest/MobileIdentity5GS): %w", err)
	}
	if a.MobileIdentity5GS.Len < 4 {
		return fmt.Errorf("invalid ie length (RegistrationRequest/MobileIdentity5GS): %d", a.MobileIdentity5GS.Len)
	}
	a.MobileIdentity5GS.SetLen(a.MobileIdentity5GS.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.MobileIdentity5GS.Buffer); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationRequest/MobileIdentity5GS): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (RegistrationRequest/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case RegistrationRequestNoncurrentNativeNASKeySetIdentifierType:
			a.NoncurrentNativeNASKeySetIdentifier = nasType.NewNoncurrentNativeNASKeySetIdentifier(ieiN)
			a.NoncurrentNativeNASKeySetIdentifier.Octet = ieiN
		case RegistrationRequestCapability5GMMType:
			a.Capability5GMM = nasType.NewCapability5GMM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Capability5GMM.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/Capability5GMM): %w", err)
			}
			if a.Capability5GMM.Len < 1 || a.Capability5GMM.Len > 13 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/Capability5GMM): %d", a.Capability5GMM.Len)
			}
			a.Capability5GMM.SetLen(a.Capability5GMM.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.Capability5GMM.Octet[:a.Capability5GMM.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/Capability5GMM): %w", err)
			}
		case RegistrationRequestUESecurityCapabilityType:
			a.UESecurityCapability = nasType.NewUESecurityCapability(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UESecurityCapability.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UESecurityCapability): %w", err)
			}
			if a.UESecurityCapability.Len < 2 || a.UESecurityCapability.Len > 8 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/UESecurityCapability): %d", a.UESecurityCapability.Len)
			}
			a.UESecurityCapability.SetLen(a.UESecurityCapability.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.UESecurityCapability.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UESecurityCapability): %w", err)
			}
		case RegistrationRequestRequestedNSSAIType:
			a.RequestedNSSAI = nasType.NewRequestedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/RequestedNSSAI): %w", err)
			}
			if a.RequestedNSSAI.Len < 2 || a.RequestedNSSAI.Len > 72 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/RequestedNSSAI): %d", a.RequestedNSSAI.Len)
			}
			a.RequestedNSSAI.SetLen(a.RequestedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RequestedNSSAI.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/RequestedNSSAI): %w", err)
			}
		case RegistrationRequestLastVisitedRegisteredTAIType:
			a.LastVisitedRegisteredTAI = nasType.NewLastVisitedRegisteredTAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, a.LastVisitedRegisteredTAI.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/LastVisitedRegisteredTAI): %w", err)
			}
		case RegistrationRequestS1UENetworkCapabilityType:
			a.S1UENetworkCapability = nasType.NewS1UENetworkCapability(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.S1UENetworkCapability.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/S1UENetworkCapability): %w", err)
			}
			if a.S1UENetworkCapability.Len < 2 || a.S1UENetworkCapability.Len > 13 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/S1UENetworkCapability): %d", a.S1UENetworkCapability.Len)
			}
			a.S1UENetworkCapability.SetLen(a.S1UENetworkCapability.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.S1UENetworkCapability.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/S1UENetworkCapability): %w", err)
			}
			if len(a.S1UENetworkCapability.Buffer)!=int(a.S1UENetworkCapability.Len){
				return fmt.Errorf("invalid payload length(RegistrationRequest/S1UENetworkCapability), expected length: %d, but get %d",a.S1UENetworkCapability.Len,len(a.S1UENetworkCapability.Buffer))
			}
		case RegistrationRequestUplinkDataStatusType:
			a.UplinkDataStatus = nasType.NewUplinkDataStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UplinkDataStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UplinkDataStatus): %w", err)
			}
			if a.UplinkDataStatus.Len < 2 || a.UplinkDataStatus.Len > 32 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/UplinkDataStatus): %d", a.UplinkDataStatus.Len)
			}
			a.UplinkDataStatus.SetLen(a.UplinkDataStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.UplinkDataStatus.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UplinkDataStatus): %w", err)
			}
		case RegistrationRequestPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/PDUSessionStatus): %w", err)
			}
			if a.PDUSessionStatus.Len < 2 || a.PDUSessionStatus.Len > 32 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/PDUSessionStatus): %d", a.PDUSessionStatus.Len)
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/PDUSessionStatus): %w", err)
			}
		case RegistrationRequestMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case RegistrationRequestUEStatusType:
			a.UEStatus = nasType.NewUEStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UEStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UEStatus): %w", err)
			}
			if a.UEStatus.Len != 1 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/UEStatus): %d", a.UEStatus.Len)
			}
			a.UEStatus.SetLen(a.UEStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.UEStatus.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UEStatus): %w", err)
			}
		case RegistrationRequestAdditionalGUTIType:
			a.AdditionalGUTI = nasType.NewAdditionalGUTI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AdditionalGUTI.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/AdditionalGUTI): %w", err)
			}
			if a.AdditionalGUTI.Len != 11 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/AdditionalGUTI): %d", a.AdditionalGUTI.Len)
			}
			a.AdditionalGUTI.SetLen(a.AdditionalGUTI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AdditionalGUTI.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/AdditionalGUTI): %w", err)
			}
		case RegistrationRequestAllowedPDUSessionStatusType:
			a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/AllowedPDUSessionStatus): %w", err)
			}
			if a.AllowedPDUSessionStatus.Len < 2 || a.AllowedPDUSessionStatus.Len > 32 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/AllowedPDUSessionStatus): %d", a.AllowedPDUSessionStatus.Len)
			}
			a.AllowedPDUSessionStatus.SetLen(a.AllowedPDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/AllowedPDUSessionStatus): %w", err)
			}
		case RegistrationRequestUesUsageSettingType:
			a.UesUsageSetting = nasType.NewUesUsageSetting(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UesUsageSetting.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UesUsageSetting): %w", err)
			}
			if a.UesUsageSetting.Len != 1 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/UesUsageSetting): %d", a.UesUsageSetting.Len)
			}
			a.UesUsageSetting.SetLen(a.UesUsageSetting.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.UesUsageSetting.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UesUsageSetting): %w", err)
			}
		case RegistrationRequestRequestedDRXParametersType:
			a.RequestedDRXParameters = nasType.NewRequestedDRXParameters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedDRXParameters.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/RequestedDRXParameters): %w", err)
			}
			if a.RequestedDRXParameters.Len != 1 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/RequestedDRXParameters): %d", a.RequestedDRXParameters.Len)
			}
			a.RequestedDRXParameters.SetLen(a.RequestedDRXParameters.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedDRXParameters.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/RequestedDRXParameters): %w", err)
			}
		case RegistrationRequestEPSNASMessageContainerType:
			a.EPSNASMessageContainer = nasType.NewEPSNASMessageContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EPSNASMessageContainer.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/EPSNASMessageContainer): %w", err)
			}
			if a.EPSNASMessageContainer.Len < 1 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/EPSNASMessageContainer): %d", a.EPSNASMessageContainer.Len)
			}
			a.EPSNASMessageContainer.SetLen(a.EPSNASMessageContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EPSNASMessageContainer.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/EPSNASMessageContainer): %w", err)
			}
		case RegistrationRequestLADNIndicationType:
			a.LADNIndication = nasType.NewLADNIndication(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LADNIndication.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/LADNIndication): %w", err)
			}
			if a.LADNIndication.Len > 808 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/LADNIndication): %d", a.LADNIndication.Len)
			}
			a.LADNIndication.SetLen(a.LADNIndication.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.LADNIndication.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/LADNIndication): %w", err)
			}
		case RegistrationRequestPayloadContainerType:
			a.PayloadContainer = nasType.NewPayloadContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/PayloadContainer): %w", err)
			}
			if a.PayloadContainer.Len < 1 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/PayloadContainer): %d", a.PayloadContainer.Len)
			}
			a.PayloadContainer.SetLen(a.PayloadContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PayloadContainer.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/PayloadContainer): %w", err)
			}
		case RegistrationRequestNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case RegistrationRequestUpdateType5GSType:
			a.UpdateType5GS = nasType.NewUpdateType5GS(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UpdateType5GS.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UpdateType5GS): %w", err)
			}
			if a.UpdateType5GS.Len != 1 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/UpdateType5GS): %d", a.UpdateType5GS.Len)
			}
			a.UpdateType5GS.SetLen(a.UpdateType5GS.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.UpdateType5GS.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/UpdateType5GS): %w", err)
			}
		case RegistrationRequestNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/NASMessageContainer): %w", err)
			}
			if a.NASMessageContainer.Len < 1 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/NASMessageContainer): %d", a.NASMessageContainer.Len)
			}
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/NASMessageContainer): %w", err)
			}
		case RegistrationRequestEPSBearerContextStatusType:
			a.EPSBearerContextStatus = nasType.NewEPSBearerContextStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EPSBearerContextStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/EPSBearerContextStatus): %w", err)
			}
			if a.EPSBearerContextStatus.Len != 2 {
				return fmt.Errorf("invalid ie length (RegistrationRequest/EPSBearerContextStatus): %d", a.EPSBearerContextStatus.Len)
			}
			a.EPSBearerContextStatus.SetLen(a.EPSBearerContextStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EPSBearerContextStatus.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/EPSBearerContextStatus): %w", err)
			}
		case RegistrationRequestNType:
			// NOTE: Add decoding for N sent by UE
			a.N = nasType.NewN(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, a.N.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationRequest/N): %w", err)
			}
		default:
		}
	}
	return nil
}
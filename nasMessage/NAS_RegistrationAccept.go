// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/machi12/nas/nasType"
)

type RegistrationAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationAcceptMessageIdentity
	nasType.RegistrationResult5GS
	*nasType.GUTI5G
	*nasType.EquivalentPlmns
	*nasType.TAIList
	*nasType.AllowedNSSAI
	*nasType.RejectedNSSAI
	*nasType.ConfiguredNSSAI
	*nasType.NetworkFeatureSupport5GS
	*nasType.PDUSessionStatus
	*nasType.PDUSessionReactivationResult
	*nasType.PDUSessionReactivationResultErrorCause
	*nasType.LADNInformation
	*nasType.MICOIndication
	*nasType.NetworkSlicingIndication
	*nasType.ServiceAreaList
	*nasType.T3512Value
	*nasType.Non3GppDeregistrationTimerValue
	*nasType.T3502Value
	*nasType.EmergencyNumberList
	*nasType.ExtendedEmergencyNumberList
	*nasType.SORTransparentContainer
	*nasType.EAPMessage
	*nasType.NSSAIInclusionMode
	*nasType.OperatordefinedAccessCategoryDefinitions
	*nasType.NegotiatedDRXParameters
	*nasType.Non3GppNwPolicies
	*nasType.EPSBearerContextStatus
}

func NewRegistrationAccept(iei uint8) (registrationAccept *RegistrationAccept) {
	registrationAccept = &RegistrationAccept{}
	return registrationAccept
}

const (
	RegistrationAcceptGUTI5GType                                   uint8 = 0x77
	RegistrationAcceptEquivalentPlmnsType                          uint8 = 0x4A
	RegistrationAcceptTAIListType                                  uint8 = 0x54
	RegistrationAcceptAllowedNSSAIType                             uint8 = 0x15
	RegistrationAcceptRejectedNSSAIType                            uint8 = 0x11
	RegistrationAcceptConfiguredNSSAIType                          uint8 = 0x31
	RegistrationAcceptNetworkFeatureSupport5GSType                 uint8 = 0x21
	RegistrationAcceptPDUSessionStatusType                         uint8 = 0x50
	RegistrationAcceptPDUSessionReactivationResultType             uint8 = 0x26
	RegistrationAcceptPDUSessionReactivationResultErrorCauseType   uint8 = 0x72
	RegistrationAcceptLADNInformationType                          uint8 = 0x79
	RegistrationAcceptMICOIndicationType                           uint8 = 0x0B
	RegistrationAcceptNetworkSlicingIndicationType                 uint8 = 0x09
	RegistrationAcceptServiceAreaListType                          uint8 = 0x27
	RegistrationAcceptT3512ValueType                               uint8 = 0x5E
	RegistrationAcceptNon3GppDeregistrationTimerValueType          uint8 = 0x5D
	RegistrationAcceptT3502ValueType                               uint8 = 0x16
	RegistrationAcceptEmergencyNumberListType                      uint8 = 0x34
	RegistrationAcceptExtendedEmergencyNumberListType              uint8 = 0x7A
	RegistrationAcceptSORTransparentContainerType                  uint8 = 0x73
	RegistrationAcceptEAPMessageType                               uint8 = 0x78
	RegistrationAcceptNSSAIInclusionModeType                       uint8 = 0x0A
	RegistrationAcceptOperatordefinedAccessCategoryDefinitionsType uint8 = 0x76
	RegistrationAcceptNegotiatedDRXParametersType                  uint8 = 0x51
	RegistrationAcceptNon3GppNwPoliciesType                        uint8 = 0x0D
	RegistrationAcceptEPSBearerContextStatusType                   uint8 = 0x60
)

func (a *RegistrationAccept) EncodeRegistrationAccept(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationAccept/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationAccept/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.RegistrationAcceptMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationAccept/RegistrationAcceptMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.RegistrationResult5GS.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationAccept/RegistrationResult5GS): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.RegistrationResult5GS.Octet); err != nil {
		return fmt.Errorf("NAS encode error (RegistrationAccept/RegistrationResult5GS): %w", err)
	}
	if a.GUTI5G != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/GUTI5G): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/GUTI5G): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/GUTI5G): %w", err)
		}
	}
	if a.EquivalentPlmns != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EquivalentPlmns): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EquivalentPlmns): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.Octet[:a.EquivalentPlmns.GetLen()]); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EquivalentPlmns): %w", err)
		}
	}
	if a.TAIList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.TAIList.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/TAIList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.TAIList.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/TAIList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.TAIList.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/TAIList): %w", err)
		}
	}
	if a.AllowedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/AllowedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/AllowedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/AllowedNSSAI): %w", err)
		}
	}
	if a.RejectedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/RejectedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/RejectedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/RejectedNSSAI): %w", err)
		}
	}
	if a.ConfiguredNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ConfiguredNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ConfiguredNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ConfiguredNSSAI): %w", err)
		}
	}
	if a.NetworkFeatureSupport5GS != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/NetworkFeatureSupport5GS): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/NetworkFeatureSupport5GS): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.Octet[:a.NetworkFeatureSupport5GS.GetLen()]); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/NetworkFeatureSupport5GS): %w", err)
		}
	}
	if a.PDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionStatus): %w", err)
		}
	}
	if a.PDUSessionReactivationResult != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionReactivationResult): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionReactivationResult): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionReactivationResult): %w", err)
		}
	}
	if a.PDUSessionReactivationResultErrorCause != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionReactivationResultErrorCause): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionReactivationResultErrorCause): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/PDUSessionReactivationResultErrorCause): %w", err)
		}
	}
	if a.LADNInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LADNInformation.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/LADNInformation): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LADNInformation.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/LADNInformation): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LADNInformation.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/LADNInformation): %w", err)
		}
	}
	if a.MICOIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MICOIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/MICOIndication): %w", err)
		}
	}
	if a.NetworkSlicingIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkSlicingIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/NetworkSlicingIndication): %w", err)
		}
	}
	if a.ServiceAreaList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ServiceAreaList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ServiceAreaList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ServiceAreaList): %w", err)
		}
	}
	if a.T3512Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3512Value.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/T3512Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3512Value.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/T3512Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3512Value.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/T3512Value): %w", err)
		}
	}
	if a.Non3GppDeregistrationTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GppDeregistrationTimerValue.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/Non3GppDeregistrationTimerValue): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GppDeregistrationTimerValue.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/Non3GppDeregistrationTimerValue): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GppDeregistrationTimerValue.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/Non3GppDeregistrationTimerValue): %w", err)
		}
	}
	if a.T3502Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3502Value.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/T3502Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3502Value.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/T3502Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3502Value.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/T3502Value): %w", err)
		}
	}
	if a.EmergencyNumberList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EmergencyNumberList.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EmergencyNumberList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EmergencyNumberList.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EmergencyNumberList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EmergencyNumberList.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EmergencyNumberList): %w", err)
		}
	}
	if a.ExtendedEmergencyNumberList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ExtendedEmergencyNumberList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ExtendedEmergencyNumberList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/ExtendedEmergencyNumberList): %w", err)
		}
	}
	if a.SORTransparentContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/SORTransparentContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/SORTransparentContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/SORTransparentContainer): %w", err)
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EAPMessage): %w", err)
		}
	}
	if a.NSSAIInclusionMode != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NSSAIInclusionMode.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/NSSAIInclusionMode): %w", err)
		}
	}
	if a.OperatordefinedAccessCategoryDefinitions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/OperatordefinedAccessCategoryDefinitions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/OperatordefinedAccessCategoryDefinitions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/OperatordefinedAccessCategoryDefinitions): %w", err)
		}
	}
	if a.NegotiatedDRXParameters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedDRXParameters.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/NegotiatedDRXParameters): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedDRXParameters.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/NegotiatedDRXParameters): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedDRXParameters.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/NegotiatedDRXParameters): %w", err)
		}
	}
	if a.Non3GppNwPolicies != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GppNwPolicies.Octet); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/Non3GppNwPolicies): %w", err)
		}
	}
	if a.EPSBearerContextStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EPSBearerContextStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EPSBearerContextStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (RegistrationAccept/EPSBearerContextStatus): %w", err)
		}
	}
	return nil
}

func (a *RegistrationAccept) DecodeRegistrationAccept(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationAccept/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationAccept/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationAcceptMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationAccept/RegistrationAcceptMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Len); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationAccept/RegistrationResult5GS): %w", err)
	}
	if a.RegistrationResult5GS.Len != 1 {
		return fmt.Errorf("invalid ie length (RegistrationAccept/RegistrationResult5GS): %d", a.RegistrationResult5GS.Len)
	}
	a.RegistrationResult5GS.SetLen(a.RegistrationResult5GS.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet); err != nil {
		return fmt.Errorf("NAS decode error (RegistrationAccept/RegistrationResult5GS): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (RegistrationAccept/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case RegistrationAcceptGUTI5GType:
			a.GUTI5G = nasType.NewGUTI5G(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.GUTI5G.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/GUTI5G): %w", err)
			}
			if a.GUTI5G.Len != 11 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/GUTI5G): %d", a.GUTI5G.Len)
			}
			a.GUTI5G.SetLen(a.GUTI5G.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.GUTI5G.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/GUTI5G): %w", err)
			}
		case RegistrationAcceptEquivalentPlmnsType:
			a.EquivalentPlmns = nasType.NewEquivalentPlmns(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EquivalentPlmns.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/EquivalentPlmns): %w", err)
			}
			if a.EquivalentPlmns.Len < 3 || a.EquivalentPlmns.Len > 45 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/EquivalentPlmns): %d", a.EquivalentPlmns.Len)
			}
			a.EquivalentPlmns.SetLen(a.EquivalentPlmns.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EquivalentPlmns.Octet[:a.EquivalentPlmns.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/EquivalentPlmns): %w", err)
			}
		case RegistrationAcceptTAIListType:
			a.TAIList = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.TAIList.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/TAIList): %w", err)
			}
			if a.TAIList.Len < 7 || a.TAIList.Len > 112 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/TAIList): %d", a.TAIList.Len)
			}
			a.TAIList.SetLen(a.TAIList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.TAIList.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/TAIList): %w", err)
			}
		case RegistrationAcceptAllowedNSSAIType:
			a.AllowedNSSAI = nasType.NewAllowedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AllowedNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/AllowedNSSAI): %w", err)
			}
			if a.AllowedNSSAI.Len < 2 || a.AllowedNSSAI.Len > 72 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/AllowedNSSAI): %d", a.AllowedNSSAI.Len)
			}
			a.AllowedNSSAI.SetLen(a.AllowedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AllowedNSSAI.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/AllowedNSSAI): %w", err)
			}
		case RegistrationAcceptRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/RejectedNSSAI): %w", err)
			}
			if a.RejectedNSSAI.Len < 2 || a.RejectedNSSAI.Len > 40 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/RejectedNSSAI): %d", a.RejectedNSSAI.Len)
			}
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/RejectedNSSAI): %w", err)
			}
		case RegistrationAcceptConfiguredNSSAIType:
			a.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/ConfiguredNSSAI): %w", err)
			}
			if a.ConfiguredNSSAI.Len < 2 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/ConfiguredNSSAI): %d", a.ConfiguredNSSAI.Len)
			}
			a.ConfiguredNSSAI.SetLen(a.ConfiguredNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ConfiguredNSSAI.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/ConfiguredNSSAI): %w", err)
			}
		case RegistrationAcceptNetworkFeatureSupport5GSType:
			a.NetworkFeatureSupport5GS = nasType.NewNetworkFeatureSupport5GS(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NetworkFeatureSupport5GS.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/NetworkFeatureSupport5GS): %w", err)
			}
			if a.NetworkFeatureSupport5GS.Len < 1 || a.NetworkFeatureSupport5GS.Len > 3 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/NetworkFeatureSupport5GS): %d", a.NetworkFeatureSupport5GS.Len)
			}
			a.NetworkFeatureSupport5GS.SetLen(a.NetworkFeatureSupport5GS.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.Octet[:a.NetworkFeatureSupport5GS.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/NetworkFeatureSupport5GS): %w", err)
			}
		case RegistrationAcceptPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/PDUSessionStatus): %w", err)
			}
			if a.PDUSessionStatus.Len < 2 || a.PDUSessionStatus.Len > 32 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/PDUSessionStatus): %d", a.PDUSessionStatus.Len)
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/PDUSessionStatus): %w", err)
			}
		case RegistrationAcceptPDUSessionReactivationResultType:
			a.PDUSessionReactivationResult = nasType.NewPDUSessionReactivationResult(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResult.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/PDUSessionReactivationResult): %w", err)
			}
			if a.PDUSessionReactivationResult.Len < 2 || a.PDUSessionReactivationResult.Len > 32 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/PDUSessionReactivationResult): %d", a.PDUSessionReactivationResult.Len)
			}
			a.PDUSessionReactivationResult.SetLen(a.PDUSessionReactivationResult.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResult.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/PDUSessionReactivationResult): %w", err)
			}
		case RegistrationAcceptPDUSessionReactivationResultErrorCauseType:
			a.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResultErrorCause.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/PDUSessionReactivationResultErrorCause): %w", err)
			}
			if a.PDUSessionReactivationResultErrorCause.Len < 2 || a.PDUSessionReactivationResultErrorCause.Len > 512 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/PDUSessionReactivationResultErrorCause): %d", a.PDUSessionReactivationResultErrorCause.Len)
			}
			a.PDUSessionReactivationResultErrorCause.SetLen(a.PDUSessionReactivationResultErrorCause.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/PDUSessionReactivationResultErrorCause): %w", err)
			}
		case RegistrationAcceptLADNInformationType:
			a.LADNInformation = nasType.NewLADNInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LADNInformation.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/LADNInformation): %w", err)
			}
			if a.LADNInformation.Len < 9 || a.LADNInformation.Len > 1712 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/LADNInformation): %d", a.LADNInformation.Len)
			}
			a.LADNInformation.SetLen(a.LADNInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.LADNInformation.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/LADNInformation): %w", err)
			}
		case RegistrationAcceptMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case RegistrationAcceptNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case RegistrationAcceptServiceAreaListType:
			a.ServiceAreaList = nasType.NewServiceAreaList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceAreaList.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/ServiceAreaList): %w", err)
			}
			if a.ServiceAreaList.Len < 4 || a.ServiceAreaList.Len > 112 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/ServiceAreaList): %d", a.ServiceAreaList.Len)
			}
			a.ServiceAreaList.SetLen(a.ServiceAreaList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceAreaList.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/ServiceAreaList): %w", err)
			}
		case RegistrationAcceptT3512ValueType:
			a.T3512Value = nasType.NewT3512Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3512Value.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/T3512Value): %w", err)
			}
			if a.T3512Value.Len != 1 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/T3512Value): %d", a.T3512Value.Len)
			}
			a.T3512Value.SetLen(a.T3512Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3512Value.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/T3512Value): %w", err)
			}
		case RegistrationAcceptNon3GppDeregistrationTimerValueType:
			a.Non3GppDeregistrationTimerValue = nasType.NewNon3GppDeregistrationTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Non3GppDeregistrationTimerValue.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/Non3GppDeregistrationTimerValue): %w", err)
			}
			if a.Non3GppDeregistrationTimerValue.Len != 1 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/Non3GppDeregistrationTimerValue): %d", a.Non3GppDeregistrationTimerValue.Len)
			}
			a.Non3GppDeregistrationTimerValue.SetLen(a.Non3GppDeregistrationTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.Non3GppDeregistrationTimerValue.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/Non3GppDeregistrationTimerValue): %w", err)
			}
		case RegistrationAcceptT3502ValueType:
			a.T3502Value = nasType.NewT3502Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3502Value.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/T3502Value): %w", err)
			}
			if a.T3502Value.Len != 1 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/T3502Value): %d", a.T3502Value.Len)
			}
			a.T3502Value.SetLen(a.T3502Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3502Value.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/T3502Value): %w", err)
			}
		case RegistrationAcceptEmergencyNumberListType:
			a.EmergencyNumberList = nasType.NewEmergencyNumberList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EmergencyNumberList.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/EmergencyNumberList): %w", err)
			}
			if a.EmergencyNumberList.Len < 3 || a.EmergencyNumberList.Len > 48 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/EmergencyNumberList): %d", a.EmergencyNumberList.Len)
			}
			a.EmergencyNumberList.SetLen(a.EmergencyNumberList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EmergencyNumberList.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/EmergencyNumberList): %w", err)
			}
		case RegistrationAcceptExtendedEmergencyNumberListType:
			a.ExtendedEmergencyNumberList = nasType.NewExtendedEmergencyNumberList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedEmergencyNumberList.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/ExtendedEmergencyNumberList): %w", err)
			}
			if a.ExtendedEmergencyNumberList.Len < 4 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/ExtendedEmergencyNumberList): %d", a.ExtendedEmergencyNumberList.Len)
			}
			a.ExtendedEmergencyNumberList.SetLen(a.ExtendedEmergencyNumberList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/ExtendedEmergencyNumberList): %w", err)
			}
		case RegistrationAcceptSORTransparentContainerType:
			a.SORTransparentContainer = nasType.NewSORTransparentContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SORTransparentContainer.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/SORTransparentContainer): %w", err)
			}
			if a.SORTransparentContainer.Len < 17 || a.SORTransparentContainer.Len > 2045 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/SORTransparentContainer): %d", a.SORTransparentContainer.Len)
			}
			a.SORTransparentContainer.SetLen(a.SORTransparentContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SORTransparentContainer.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/SORTransparentContainer): %w", err)
			}
		case RegistrationAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/EAPMessage): %w", err)
			}
			if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/EAPMessage): %d", a.EAPMessage.Len)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/EAPMessage): %w", err)
			}
		case RegistrationAcceptNSSAIInclusionModeType:
			a.NSSAIInclusionMode = nasType.NewNSSAIInclusionMode(ieiN)
			a.NSSAIInclusionMode.Octet = ieiN
		case RegistrationAcceptOperatordefinedAccessCategoryDefinitionsType:
			a.OperatordefinedAccessCategoryDefinitions = nasType.NewOperatordefinedAccessCategoryDefinitions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/OperatordefinedAccessCategoryDefinitions): %w", err)
			}
			a.OperatordefinedAccessCategoryDefinitions.SetLen(a.OperatordefinedAccessCategoryDefinitions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/OperatordefinedAccessCategoryDefinitions): %w", err)
			}
		case RegistrationAcceptNegotiatedDRXParametersType:
			a.NegotiatedDRXParameters = nasType.NewNegotiatedDRXParameters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NegotiatedDRXParameters.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/NegotiatedDRXParameters): %w", err)
			}
			if a.NegotiatedDRXParameters.Len != 1 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/NegotiatedDRXParameters): %d", a.NegotiatedDRXParameters.Len)
			}
			a.NegotiatedDRXParameters.SetLen(a.NegotiatedDRXParameters.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.NegotiatedDRXParameters.Octet); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/NegotiatedDRXParameters): %w", err)
			}
		case RegistrationAcceptNon3GppNwPoliciesType:
			a.Non3GppNwPolicies = nasType.NewNon3GppNwPolicies(ieiN)
			a.Non3GppNwPolicies.Octet = ieiN
		case RegistrationAcceptEPSBearerContextStatusType:
			a.EPSBearerContextStatus = nasType.NewEPSBearerContextStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EPSBearerContextStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/EPSBearerContextStatus): %w", err)
			}
			if a.EPSBearerContextStatus.Len != 2 {
				return fmt.Errorf("invalid ie length (RegistrationAccept/EPSBearerContextStatus): %d", a.EPSBearerContextStatus.Len)
			}
			a.EPSBearerContextStatus.SetLen(a.EPSBearerContextStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EPSBearerContextStatus.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (RegistrationAccept/EPSBearerContextStatus): %w", err)
			}
		default:
		}
	}
	return nil
}

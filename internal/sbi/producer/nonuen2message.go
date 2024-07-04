package producer

import (
	"encoding/json"
	"log"

	amf_context "github.com/free5gc/amf/internal/context"
	ngap_message "github.com/free5gc/amf/internal/ngap/message"
	"github.com/free5gc/openapi/models"
)

func NonUeN2MessageTransferProcedure(amfSelf *amf_context.AMFContext, message models.NonUeN2MessageTransferRequest) error {
    var keyValueN2Information map[string]string
    err := json.Unmarshal(message.BinaryDataN2Information, &keyValueN2Information)
    if err != nil {
        log.Printf("Failed to unmarshal BinaryDataN2Information: %v", err)
        return err
    }
    log.Printf("Successfully unmarshaled BinaryDataN2Information: %v", keyValueN2Information)

    amf_context.GetSelf().AmfRanPool.Range(func(key, value interface{}) bool {
        amfRan := value.(*amf_context.AmfRan)
        log.Printf("Sending WriteReplaceWarningRequest to RAN: %v", amfRan)
        ngap_message.SendWriteReplaceWarningRequest(amfRan, keyValueN2Information)
        return true
    })

    return nil
}


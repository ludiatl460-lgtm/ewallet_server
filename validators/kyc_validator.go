package validators

import (
	"strings"
	"time"

	"backend/models"
)

const dateLayout = "2006-01-02" // ຮູບແບບວັນທີ YYYY-MM-DD

var documentTypes = map[string]bool{"id_card": true, "family_book": true, "passport": true}
var detectionTypes = map[string]bool{"doc_img": true, "face_detection_img": true, "face_detection_vdo": true}

// ຄືນຂໍ້ຄວາມ error; ຖ້າຂໍ້ມູນຖືກຕ້ອງຄືນ "" (ຫວ່າງ)
func ValidateKyc(k *models.Kyc) string {
	if k.UserId <= 0 {
		return "user_id is required"
	}
	if !documentTypes[k.DocumentType] {
		return "document_type must be id_card, family_book or passport"
	}
	if strings.TrimSpace(k.DocumentNo) == "" || len(k.DocumentNo) > 15 {
		return "document_no is required (max 15 characters)"
	}
	if _, err := time.Parse(dateLayout, k.DocumentIssue); err != nil {
		return "document_issue must be YYYY-MM-DD"
	}
	expiry, err := time.Parse(dateLayout, k.DocumentExpiry)
	if err != nil {
		return "document_expiry must be YYYY-MM-DD"
	}
	if !expiry.After(time.Now()) {
		return "document is already expired"
	}
	return ""
}

func ValidateKycDetail(d *models.KycDetail) string {
	if strings.TrimSpace(d.ImgOrVdo) == "" {
		return "img_or_vdo is required"
	}
	if !detectionTypes[d.TypeOfDetection] {
		return "type_of_detection must be doc_img, face_detection_img or face_detection_vdo"
	}
	return ""
}
package services

import (
	"backend/config"
	"backend/models"
)

func CreateKyc(k *models.Kyc) (*models.Kyc, error) {
	var created models.Kyc
	// ໃສ່ສະເພາະຊ່ອງທີ່ຜູ້ໃຊ້ກຳນົດໄດ້; status ເປັນ 'pending' ຕາມຄ່າເລີ່ມຕົ້ນຂອງ DB
	err := config.DB.Get(&created,
		`INSERT INTO kyc (user_id, address, document_type, document_no, document_issue, document_expiry)
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		k.UserId, k.Address, k.DocumentType, k.DocumentNo, k.DocumentIssue, k.DocumentExpiry)
	return &created, err
}

func GetKycByID(id int64) (*models.Kyc, error) {
	var k models.Kyc
	err := config.DB.Get(&k, `SELECT * FROM kyc WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	k.Details, err = GetDetailsByKycID(id)
	return &k, err
}

// ອະນຸມັດໄດ້ສະເພາະໃບທີ່ຍັງ pending (ກັນອະນຸມັດຊ້ຳ)
// ອາຍຸ KYC 2 ປີ ເປັນຄ່າສົມມຸດ, ຕ້ອງຢືນຢັນກັບທີມ
func ApproveKyc(id int64, adminName string) (*models.Kyc, error) {
	var k models.Kyc
	err := config.DB.Get(&k,
		`UPDATE kyc
		 SET status='approved', is_active=true, rejected_reason=NULL,
		     issue_date=CURRENT_DATE, expiry_date=CURRENT_DATE + INTERVAL '2 years',
		     manage_by=$1, manage_at=NOW()
		 WHERE id=$2 AND status='pending'
		 RETURNING *`,
		adminName, id)
	return &k, err
}

func RejectKyc(id int64, adminName, reason string) (*models.Kyc, error) {
	var k models.Kyc
	err := config.DB.Get(&k,
		`UPDATE kyc
		 SET status='rejected', is_active=false, rejected_reason=$1,
		     manage_by=$2, manage_at=NOW()
		 WHERE id=$3 AND status='pending'
		 RETURNING *`,
		reason, adminName, id)
	return &k, err
}
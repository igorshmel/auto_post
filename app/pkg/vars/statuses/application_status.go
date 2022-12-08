package status

// ApplicationStatus _
type ApplicationStatus string

const (
	// ApplicationOperatorApproval - На согласовании у oперациониста
	ApplicationOperatorApproval ApplicationStatus = "operator_approval"
	// ApplicationLawyerApproval - На согласовании у Юриста
	ApplicationLawyerApproval ApplicationStatus = "lawyer_approval"
	// ApplicationDeclined - Отклонена
	ApplicationDeclined ApplicationStatus = "declined"
	// ApplicationUnderRevision - На доработке
	ApplicationUnderRevision ApplicationStatus = "under_revision"
	// ApplicationApproved - Согласовано
	ApplicationApproved ApplicationStatus = "approved"
)

func (ths ApplicationStatus) String() string {
	return string(ths)
}

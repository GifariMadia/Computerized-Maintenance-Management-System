package domain

// Four Engineering divisions defined in URS PT Bintang Toedjoe Site Pulogadung.
type DivisionCode string

const (
	DivisionMaintenance DivisionCode = "MAINTENANCE"
	DivisionUtility     DivisionCode = "UTILITY"
	DivisionBuilding    DivisionCode = "BUILDING"
	DivisionSparePart   DivisionCode = "SPARE_PART"
)

// Roles defined in URS ALCOA 1.1 Access Control Matrix.
type Role string

const (
	RoleAdministrator  Role = "ADMINISTRATOR"
	RoleTechnician     Role = "TECHNICIAN"
	RoleInspector      Role = "INSPECTOR"
	RoleSupervisor     Role = "SUPERVISOR"
	RoleManager        Role = "MANAGER"
	RoleSparePartStaff Role = "SPARE_PART_STAFF"
)

// Asset statuses (URS 3.1 - 3.6).
type AssetStatus string

const (
	AssetActive           AssetStatus = "ACTIVE"
	AssetBreakdown        AssetStatus = "BREAKDOWN"
	AssetUnderMaintenance AssetStatus = "UNDER_MAINTENANCE"
	AssetStandby          AssetStatus = "STANDBY"
	AssetDeactivated      AssetStatus = "DEACTIVATED"
)

// Work Order statuses kept distinct from asset condition and work result (URS 4.10b).
type WorkOrderStatus string

const (
	WOOpen                WorkOrderStatus = "OPEN"
	WOInProgress          WorkOrderStatus = "IN_PROGRESS"
	WOWaitingPart         WorkOrderStatus = "WAITING_PART"
	WOPendingVerification WorkOrderStatus = "PENDING_VERIFICATION"
	WORework              WorkOrderStatus = "REWORK"
	WOClosed              WorkOrderStatus = "CLOSED"
	WOCancelled           WorkOrderStatus = "CANCELLED"
)

// Priority levels (URS 4.1c).
type Priority string

const (
	PriorityEmergency Priority = "EMERGENCY"
	PriorityHigh      Priority = "HIGH"
	PriorityMedium    Priority = "MEDIUM"
	PriorityLow       Priority = "LOW"
)

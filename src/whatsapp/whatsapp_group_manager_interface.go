package whatsapp

// QpGroupSettings holds optional settings for group creation.
// All boolean pointer fields are optional; nil means use the WhatsApp default.
type QpGroupSettings struct {
	// AllMembersCanEditInfo controls whether all members or only admins can edit group info.
	// nil = WhatsApp default; false = only admins; true = all members.
	AllMembersCanEditInfo *bool

	// AllMembersCanSendMessages controls whether all members or only admins can send messages.
	// nil = WhatsApp default; true = all members; false = only admins.
	AllMembersCanSendMessages *bool

	// AllMembersCanAddMembers controls whether all members or only admins can add new members.
	// nil = WhatsApp default; true = all members; false = only admins.
	AllMembersCanAddMembers *bool

	// GenerateInviteLink if true, generates and returns the group invite link after creation.
	GenerateInviteLink bool
}

// WhatsappGroupManagerInterface defines the interface for group management operations
// This interface should be implemented by the group manager in the whatsmeow package
type WhatsappGroupManagerInterface interface {
	// Get group invite link
	GetInvite(groupId string) (string, error)

	// Get a list of all groups
	GetJoinedGroups() ([]interface{}, error)

	// Get a specific group
	GetGroupInfo(string) (interface{}, error)

	// Create a group
	CreateGroup(string, []string) (interface{}, error)

	// Create a group with extended options (title and participants)
	CreateGroupExtended(title string, participants []string) (interface{}, error)

	// Create group with extended options (map-based for QP level)
	CreateGroupExtendedWithOptions(options map[string]interface{}) (interface{}, error)

	// Create a group with additional permission settings.
	// Returns the group info and, when settings.GenerateInviteLink is true, the invite link.
	CreateGroupWithSettings(title string, participants []string, settings QpGroupSettings) (groupInfo interface{}, inviteLink string, err error)

	// Update Group Name
	UpdateGroupSubject(string, string) (interface{}, error)

	// Update Group Topic (Description)
	UpdateGroupTopic(string, string) (interface{}, error)

	// Update Group Photo
	UpdateGroupPhoto(string, []byte) (string, error)

	// Update group participants (add, remove, promote, demote)
	UpdateGroupParticipants(groupJID string, participants []string, action string) ([]interface{}, error)

	// Get list of pending join requests for a group
	GetGroupJoinRequests(groupJID string) ([]interface{}, error)

	// Handle join requests (approve/reject)
	HandleGroupJoinRequests(groupJID string, participants []string, action string) ([]interface{}, error)

	// Leave a group
	LeaveGroup(groupID string) error
}

// IWhatsappConnectionWithGroups extends IWhatsappConnection with group management
// Use this interface when you need both connection and group operations
type IWhatsappConnectionWithGroups interface {
	IWhatsappConnection

	// GetGroupManager returns the group manager for group operations
	GetGroupManager() WhatsappGroupManagerInterface
}

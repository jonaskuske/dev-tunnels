// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
// Generated from ../../../cs/src/Contracts/TunnelConstraints.cs

package tunnels

import (
	"regexp"
)

const (
	// Min length of tunnel cluster ID.
	TunnelConstraintsClusterIDMinLength = 3

	// Max length of tunnel cluster ID.
	TunnelConstraintsClusterIDMaxLength = 12

	// Length of tunnel id.
	TunnelConstraintsTunnelIDLength = 8

	// Min length of tunnel name.
	TunnelConstraintsTunnelNameMinLength = 3

	// Max length of tunnel name.
	TunnelConstraintsTunnelNameMaxLength = 60

	// Max length of tunnel or port description.
	TunnelConstraintsDescriptionMaxLength = 400

	// Min length of a single tunnel or port tag.
	TunnelConstraintsTagMinLength = 1

	// Max length of a single tunnel or port tag.
	TunnelConstraintsTagMaxLength = 50

	// Maximum number of tags that can be applied to a tunnel or port.
	TunnelConstraintsMaxTags = 100

	// Min length of a tunnel domain.
	TunnelConstraintsTunnelDomainMinLength = 4

	// Max length of a tunnel domain.
	TunnelConstraintsTunnelDomainMaxLength = 180

	// Maximum number of items allowed in the tunnel ports array. The actual limit on number
	// of ports that can be created may be much lower, and may depend on various resource
	// limitations or policies.
	TunnelConstraintsTunnelMaxPorts = 1000

	// Maximum number of access control entries (ACEs) in a tunnel or tunnel port access
	// control list (ACL).
	TunnelConstraintsAccessControlMaxEntries = 40

	// Maximum number of subjects (such as user IDs) in a tunnel or tunnel port access
	// control entry (ACE).
	TunnelConstraintsAccessControlMaxSubjects = 100

	// Max length of an access control subject or organization ID.
	TunnelConstraintsAccessControlSubjectMaxLength = 200

	// Max length of an access control subject name, when resolving names to IDs.
	TunnelConstraintsAccessControlSubjectNameMaxLength = 200

	// Maximum number of scopes in an access control entry.
	TunnelConstraintsAccessControlMaxScopes = 10

	// Regular expression that can match or validate tunnel cluster ID strings.
	//
	// Cluster IDs are alphanumeric; hyphens are not permitted.
	TunnelConstraintsClusterIDPattern = "[a-z][a-z0-9]{2,11}"

	// Characters that are valid in tunnel IDs. Includes numbers and lowercase letters,
	// excluding vowels and 'y' (to avoid accidentally generating any random words).
	TunnelConstraintsTunnelIDChars = "0123456789bcdfghjklmnpqrstvwxz"

	// Regular expression that can match or validate tunnel ID strings.
	//
	// Tunnel IDs are fixed-length and have a limited character set of numbers and lowercase
	// letters (minus vowels and y).
	TunnelConstraintsTunnelIDPattern = "[" + TunnelConstraintsTunnelIDChars + "]{8}"

	// Regular expression that can match or validate tunnel names.
	//
	// Tunnel names are alphanumeric and may contain hyphens. The pattern also allows an
	// empty string because tunnels may be unnamed.
	TunnelConstraintsTunnelNamePattern = "([a-z0-9][a-z0-9-]{1,58}[a-z0-9])|"

	// Regular expression that can match or validate tunnel or port tags.
	TunnelConstraintsTagPattern = "[\\w-=]{1,50}"

	// Regular expression that can match or validate tunnel domains.
	//
	// The tunnel service may perform additional contextual validation at the time the domain
	// is registered.
	TunnelConstraintsTunnelDomainPattern = "[0-9a-z][0-9a-z-.]{1,158}[0-9a-z]"

	// Regular expression that can match or validate an access control subject or
	// organization ID.
	//
	// The : and / characters are allowed because subjects may include IP addresses and
	// ranges. The @ character is allowed because MSA subjects may be identified by email
	// address.
	TunnelConstraintsAccessControlSubjectPattern = "[0-9a-zA-Z-._:/@]{0,200}"

	// Regular expression that can match or validate an access control subject name, when
	// resolving subject names to IDs.
	//
	// Note angle-brackets are only allowed when they wrap an email address as part of a
	// formatted name with email. The service will block any other use of angle-brackets, to
	// avoid any XSS risks.
	TunnelConstraintsAccessControlSubjectNamePattern = "[ \\w\\d-.,/'\"_@()<>]{0,200}"
)
var (
	// Regular expression that can match or validate tunnel cluster ID strings.
	//
	// Cluster IDs are alphanumeric; hyphens are not permitted.
	TunnelConstraintsClusterIDRegex = regexp.MustCompile(TunnelConstraintsClusterIDPattern)

	// Regular expression that can match or validate tunnel ID strings.
	//
	// Tunnel IDs are fixed-length and have a limited character set of numbers and lowercase
	// letters (minus vowels and y).
	TunnelConstraintsTunnelIDRegex = regexp.MustCompile(TunnelConstraintsTunnelIDPattern)

	// Regular expression that can match or validate tunnel names.
	//
	// Tunnel names are alphanumeric and may contain hyphens. The pattern also allows an
	// empty string because tunnels may be unnamed.
	TunnelConstraintsTunnelNameRegex = regexp.MustCompile(TunnelConstraintsTunnelNamePattern)

	// Regular expression that can match or validate tunnel or port tags.
	TunnelConstraintsTagRegex = regexp.MustCompile(TunnelConstraintsTagPattern)

	// Regular expression that can match or validate tunnel domains.
	//
	// The tunnel service may perform additional contextual validation at the time the domain
	// is registered.
	TunnelConstraintsTunnelDomainRegex = regexp.MustCompile(TunnelConstraintsTunnelDomainPattern)

	// Regular expression that can match or validate an access control subject or
	// organization ID.
	TunnelConstraintsAccessControlSubjectRegex = regexp.MustCompile(TunnelConstraintsAccessControlSubjectPattern)

	// Regular expression that can match or validate an access control subject name, when
	// resolving subject names to IDs.
	TunnelConstraintsAccessControlSubjectNameRegex = regexp.MustCompile(TunnelConstraintsAccessControlSubjectNamePattern)
)

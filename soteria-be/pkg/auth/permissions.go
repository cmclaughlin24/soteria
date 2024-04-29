package auth

import "strings"

type UserPermission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func PackPermissions(permissions []UserPermission) []string {
	if len(permissions) == 0 {
		return make([]string, 0)
	}

	m := make(map[string]string)

	for _, permission := range permissions {
		val, ok := m[permission.Resource]

		if !ok {
			m[permission.Resource] = permission.Action
		} else {
			m[permission.Resource] = val + "," + permission.Action
		}
	}

	pack := make([]string, 0, len(m))

	for resource, actions := range m {
		pack = append(pack, resource+"="+actions)
	}

	return pack
}

func UnpackPermissions(permissions []string) map[string][]string {
	if len(permissions) == 0 {
		return make(map[string][]string)
	}

	m := make(map[string][]string, len(permissions))

	for _, permission := range permissions {
		p := strings.Split(permission, "=")
		m[p[0]] = strings.Split(p[1], ",")
	}

	return m
}

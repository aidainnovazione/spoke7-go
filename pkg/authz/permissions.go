package authz

import "github.com/casbin/casbin/v2"

func GetJSPermissionForUser(e casbin.IEnforcer, user string, groups []string) (map[string]interface{}, error) {
	model := e.GetModel()
	m := map[string]interface{}{}

	m["m"] = model.ToText()

	pRules := [][]string{}
	e.GetFilteredPolicy(0, user)

	m["p"] = pRules

	gRules := [][]string{}
	for ptype := range model["g"] {
		policies, err := model.GetPolicy("g", ptype)
		if err != nil {
			return nil, err
		}
		for _, rules := range policies {
			gRules = append(gRules, append([]string{ptype}, rules...))
		}
	}
	m["g"] = gRules

	return m, nil
}

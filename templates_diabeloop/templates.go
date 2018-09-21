package templates_diabeloop

import (
	"fmt"

	"github.com/tidepool-org/hydrophone/models"
)

func NewDiabeloop() (models.Templates, error) {
	templates := models.Templates{}
	fmt.Printf("TEMPLATES:DIA")
	if template, err := NewCareteamInviteTemplate(); err != nil {
		return nil, fmt.Errorf("templates: failure to create careteam invite template: %s", err)
	} else {
		templates[template.Name()] = template
	}

	if template, err := NewNoAccountTemplate(); err != nil {
		return nil, fmt.Errorf("templates: failure to create no account template: %s", err)
	} else {
		templates[template.Name()] = template
	}

	if template, err := NewPasswordResetTemplate(); err != nil {
		return nil, fmt.Errorf("templates: failure to create password reset template: %s", err)
	} else {
		templates[template.Name()] = template
	}

	if template, err := NewSignupTemplate(); err != nil {
		return nil, fmt.Errorf("templates: failure to create signup template: %s", err)
	} else {
		templates[template.Name()] = template
	}

	if template, err := NewSignupClinicTemplate(); err != nil {
		return nil, fmt.Errorf("templates: failure to create signup clinic template: %s", err)
	} else {
		templates[template.Name()] = template
	}

	if template, err := NewSignupCustodialTemplate(); err != nil {
		return nil, fmt.Errorf("templates: failure to create signup custodial template: %s", err)
	} else {
		templates[template.Name()] = template
	}

	if template, err := NewSignupCustodialClinicTemplate(); err != nil {
		return nil, fmt.Errorf("templates: failure to create signup custodial clinic template: %s", err)
	} else {
		templates[template.Name()] = template
	}

	return templates, nil
}

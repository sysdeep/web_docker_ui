package service_model

type ServiceMode struct {
	Replicated *ReplicatedService `json:"replicated"`
	Global     *GlobalService     `json:"global"`
	// ReplicatedJob *ReplicatedJob     `json:",omitempty"`
	// GlobalJob     *GlobalJob         `json:",omitempty"`
}

// ReplicatedService is a kind of ServiceMode.
type ReplicatedService struct {
	Replicas int `json:"replicas"`
}

// GlobalService is a kind of ServiceMode.
type GlobalService struct{}

// Service represents a service.
type ServiceModel struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Mode      ServiceMode `json:"mode"`
	Image     string      `json:"image"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	// Spec      serviceSpec `json:"spec"`
	// PreviousSpec *ServiceSpec  `json:",omitempty"`
	// Endpoint     Endpoint      `json:",omitempty"`
	// UpdateStatus *UpdateStatus `json:",omitempty"`
	//
	// // ServiceStatus is an optional, extra field indicating the number of
	// // desired and running tasks. It is provided primarily as a shortcut to
	// // calculating these values client-side, which otherwise would require
	// // listing all tasks for a service, an operation that could be
	// // computation and network expensive.
	// ServiceStatus *ServiceStatus `json:",omitempty"`
	//
	// // JobStatus is the status of a Service which is in one of ReplicatedJob or
	// // GlobalJob modes. It is absent on Replicated and Global services.
	// JobStatus *JobStatus `json:",omitempty"`
}

package deployment

type MetadataOwnerReference struct {
	apiVersion         string
	blockOwnerDeletion bool
	controller         bool
	kind               string
	name               string
	uid                string
}

type MetadataManagedField struct {
	apiVersion string
	fieldsType string
	fieldsV1   map[string]string
	manager    string
	operation  string
	time       string
}

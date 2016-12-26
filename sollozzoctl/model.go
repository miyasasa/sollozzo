package sollozzoctl

type Version struct {
	Major       uint8
	Minor       uint8
	BuildNumber uint16
}

type Project struct {
	Key     string
	Desc    string
	Version *Version
}

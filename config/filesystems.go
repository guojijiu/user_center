package config

const (
	MAX_MULTIPART_MEMORY = 8 << 20 // 8 MiB
)

type filesystems struct {
	Default string
	Cloud   string
	Disks   Disks
}

type Disks struct {
	Local Local
}

type Local struct {
	Driver string
	Root   string
}

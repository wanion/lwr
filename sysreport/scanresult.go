package sysreport

// ScanResult structure to match JSON used by LSAgent.
type ScanResult struct {
	ScanResultType int      `json:"ScanResultType"`
	ScanType       int      `json:"ScanType"`
	LsAgentVersion string   `json:"LsAgentVersion"`
	MacAddresses   []string `json:"MacAddresses"`
	AssetErrors    []string `json:"AssetErrors"`
	Uptime         int      `json:"Uptime"`
	IPAddress      string   `json:"IpAddress"`
	AssetName      string   `json:"AssetName"`
	SystemInfo     struct {
		PCManufacturer      *string `json:"PCManufacturer"`
		TotalMemory         string  `json:"TotalMemory"`
		CPUCount            int     `json:"CpuCount"`
		CPUName             string  `json:"CpuName"`
		NetworkNodeHostname string  `json:"NetworkNodeHostname"`
		KernelName          string  `json:"KernelName"`
		KernelRelease       string  `json:"KernelRelease"`
		KernelVersion       string  `json:"KernelVersion"`
		MachineHardwareName string  `json:"MachineHardwareName"`
		ProcessorType       string  `json:"ProcessorType"`
		HardwarePlatform    string  `json:"HardwarePlatform"`
		OperatingSysem      string  `json:"OperatingSystem"`
		OSRelease           string  `json:"OSRelease"`
		Model               string  `json:"Model"`
		Serial              string  `json:"Serial"`
		SystemSku           *string `json:"SystemSku"`
		Firmware            *string `json:"Firmware"`
	} `json:"SystemInfo"`
	GraphicsCards    []scanDevice       `json:"GraphicsCards"`
	PciCards         []scanDevice       `json:"PciCards"`
	SoundCards       []scanDevice       `json:"SoundCards"`
	NetworkDetection []networkInterface `json:"NetworkDetection"`
	HardDisks        []hardDisk         `json:"HardDisks"`
	OpticalDrives    []string           `json:"OpticalDrives"`
	Volumes          []string           `json:"Volumes"`
	Software         []softwarePackage  `json:"Software"`
	// DmidecodeNotInstalled bool               `json:"DmidecodeNotInstalled"`
	// SmBiosInfo            smBios             `json:"SmBiosInfo"`
}

type hardDisk struct {
	Filesystem string `json:"Filesystem"`
	Size       string `json:"Size"`
	Used       string `json:"Used"`
	Available  string `json:"Available"`
	Percentage string `json:"Percentage"`
	MountedOn  string `json:"MountedOn"`
}

type scanDevice struct {
	DeviceID              string `json:"DeviceId"`
	Type                  string `json:"Type"`
	Name                  string `json:"Name"`
	Manufacturer          string `json:"Manufacturer"`
	SubsystemName         string `json:"SubsystemName"`
	SubsystemManufacturer string `json:"SubsystemManufacturer"`
}

type softwarePackage struct {
	Name         string `json:"Name"`
	Version      string `json:"Version"`
	Description  string `json:"Description"`
	InstallDate  string `json:"InstallDate"`
	Release      string `json:"Release"`
	Architecture string `json:"Architecture"`
}

type smBios struct {
	BiosInfo struct {
		Vendor      scanSetting `json:"Vendor"`
		Version     scanSetting `json:"Version"`
		ReleaseDate scanSetting `json:"ReleaseDate"`
		Address     scanSetting `json:"Address"`
		RuntimeSize scanSetting `json:"RuntimeSize"`
		ROMSize     scanSetting `json:"ROMSize"`
	} `json:"BiosBiosInfo"`
	SystemInfo struct {
		Manufacturer scanSetting `json:"Manufacturer"`
		ProductName  scanSetting `json:"ProductName"`
		Version      scanSetting `json:"Version"`
		Serial       scanSetting `json:"Serial"`
		UUID         scanSetting `json:"UUID"`
		WakeupTime   scanSetting `json:"WakeupTime"`
		SystemSku    scanSetting `json:"SystemSku"`
	} `json:"SystemInfo"`
	BootStatus struct {
		BootStatus scanSetting `json:"BootStatus"`
	} `json:"BootStatus"`
	BaseboardInfo struct {
		Manufacturer      scanSetting `json:"Manufacturer"`
		ProductName       scanSetting `json:"ProductName"`
		Version           scanSetting `json:"Version"`
		SerialNumber      scanSetting `json:"SerialNumber"`
		LocationInChassis scanSetting `json:"LocationInChassis"`
		Type              scanSetting `json:"Type"`
	} `json:"BaseBoardInfo"`
	MemoryControllerInfo struct {
		SupportedInterleave scanSetting `json:"SupportedInterleave"`
		CurrentInterleave   scanSetting `json:"CurrentInterleave"`
		MaxMemoryModuleSize scanSetting `json:"MaxMemoryModuleSize"`
		MaxTotalMemorySize  scanSetting `json:"MaxTotalMemorySize"`
		MemoryModuleVoltage scanSetting `json:"MemoryModuleVoltage"`
		NumberOfSlots       scanSetting `json:"NumberOfSlots"`
		SupportedSpeeds     scanSetting `json:"SupportedSpeeds"`
		SupportedMemTypes   scanSetting `json:"SupportedMemTypes"`
	} `json:"MemoryControllerInfo"`
	Enclosure struct {
		Manufacturer   scanSetting `json:"Manufacturer"`
		SMBIOSAssetTag scanSetting `json:"SMBIOSAssetTag"`
		SerialNumber   scanSetting `json:"SerialNumber"`
		Version        scanSetting `json:"Version"`
		LockPresent    scanSetting `json:"LockPresent"`
		SecurityStatus scanSetting `json:"SecurityStatus"`
		ChassisTypes   scanSetting `json:"ChassisTypes"`
	} `json:"Enclosure"`
	ProcessorInfo []struct {
		Type          scanSetting `json:"Type"`
		Socket        scanSetting `json:"Socket"`
		Voltage       scanSetting `json:"Voltage"`
		MaxSpeed      scanSetting `json:"MaxSpeed"`
		Status        scanSetting `json:"Status"`
		Family        scanSetting `json:"Family"`
		ExternalClock scanSetting `json:"ExternalClock"`
		CurrentSpeed  scanSetting `json:"CurrentSpeed"`
		SerialNumber  scanSetting `json:"SerialNumber"`
		Manufacturer  scanSetting `json:"Manufacturer"`
		ID            scanSetting `json:"ID"`
		Version       scanSetting `json:"Version"`
	} `json:"ProcessorInfo"`
	MemoryModuleInfo []struct {
		Socket          scanSetting `json:"Socket"`
		BankConnections scanSetting `json:"BankConnections"`
		CurrentSpeed    scanSetting `json:"CurrentSpeed"`
		Type            scanSetting `json:"Type"`
		InstalledSize   scanSetting `json:"InstalledSize"`
		EnabledSize     scanSetting `json:"EnabledSize"`
		ErrorStatus     scanSetting `json:"ErrorStatus"`
	} `json:"MemoryModuleInfo"`
	MemoryDeviceInfo []struct {
		TotalWidth   scanSetting `json:"TotalWidth"`
		DataWidth    scanSetting `json:"DataWidth"`
		Size         scanSetting `json:"Size"`
		FormFactor   scanSetting `json:"FormFactor"`
		Set          scanSetting `json:"Set"`
		Locator      scanSetting `json:"Locator"`
		BankLocator  scanSetting `json:"BankLocator"`
		Type         scanSetting `json:"Type"`
		TypeDetail   scanSetting `json:"TypeDetail"`
		Speed        scanSetting `json:"Speed"`
		Manufacturer scanSetting `json:"Manufacturer"`
		SerialNumber scanSetting `json:"SerialNumber"`
	} `json:"MemoryDeviceInfo"`
}

type scanSetting struct {
	Value     string `json:"Value"`
	BiosName  string `json:"BiosName"`
	MaxLength int    `json:"MaxLength"`
}

type networkInterface struct {
	Name      string  `json:"Name"`
	LinkEncap string  `json:"LinkEncap"`
	Mac       *string `json:"Mac"`
	IP4       string  `json:"Ip4"`
	Broadcast string  `json:"Broadcast"`
	Mask      string  `json:"Mask"`
	IP6       string  `json:"Ip6"`
	Scope     string  `json:"Scope"`
	Gateway   string  `json:"Gateway"`
}

/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package system

import "math"

const (
	// subfs name
	CgroupCPUDir     string = "cpu/"
	CgroupCPUSetDir  string = "cpuset/"
	CgroupCPUacctDir string = "cpuacct/"
	CgroupMemDir     string = "memory/"
)

const (
	CPUStatFileName   = "cpu.stat"
	CPUSharesFileName = "cpu.shares"
	CPUCFSQuotaName   = "cpu.cfs_quota_us"
	CPUCFSPeriodName  = "cpu.cfs_period_us"
	CPUBVTWarpNsName  = "cpu.bvt_warp_ns"
	CPUSFileName      = "cpuset.cpus"

	CpuacctStatFileName = "cpuacct.stat"

	MemWmarkRatioFileName = "memory.wmark_ratio"
	MemHighFileName       = "memory.high"
	MemoryLimitFileName   = "memory.limit_in_bytes"
	MemStatFileName       = "memory.stat"
)

var (
	MemHighValidator = &RangeValidator{name: MemHighFileName, min: 0, max: math.MaxInt64}
)

var (
	CPUStat      = CgroupFile{ResourceFileName: CPUStatFileName, Subfs: CgroupCPUDir, IsAliOS: false}
	CPUShares    = CgroupFile{ResourceFileName: CPUSharesFileName, Subfs: CgroupCPUDir, IsAliOS: false}
	CPUCFSQuota  = CgroupFile{ResourceFileName: CPUCFSQuotaName, Subfs: CgroupCPUDir, IsAliOS: false}
	CPUCFSPeriod = CgroupFile{ResourceFileName: CPUCFSPeriodName, Subfs: CgroupCPUDir, IsAliOS: false}

	CPUSet = CgroupFile{ResourceFileName: CPUSFileName, Subfs: CgroupCPUSetDir, IsAliOS: false}

	CpuacctStat = CgroupFile{ResourceFileName: CpuacctStatFileName, Subfs: CgroupCPUacctDir, IsAliOS: false}

	MemStat     = CgroupFile{ResourceFileName: MemStatFileName, Subfs: CgroupMemDir, IsAliOS: false}
	MemHigh     = CgroupFile{ResourceFileName: MemHighFileName, Subfs: CgroupMemDir, IsAliOS: true, Validator: MemHighValidator}
	MemoryLimit = CgroupFile{ResourceFileName: MemoryLimitFileName, Subfs: CgroupMemDir, IsAliOS: false}
)

type CgroupFile struct {
	ResourceFileName string
	Subfs            string
	IsAliOS          bool
	Validator        Validate
}
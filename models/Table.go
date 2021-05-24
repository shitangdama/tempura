package models

// https://www.cnblogs.com/itsone/p/13602999.html
// Table is
type Table struct {
	ID         int32  `json:"id"`
	Schema     string `json:"schema"`
	Name       string `json:"name"`
	RlsEnabled bool   `json:"rls_enabled"`
	RlsForced  bool   `json:"rls_forced"`
	// relreplident string //用于为行构成“副本标识”的列：d = default（主键，如果有的话），n =无，f =所有列 i =具有indisreplident set的索引或default
	Bytes            int64  `json:"bytes"`
	Size             int64  `json:"size"`
	LiveRowsEstimate int    `json:"live_rows_estimate"`
	DeadRowsEstimate int    `json:"dead_rows_estimate"`
	Comment          string `json:"comment"`
}

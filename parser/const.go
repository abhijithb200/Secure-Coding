package parser

const (
	_ int = iota + 9

	// Reflected XSS
	CONCAT_WITH_VULN_ARRAY // 3
	CONCAT_WITH_TAINT_VAR  //  4
)

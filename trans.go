package gbutility

func GetDocumentNumber(p string) string {
	return GetYMDTrans() + lpad(p, "0", 4)
}

func GetDocumentNumberArea(area string, p string) string {
	return GetYMDTrans() + lpad(area, "0", 2) + lpad(p, "0", 4)
}

func lpad(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}

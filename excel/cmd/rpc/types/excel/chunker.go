package excel

type ChunkerSrv []byte
const chunkSize = 64 * 1024 // 64 KiB

func (c ChunkerSrv) ChunkerSend(srv User_ParseExcelClient) error {
	chnk := &ExcelRequest{}
	for currentByte := 0; currentByte < len(c); currentByte += chunkSize {
		if currentByte+chunkSize > len(c) {
			chnk.File = c[currentByte:len(c)]
		} else {
			chnk.File = c[currentByte : currentByte+chunkSize]
		}
		if err := srv.Send(chnk); err != nil {
			return err
		}
	}
	return nil
}
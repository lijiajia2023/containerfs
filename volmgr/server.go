package volmgr

type Server struct{}

func (s *Server) MakeFilesystem(inodeQuota, spaceQuota int64) (int, error) {}

func (s *Server) TuneFilesystem(volumeID int, inodeQuota, spaceQuota int64) error {}

func (s *Server) StatFilesystem(volumeID int) (info *VolumeInfo, e error) {}

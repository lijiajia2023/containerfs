package volmgr

type Volume struct {
	inodeGroupID int
	blockGroups  []int
}
